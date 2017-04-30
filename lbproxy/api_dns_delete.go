package main

import "net/http"

// apiDNSDelete removes a record from a domain.
func apiDNSDelete(w http.ResponseWriter, r *http.Request) error {
	return respond(w, &cfg.InfoDump)
}
