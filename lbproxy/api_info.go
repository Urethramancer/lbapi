package main

import "net/http"

// apiInfo simply passes the InfoDump configuration data to the client.
func apiInfo(w http.ResponseWriter, r *http.Request) error {
	return respond(w, &cfg.InfoDump)
}
