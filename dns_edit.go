package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

// EditARecord modifies A or AAAA records.
func (c *Client) EditARecord(domain, oldip, newip, host string, ttl int64, six bool) error {
	if six {
		return c.editRecord(apiDNSUpdateIPv6, domain, oldip, newip, host, ttl)
	}
	return c.editRecord(apiDNSUpdateIPv4, domain, oldip, newip, host, ttl)
}

// EditCNAME modifies CNAME (canonical name) records.
func (c *Client) EditCNAME(domain, oldip, newip, host string, ttl int64) error {
	return c.editRecord(apiDNSUpdateCNAME, domain, oldip, newip, host, ttl)
}

// EditMX modifies MX (mail server) records.
func (c *Client) EditMX(domain, oldip, newip, host string, ttl int64, priority uint16) error {
	return c.editRecordPri(apiDNSUpdateMX, domain, oldip, newip, host, ttl, priority)
}

// EditNS modifies NS (nameserver) records.
func (c *Client) EditNS(domain, oldip, newip, host string, ttl int64) error {
	return c.editRecord(apiDNSUpdateNS, domain, oldip, newip, host, ttl)
}

// EditTXT modfies TXT records.
func (c *Client) EditTXT(domain, oldip, newip, host string, ttl int64) error {
	return c.editRecord(apiDNSUpdateTXT, domain, oldip, newip, host, ttl)
}

func (c *Client) editRecord(call, domain, oldip, newip, host string, ttl int64) error {
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

	res, err := PostResponse(c.Client, u.String())
	if err != nil {
		return err
	}

	list := *res
	if list["status"] == "ERROR" {
		return errors.New((fmt.Sprintf("%v", list["message"])))
	}

	return nil
}

func (c *Client) editRecordPri(call, domain, oldip, newip, host string, ttl int64, priority uint16) error {
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
	q.Set("priority", fmt.Sprintf("%d", priority))
	u.RawQuery = q.Encode()

	res, err := PostResponse(c.Client, u.String())
	if err != nil {
		return err
	}

	list := *res
	if list["status"] == "ERROR" {
		return errors.New((fmt.Sprintf("%v", list["message"])))
	}

	return nil
}

// EditSRV modifies a SRV (service) record.
func (c *Client) EditSRV(domain, oldval, newval, host string, ttl int64, priority, port, weight uint) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = apiDNSUpdateSRV
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("current-value", oldval)
	q.Set("new-value", newval)
	q.Set("host", host)
	if ttl < 7200 {
		ttl = 7200
	}
	q.Set("ttl", fmt.Sprintf("%d", ttl))
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

	res, err := PostResponse(c.Client, u.String())
	if err != nil {
		return err
	}

	list := *res
	if list["status"] == "ERROR" {
		return errors.New((fmt.Sprintf("%v", list["message"])))
	}

	return nil
}

// EditSOA modifies a SOA (Start of Authority) record.
func (c *Client) EditSOA(domain, person string, refresh, retry, expire, ttl int64) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = apiDNSUpdateSOA
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("responsible-person", person)
	if refresh < 7200 {
		refresh = 7200
	}
	q.Set("refresh", fmt.Sprintf("%d", refresh))
	if retry < 7200 {
		retry = 7200
	}
	q.Set("retry", fmt.Sprintf("%d", retry))
	if expire < 172800 {
		expire = 172800
	}
	q.Set("expire", fmt.Sprintf("%d", expire))
	if ttl < 14400 {
		ttl = 14400
	}
	q.Set("ttl", fmt.Sprintf("%d", ttl))
	u.RawQuery = q.Encode()

	res, err := PostResponse(c.Client, u.String())
	if err != nil {
		return err
	}

	list := *res
	if list["status"] == "ERROR" {
		return errors.New((fmt.Sprintf("%v", list["message"])))
	}

	return nil
}
