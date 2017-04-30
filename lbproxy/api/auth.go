package api

import (
	"fmt"
	"net/url"

	"github.com/Urethramancer/lbapi"
)

func (c *Client) Authenticate() bool {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return false
	}

	u.Path = PathAuth
	q := u.Query()
	q.Set("username", c.Username)
	q.Set("password", c.Password)
	u.RawQuery = q.Encode()

	res, err := lbapi.GetResponse(c.Client, u.String())
	if err != nil {
		return false
	}

	s := *res
	fmt.Printf("%#v\n", s)
	return true
}
