package main

import "net/http"

// apiDNSNuke wipes out all but the primary A and AAAA records for a domain.
func apiDNSNuke(w http.ResponseWriter, r *http.Request) error {
	return respond(w, &cfg.InfoDump)
}
