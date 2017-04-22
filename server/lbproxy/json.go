package main

import (
	"encoding/json"
	"io"
)

// Request is a generic all-purpose JSON request structure.
type Request struct {
	// Token is required for all requests to manipulate data, or read privileged information.
	Token string `json:"token,omitempty"`
	// Username is used for authentication requests only.
	Username string `json:"username,omitempty"`
	// Password is used for authentication requests only.
	Password string `json:"password,omitempty"`
	// ID is used to retrieve notes and lists.
	ID uint64 `json:"id,omitempty"`
	// Note is used when adding notes.
	// For changing public/private status
	Public bool   `json:"public,omitempty"`
	Note   string `json:"note,omitempty"`
	// List is used when adding lists.
	List string `json:"list,omitempty"`
	// ListItems are used when creating lists or appending to them.
	ListItems []string `json:"listitems,omitempty"`
	// Start of a range to fetch.
	Start int `json:"start,omitempty"`
	// Count is the number of items to fetch.
	Count int `json:"count,omitempty"`
}

// Result is returned from any JSON request.
type Result struct {
	Status     string `json:"status,omitempty"`
	StatusCode int    `json:"statuscode"`
	Token      string `json:"token,omitempty"`
}

// SetStatus fills in the status text from a code.
func (r *Result) SetStatus(code int) {
	s, ok := messages[code]
	if !ok {
		r.StatusCode = StatusOK
		r.Status = messages[StatusOK]
		return
	}
	r.StatusCode = code
	r.Status = s
}

func respond(w io.Writer, r interface{}) error {
	res, err := json.Marshal(r)
	if err != nil {
		return err
	}

	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}
