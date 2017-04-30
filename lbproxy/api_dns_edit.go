package main

import "net/http"

// apiDNSEdit modifies a record for a domain.
func apiDNSEdit(w http.ResponseWriter, r *http.Request) error {
	return respond(w, &cfg.InfoDump)
}
