package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Urethramancer/lbapi"
)

type maplist map[string]interface{}

// Client is the structure used by end-user software to access lbproxy middleman servers.
type Client struct {
	http.Client
	URL      string
	Token    string
	Username string
	Password string
}

// NewClient creates a client structure with a HTTP client for the specified proxy.
func NewClient(api, username, password string) *Client {
	if api == "" {
		api = "http://localhost:11000"
	}

	c := Client{
		Client:   http.Client{Timeout: time.Second * 30},
		URL:      api,
		Token:    "",
		Username: username,
		Password: password,
	}

	if !c.Authenticate() {
		return nil
	}

	return &c
}

func (c *Client) DNSActive(string) bool {
	return true
}

// GetDNSRecords gets the first up to 50 records of one type for a domain.
// Pass a higher page number to get the next set of up to 50.
func (c *Client) GetDNSRecords(domain, value, host, t string, page int) (*lbapi.DNSRecordList, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	if page == 0 {
		page = 1
	}

	u.Path = PathDNSGet
	q := u.Query()
	q.Set("token", c.Token)
	q.Set("domain", domain)
	q.Set("type", t)
	q.Set("count", "50")
	q.Set("page", fmt.Sprintf("%d", page))
	if host != "" {
		q.Set("host", host)
	}
	if value != "" {
		q.Set("value", value)
	}
	u.RawQuery = q.Encode()

	res, err := lbapi.GetResponse(c.Client, u.String())
	if err != nil {
		return nil, err
	}

	list := *res
	rl := lbapi.DNSRecordList{
		Count:      lbapi.Atoi(fmt.Sprintf("%v", list["recsonpage"])),
		MaxRecords: lbapi.Atoi(fmt.Sprintf("%v", list["recsindb"])),
	}
	delete(list, "recsonpage")
	delete(list, "recsindb")

	for _, rec := range list {
		r := lbapi.ParseDNS(rec)
		if r != nil {
			rl.Records = append(rl.Records, r)
		}
	}
	return &rl, nil
}
