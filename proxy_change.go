package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

// ChangeARecord modifies A or AAAA records.
func (c *ProxyClient) ChangeARecord(domain, oldip, newip, host string, ttl int64, six bool) error {
	if six {
		return c.changeRecord("api/dns/manage/update-ipv6-record.json", domain, oldip, newip, host, ttl)
	}
	return c.changeRecord("api/dns/manage/update-ipv4-record.json", domain, oldip, newip, host, ttl)
}

func (c *ProxyClient) ChangeCNAME(domain, oldip, newip, host string, ttl int64) error {
	return c.changeRecord("api/dns/manage/update-cname-record.json", domain, oldip, newip, host, ttl)
}

func (c *ProxyClient) ChangeMX(domain, oldip, newip, host string, ttl int64, priority uint16) error {
	return c.changeRecordPri("api/dns/manage/update-mx-record.json", domain, oldip, newip, host, ttl, priority)
}

func (c *ProxyClient) ChangeNS(domain, oldip, newip, host string, ttl int64) error {
	return c.changeRecord("api/dns/manage/update-ns-record.json", domain, oldip, newip, host, ttl)
}

func (c *ProxyClient) ChangeTXT(domain, oldip, newip, host string, ttl int64) error {
	return c.changeRecord("api/dns/manage/update-txt-record.json", domain, oldip, newip, host, ttl)
}

func (c *ProxyClient) changeRecord(call, domain, oldip, newip, host string, ttl int64) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = call
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
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

func (c *ProxyClient) changeRecordPri(call, domain, oldip, newip, host string, ttl int64, priority uint16) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = call
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
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

// ChangeSRV modifies a SRV record.
func (c *ProxyClient) ChangeSRV(domain, oldval, newval, host string, ttl int64, priority, port, weight uint) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = "api/dns/manage/add-srv-record.json"
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
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

// ChangeSOA modifies a SOA (Start of Authority) record.
func (c *ProxyClient) ChangeSOA(domain, person string, refresh, retry, expire, ttl int64) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = "api/dns/manage/add-srv-record.json"
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
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
