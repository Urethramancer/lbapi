package lbapi

import (
	"errors"
	"fmt"
	"net/url"
)

// DNSRecordList is what client software gets.
// It's not guaranteed to hold all records, so check Count against MaxRecords.
type DNSRecordList struct {
	// Count of records returned in this structure.
	Count int64
	// MaxRecords is the total available for this query.
	MaxRecords int64
	// Domains for the specified search query.
	Records DNSRecords
}

// DNSRecords is a special sortable structure.
type DNSRecords []*DNSRecord

// Len reports the number of records.
func (slice DNSRecords) Len() int {
	return len(slice)
}

// Less checks if host name i comes before host name j.
func (slice DNSRecords) Less(i, j int) bool {
	return slice[i].Host < slice[j].Host
}

// Swap does what it says.
func (slice DNSRecords) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// DNSRecord is an individual record.
type DNSRecord struct {
	TTL     int64  // 7200 is a safe default
	Host    string // subdomain or @ for the primary domain
	Type    string // A, AAAA, MX etc.
	Address string // IPv4 or IPv6 address
	Status  string // Normally "Active"
}

// DNSActive reports if an order has activated DNS yet.
// This is normally on by default, but will be activated when
// this is called otherwise.
func (c *Client) DNSActive(id string) bool {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return false
	}

	u.Path = "api/dns/activate.json"
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
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
func (c *Client) GetDNSRecords(domain, host, t string, page int) (*DNSRecordList, error) {
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
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("type", t)
	q.Set("no-of-records", "50")
	q.Set("page-no", fmt.Sprintf("%d", page))
	if host != "" {
		q.Set("host", host)
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

func parseDNS(in interface{}) *DNSRecord {
	data := in.(map[string]interface{})
	return &DNSRecord{
		Host:    data["host"].(string),
		Type:    data["type"].(string),
		Address: data["value"].(string),
		TTL:     atoi(data["timetolive"].(string)),
		Status:  data["status"].(string),
	}
}

func (c *Client) AddDNSA(domain, host, address string, ttl int64, six bool) error {
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
	if host != "" {
		q.Set("host", host)
	}
	q.Set("value", address)
	if ttl == 0 || ttl < 7200 {
		ttl = 7200
	}
	q.Set("ttl", fmt.Sprintf("%d", ttl))
	if host != "" {
		q.Set("host", host)
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

func (c *Client) ChangeDNSRecord(rec *DNSRecord) error {
	return nil
}

func (c *Client) DeleteDNSRecord(domain, host, address string, six bool) error {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	if six {
		u.Path = "api/dns/manage/delete-ipv6-record.json"
	} else {
		u.Path = "api/dns/manage/delete-ipv4-record.json"
	}
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("domain-name", domain)
	q.Set("value", address)
	if host != "" {
		q.Set("host", host)
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
