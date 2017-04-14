package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) ChangeARecord(domain, oldip, newip, host string, ttl int64, six bool) error {
	if six {
		return c.changeRecord("api/dns/manage/update-ipv6-record.json", domain, oldip, newip, host, ttl)
	}
	return c.changeRecord("api/dns/manage/update-ipv4-record.json", domain, oldip, newip, host, ttl)
}

func (c *Client) changeRecord(call, domain, oldip, newip, host string, ttl int64) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = call
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("current-value", oldip)
	q.Set("new-value", newip)
	if host != "" {
		q.Set("host", host)
	}
	if ttl == 0 || ttl < 7200 {
		ttl = 7200
	}
	q.Set("ttl", fmt.Sprintf("%d", ttl))
	u.RawQuery = q.Encode()

	res, err := c.postResponse(u.String())
	if err != nil {
		return err
	}

	list := *res
	if list["status"] == "ERROR" {
		return errors.New((fmt.Sprintf("%v", list["message"])))
	}

	return nil
}
