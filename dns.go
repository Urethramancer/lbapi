package lbapi

import "net/url"
import "fmt"

type DNSRecordList struct {
	// Count of records returned in this structure.
	Count int64
	// MaxRecords is the total available for this query.
	MaxRecords int64
	// Domains for the specified search query.
	Records DNSRecords
}

type DNSRecords []*DNSRecord

func (slice DNSRecords) Len() int {
	return len(slice)
}

func (slice DNSRecords) Less(i, j int) bool {
	return slice[i].Host < slice[j].Host
}

func (slice DNSRecords) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type DNSRecord struct {
	TTL     int64
	Host    string
	Type    string
	Address string
	Status  string
}

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

func (c *Client) GetDNSRecords(domain, t string, page int) (*DNSRecordList, error) {
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
