package main

import "net/http"

// apiDNSAdd adds a record for a domain.
func apiDNSAdd(w http.ResponseWriter, r *http.Request) error {
	return respond(w, &cfg.InfoDump)
}
