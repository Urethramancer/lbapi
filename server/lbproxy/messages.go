package main

const (
	// StatusOK means no problems.
	StatusOK = iota
	// StatusNoCredentials means wrong inputs to login.
	StatusNoCredentials = iota
	// StatusNoAuth means no or wrong auth token was passed.
	StatusNoAuth = iota
	// StatusAuthFailed means username or password was wrong, or account non-existent.
	StatusAuthFailed = iota
	// StatusNoAccess means the token given does not have access to the requested data.
	StatusNoAccess = iota
)

var messages = map[int]string{
	StatusOK:            "OK",
	StatusNoCredentials: "Missing username or password.",
	StatusNoAuth:        "No authorisation token provided.",
	StatusAuthFailed:    "Login failed.",
	StatusNoAccess:      "Not accessible with provided credentials.",
}
