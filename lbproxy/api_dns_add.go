package main

import (
	"net/http"
	"strings"

	"github.com/Urethramancer/lbapi"
)

// apiDNSAdd endpoint.
func apiDNSAdd(w http.ResponseWriter, r *http.Request) error {
	var res Result
	var err error
	q := r.URL.Query()
	t := getToken(q.Get("token"))
	if t == nil || !t.IsValid(r) {
		res.SetStatus(StatusNoAuth)
		return respond(w, &res)
	}

	rec := strings.ToUpper(q.Get("type"))
	domain := q.Get("domain")
	value := q.Get("value")
	host := q.Get("host")
	x := q.Get("ttl")
	if x == "" {
		x = "7200"
	}
	ttl := lbapi.Atoi(x)
	pri := uint16(lbapi.Atoi(q.Get("priority")))
	port := uint16(lbapi.Atoi(q.Get("port")))
	weight := uint16(lbapi.Atoi(q.Get("weight")))
	switch rec {
	case "A":
		err = client.AddARecord(domain, value, host, ttl, false)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	case "AAAA":
		err = client.AddARecord(domain, value, host, ttl, true)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	case "CNAME":
		err = client.AddCNAME(domain, value, host, ttl)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	case "MX":
		err = client.AddMX(domain, value, host, ttl, pri)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	case "NS":
		err = client.AddNS(domain, value, host, ttl, pri)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	case "TXT":
		err = client.AddTXT(domain, value, host, ttl, pri)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	case "SRV":
		err = client.AddSRV(domain, value, host, ttl, pri, port, weight)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
	}
	res.SetStatus(StatusOK)
	return respond(w, &res)
}
