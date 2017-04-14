package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

// ADDARecord adds A or AAAA records.
func (c *Client) AddARecord(domain, address, host string, ttl int64, six bool) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	if six {
		u.Path = "api/dns/manage/add-ipv6-record.json"
	} else {
		u.Path = "api/dns/manage/add-ipv4-record.json"
	}
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", address)
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

// AddCNAME does exactly that.
func (c *Client) AddCNAME(domain, address, host string, ttl int64) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = "api/dns/manage/add-cname-record.json"
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", address)
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

func (c *Client) AddMX(domain, address, host string, ttl int64, priority uint) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = "api/dns/manage/add-mx-record.json"
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", address)
	if host != "" {
		q.Set("host", host)
	}
	if ttl == 0 || ttl < 7200 {
		ttl = 7200
	}
	q.Set("ttl", fmt.Sprintf("%d", ttl))
	if priority < 0 {
		priority = 0
	}
	q.Set("priority", fmt.Sprintf("%d", priority))
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
