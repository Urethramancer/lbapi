package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/Urethramancer/lbapi/lbproxy/api"
	"github.com/gorilla/mux"
)

// Servers hold the two possible web servers which can be started.
var webserver *http.Server

func initWeb() {
	var address string

	if cfg.Security.SSL {
		address = net.JoinHostPort(cfg.Web.Address, cfg.Web.Port)
		cfg.Web.url = "https://" + cfg.Web.Domain
		if cfg.Web.Port != "443" {
			cfg.Web.url += ":" + cfg.Web.Port
		}
	} else {
		address = net.JoinHostPort(cfg.Web.Address, cfg.Web.Port)
		cfg.Web.url = "http://" + cfg.Web.Domain
		if cfg.Web.Port != "80" {
			cfg.Web.url += ":" + cfg.Web.Port
		}
	}

	router := mux.NewRouter().StrictSlash(true)
	api := "/"
	if cfg.APIPath != "" {
		api += cfg.APIPath + "/"
	}
	sub := router.Host(cfg.Web.Domain).Methods("POST").PathPrefix(api).Subrouter()

	if cfg.Security.SSL {
		secure := sub.Schemes("https").Subrouter()
		initSSLHandlers(secure)
		syslog("Starting secure web server on %s (%s)", address, cfg.Web.url)
		go startHTTPS(address, secure)
	} else {
		insecure := sub.Schemes("http").Subrouter()
		initHandlers(insecure)
		syslog("Starting plain web server on %s (%s)", address, cfg.Web.url)
		go startHTTP(address, insecure)
	}
}

func startHTTP(address string, r http.Handler) {
	webserver = &http.Server{
		Addr:    address,
		Handler: r,
		// Safe numbers which should help against DoS.
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
	err := webserver.ListenAndServe()
	if err == http.ErrServerClosed {
		syslog("HTTP server shut down cleanly.")
	}
}

func startHTTPS(address string, r http.Handler) {
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	config.NextProtos = []string{"http/1.1"}
	// This should give a pretty good score from SSL Labs' tests.
	config.MinVersion = tls.VersionTLS12
	config.CurvePreferences = []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256}
	config.PreferServerCipherSuites = true
	config.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	}
	webserver = &http.Server{
		Addr:         address,
		Handler:      r,
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		TLSConfig:    config,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	// The certificate should be the fullchain version if using Let's Encrypt.
	err := webserver.ListenAndServeTLS(cfg.Security.Certificate, cfg.Security.Key)
	if err == http.ErrServerClosed {
		syslog("HTTPS server shut down: %s", err)
	}
}

func stopServers() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	err := webserver.Shutdown(ctx)
	if err != nil {
		crit("Error: %s", err.Error())
	}
}

type (
	// Handler is the signature for web route handlers.
	Handler func(w http.ResponseWriter, r *http.Request) error
)

// Handle walks through its chain of middleware for a web request.
func Handle(handlers ...Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle handlers until one can't handle it.
		for _, handler := range handlers {
			err := handler(w, r)
			if err != nil {
				crit("Error: %s", err.Error())
				return
			}
		}
	})
}

func addJSONHeaders(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return nil
}

// addSSLHeaders ensures browsers stick to SSL after they start using it.
func addSSLHeaders(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Strict-Transport-Security", "max-age=15768000; includeSubDomains")
	w.Header().Set("Content-Type", "charset=UTF-8")
	return nil
}

func initHandlers(r *mux.Router) {
	r.Handle(api.PathAuth, Handle(addJSONHeaders, apiAuth))
	r.Handle(api.PathInfo, Handle(addJSONHeaders, apiInfo))

	r.Handle(api.PathDNSGet, Handle(addJSONHeaders, apiDNSGet))

	r.Handle(api.PathDNSAddIPv4, Handle(addJSONHeaders, apiDNSAddIPv4))
	r.Handle(api.PathDNSAddIPv6, Handle(addJSONHeaders, apiDNSAddIPv6))
	r.Handle(api.PathDNSAddCNAME, Handle(addJSONHeaders, apiDNSAddCNAME))
	r.Handle(api.PathDNSAddMX, Handle(addJSONHeaders, apiDNSAddMX))
	r.Handle(api.PathDNSAddNS, Handle(addJSONHeaders, apiDNSAddNS))
	r.Handle(api.PathDNSAddTXT, Handle(addJSONHeaders, apiDNSAddTXT))
	r.Handle(api.PathDNSAddSRV, Handle(addJSONHeaders, apiDNSAddSRV))

	r.Handle(api.PathDNSEditIPv4, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditIPv6, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditCNAME, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditMX, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditNS, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditTXT, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditSRV, Handle(addJSONHeaders, apiInfo))

	r.Handle(api.PathDNSDeleteIPv4, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteIPv6, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteCNAME, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteMX, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteNS, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteTXT, Handle(addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteSRV, Handle(addJSONHeaders, apiInfo))

	r.Handle(api.PathDNSNuke, Handle(addJSONHeaders, apiInfo))
}

func initSSLHandlers(r *mux.Router) {
	r.Handle(api.PathAuth, Handle(addSSLHeaders, addJSONHeaders, apiAuth))
	r.Handle(api.PathInfo, Handle(addSSLHeaders, addJSONHeaders, apiInfo))

	r.Handle(api.PathDNSGet, Handle(addSSLHeaders, addJSONHeaders, apiDNSGet))

	r.Handle(api.PathDNSAddIPv4, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddIPv4))
	r.Handle(api.PathDNSAddIPv6, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddIPv6))
	r.Handle(api.PathDNSAddCNAME, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddCNAME))
	r.Handle(api.PathDNSAddMX, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddMX))
	r.Handle(api.PathDNSAddNS, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddNS))
	r.Handle(api.PathDNSAddTXT, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddTXT))
	r.Handle(api.PathDNSAddSRV, Handle(addSSLHeaders, addJSONHeaders, apiDNSAddSRV))

	r.Handle(api.PathDNSEditIPv4, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditIPv6, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditCNAME, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditMX, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditNS, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditTXT, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSEditSRV, Handle(addSSLHeaders, addJSONHeaders, apiInfo))

	r.Handle(api.PathDNSDeleteIPv4, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteIPv6, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteCNAME, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteMX, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteNS, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteTXT, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
	r.Handle(api.PathDNSDeleteSRV, Handle(addSSLHeaders, addJSONHeaders, apiInfo))

	r.Handle(api.PathDNSNuke, Handle(addSSLHeaders, addJSONHeaders, apiInfo))
}
