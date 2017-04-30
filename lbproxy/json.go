package main

import (
	"encoding/json"
	"io"
)

// Result is returned from many JSON requests.
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
