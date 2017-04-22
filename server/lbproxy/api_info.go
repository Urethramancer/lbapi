package main

import "net/http"

func apiInfo(w http.ResponseWriter, r *http.Request) error {
	var info InfoDump
	return respond(w, &info)
}
