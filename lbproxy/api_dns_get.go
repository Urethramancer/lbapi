package main

import "net/http"
import "github.com/Urethramancer/lbapi"
import "strings"

// apiDNSGet implementation.
func apiDNSGet(w http.ResponseWriter, r *http.Request) error {
	var res Result
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
	var list lbapi.DNSRecords
	var count int64
	page := 1
	for {
		recs, err := client.GetDNSRecords(domain, value, host, rec, page)
		if err != nil {
			res.StatusCode = StatusError
			res.Status = err.Error()
			return respond(w, &res)
		}
		list = append(list, recs.Records...)
		page++
		count += recs.Count
		if count >= recs.MaxRecords {
			break
		}
	}

	return respond(w, &list)
}
