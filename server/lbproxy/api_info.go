package main

import "net/http"

// InfoDump is just that - a collection of information relating to the service.
type InfoDump struct {
}

func apiInfo(w http.ResponseWriter, r *http.Request) error {
	var info InfoDump
	return respond(w, &info)
}
