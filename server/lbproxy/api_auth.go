package main

import (
	"net/http"
	"strings"
)

func apiAuth(w http.ResponseWriter, r *http.Request) error {
	uname := strings.TrimSpace(r.URL.Query().Get("username"))
	pw := r.URL.Query().Get("password")

	var res Result
	cd, err := client.Authenticate(uname, pw)
	if err != nil {
		res.SetStatus(StatusAuthFailed)
		return respond(w, &res)
	}

	t := createToken(0, r)
	syslog("User %s (%d) logged in.", cd.Email, cd.ID)
	res.SetStatus(StatusOK)
	res.Token = t.Hash
	return respond(w, &res)
}
