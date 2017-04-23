package lbapi

import (
	"encoding/json"
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

// DNSActive reports if an order has activated DNS yet.
// This is normally on by default, but will be activated when
// this is called otherwise.
func (c *ProxyClient) DNSActive(id string) bool {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return false
	}

	u.Path = "api/dns/activate.json"
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
	q.Set("order-id", id)
	u.RawQuery = q.Encode()

	res, err := c.postResponse(u.String())
	if err != nil {
		return false
	}

	list := *res
	return list["status"] == "Success"
}

// GetDNSRecords gets the first up to 50 records of one type for a domain.
// Pass a higher page number to get the next set of up to 50.
func (c *ProxyClient) GetDNSRecords(domain, value, host, t string, page int) (*DNSRecordList, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	if page == 0 {
		page = 1
	}

	u.Path = "api/dns/manage/search-records.json"
	q := u.Query()
	// q.Set("auth-userid", c.ID)
	// q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("type", t)
	q.Set("no-of-records", "50")
	q.Set("page-no", fmt.Sprintf("%d", page))
	if host != "" {
		q.Set("host", host)
	}
	if value != "" {
		q.Set("value", value)
	}
	u.RawQuery = q.Encode()

	res, err := c.getResponse(u.String())
	if err != nil {
		return nil, err
	}

	list := *res
	rl := DNSRecordList{
		Count:      atoi(fmt.Sprintf("%v", list["recsonpage"])),
		MaxRecords: atoi(fmt.Sprintf("%v", list["recsindb"])),
	}
	delete(list, "recsonpage")
	delete(list, "recsindb")

	for _, rec := range list {
		r := parseDNS(rec)
		if r != nil {
			rl.Records = append(rl.Records, r)
		}
	}
	return &rl, nil
}
