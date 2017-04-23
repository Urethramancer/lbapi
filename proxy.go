package lbapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ProxyClient is the structure used by end-user software to access lbproxy middleman servers.
type ProxyClient struct {
	http.Client
	URL      string
	Token    string
	Username string
	Password string
}

// NewProxyClient creates a client structure with a HTTP client for the specified proxy, and logs in.
func NewProxyClient(api, username, password string) *ProxyClient {
	if api == "" {
		api = "http://localhost:11000"
	}

	return &ProxyClient{
		Client:   http.Client{Timeout: time.Second * 30},
		URL:      api,
		Token:    "",
		Username: username,
		Password: password,
	}
}

// getReponse is the internal call to fetch a URL's JSON,
// decode it into a maplist and hand it over to massage into
// a proper structure.
func (c *ProxyClient) getResponse(url string) (*maplist, error) {
	res, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var list maplist
	err = decoder.Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// postResponse does pretty much the same as getResponse().
// Some LogicBoxes API calls use POST instead, but there doesn't
// seem to be any actual logic to why, as they aren't actually
// posting any body content.
func (c *ProxyClient) postResponse(url string) (*maplist, error) {
	res, err := c.Client.Post(url, "", nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var list maplist
	err = decoder.Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// AddARecord adds A or AAAA records.
func (c *ProxyClient) AddARecord(domain, address, host string, ttl int64, six bool) error {
	if six {
		return c.addRecord("api/dns/manage/add-ipv6-record.json", domain, address, host, ttl)
	}
	return c.addRecord("api/dns/manage/add-ipv4-record.json", domain, address, host, ttl)
}

// AddCNAME does exactly that.
func (c *ProxyClient) AddCNAME(domain, value, host string, ttl int64) error {
	return c.addRecord("api/dns/manage/add-cname-record.json", domain, value, host, ttl)
}

// AddMX adds MX records for mail servers.
func (c *ProxyClient) AddMX(domain, value, host string, ttl int64, priority uint16) error {
	return c.addRecordPri("api/dns/manage/add-mx-record.json", domain, value, host, ttl, priority)
}

// AddNS adds name server records.
func (c *ProxyClient) AddNS(domain, value, host string, ttl int64, priority uint16) error {
	return c.addRecord("api/dns/manage/add-ns-record.json", domain, value, host, ttl)
}

// AddTXT adds TXT records.
func (c *ProxyClient) AddTXT(domain, value, host string, ttl int64, priority uint16) error {
	return c.addRecord("api/dns/manage/add-txt-record.json", domain, value, host, ttl)
}

// AddSRV adds SRV records.
func (c *ProxyClient) AddSRV(domain, value, host string, ttl int64, priority, port, weight uint16) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	u.Path = "dns/manage/add-srv-record.json"
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", value)
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
	if list["status"] != "Success" {
		return errors.New("couldn't add SRV record - check that FQDN is correct")
	}

	return nil
}

func (c *ProxyClient) addRecord(call, domain, address, host string, ttl int64) error {
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

func (c *ProxyClient) addRecordPri(call, domain, address, host string, ttl int64, priority uint16) error {
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
