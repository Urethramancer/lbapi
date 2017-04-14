package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

// ADDARecord adds A or AAAA records.
func (c *Client) AddARecord(domain, address, host string, ttl int64, six bool) error {
	if six {
		return c.addRecord("api/dns/manage/add-ipv6-record.json", domain, address, host, ttl)
	}
	return c.addRecord("api/dns/manage/add-ipv4-record.json", domain, address, host, ttl)
}

// AddCNAME does exactly that.
func (c *Client) AddCNAME(domain, value, host string, ttl int64) error {
	return c.addRecord("api/dns/manage/add-cname-record.json", domain, value, host, ttl)
}

// AddMX adds MX records for mail servers.
func (c *Client) AddMX(domain, value, host string, ttl int64, priority uint) error {
	return c.addRecordPri("api/dns/manage/add-mx-record.json", domain, value, host, ttl, priority)
}

// AddNS adds name server records.
func (c *Client) AddNS(domain, value, host string, ttl int64, priority uint) error {
	return c.addRecord("api/dns/manage/add-ns-record.json", domain, value, host, ttl)
}

// AddTXT adds TXT records.
func (c *Client) AddTXT(domain, value, host string, ttl int64, priority uint) error {
	return c.addRecord("api/dns/manage/add-txt-record.json", domain, value, host, ttl)
}

// AddSRV adds SRV records.
func (c *Client) AddSRV(domain, value, host string, ttl int64, priority, port, weight uint) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = "api/dns/manage/add-srv-record.json"
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", value)
	if host != "" {
		q.Set("host", host)
	}
	if ttl == 0 || ttl < 7200 {
		ttl = 7200
	}
	if ttl > 0 {
		q.Set("ttl", fmt.Sprintf("%d", ttl))
	}
	if priority > 0 {
		q.Set("priority", fmt.Sprintf("%d", priority))
	}
	if port > 0 {
		q.Set("port", fmt.Sprintf("%d", port))
	}
	if weight > 0 {
		q.Set("weight", fmt.Sprintf("%d", weight))
	}
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

func (c *Client) addRecord(call, domain, address, host string, ttl int64) error {
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

func (c *Client) addRecordPri(call, domain, address, host string, ttl int64, priority uint) error {
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
	q.Set("value", address)
	if host != "" {
		q.Set("host", host)
	}
	if ttl == 0 || ttl < 7200 {
		ttl = 7200
	}
	q.Set("ttl", fmt.Sprintf("%d", ttl))
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
