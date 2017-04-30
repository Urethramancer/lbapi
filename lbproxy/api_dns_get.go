package main

import "net/http"

// apiDNSGet gets records for a domain.
func apiDNSGet(w http.ResponseWriter, r *http.Request) error {

	return respond(w, &cfg.InfoDump)
}
