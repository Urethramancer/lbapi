package main

import (
	"net/http"
	"time"
)

// Token results from authentication.
type Token struct {
	Hash    string
	IP      string
	UserID  int64
	Expires time.Time
}

var apitokens = make(map[string]*Token)

func getToken(hash string) *Token {
	t := apitokens[hash]
	if t == nil {
		return nil
	}

	return t
}

// createToken creates a token with an expiry and the user's IP address, then adds it to the token map.
func createToken(id int64, r *http.Request) *Token {
	hash := hashString(genString(32))
	t := &Token{
		Hash:    hash,
		IP:      getVisitorIP(r),
		Expires: time.Now().Add(cfg.sessionlength),
	}
	apitokens[hash] = t
	return t
}

// IsValid checks for valid IP address and expiry time, then refreshes the token.
func (t *Token) IsValid(r *http.Request) bool {
	if getVisitorIP(r) != t.IP {
		return false
	}

	if t.Expires.Unix() < time.Now().Unix() {
		return false
	}

	// Reset expiry to a new period if it's still valid.
	t.Refresh()
	return true
}

// Refresh sets the lifetime to cfg.sessionLength.
func (t *Token) Refresh() {
	t.Expires = time.Now().Add(cfg.sessionlength)
}

// Delete the token.
func (t *Token) Delete() {
	delete(apitokens, t.Hash)
}

// clearTokens is run by the janitor every now and then to clear out expired tokens.
func clearTokens() {
	for _, t := range apitokens {
		if t.Expires.After(time.Now()) {
			t.Delete()
		}
	}
}
