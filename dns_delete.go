package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

// DeleteARecord deletes A or AAAA records.
func (c *Client) DeleteARecord(domain, value, host string, six bool) error {
	if six {
		return c.deleteRecord(APIDNSDeleteIPv6, domain, value, host)
	}
	return c.deleteRecord(APIDNSDeleteIPv4, domain, value, host)
}

// DeleteCNAME does exactly that.
func (c *Client) DeleteCNAME(domain, value, host string) error {
	return c.deleteRecord(APIDNSDeleteCNAME, domain, value, host)
}

// DeleteMX holds no surprises.
func (c *Client) DeleteMX(domain, value, host string) error {
	return c.deleteRecord(APIDNSDeleteMX, domain, value, host)
}

// DeleteNS is as boring as the above.
func (c *Client) DeleteNS(domain, value, host string) error {
	return c.deleteRecord(APIDNSDeleteNS, domain, value, host)
}

// DeleteTXT deletes TXT records.
func (c *Client) DeleteTXT(domain, value, host string) error {
	return c.deleteRecord(APIDNSDeleteTXT, domain, value, host)
}

// DeleteSRV deletes SRV records.
func (c *Client) DeleteSRV(domain, value, host string, port, weight uint16) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = APIDNSDeleteSRV
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", value)
	q.Set("host", host)
	q.Set("port", fmt.Sprintf("%d", port))
	q.Set("weight", fmt.Sprintf("%d", weight))
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

func (c *Client) deleteRecord(call, domain, value, host string) error {
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
	q.Set("value", value)
	q.Set("host", host)
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
