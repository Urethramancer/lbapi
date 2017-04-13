package lbapi

import (
	"fmt"
	"net/url"
	"time"
)

type DomainList struct {
	// Count of records returned in this structure.
	Count int64
	// MaxRecords can be used to calculate pagination.
	MaxRecords int64
	// Domains for the specified search query.
	Domains Domains
}

type Domains []*Domain
type DomainsByName []*Domain

func (slice Domains) Len() int {
	return len(slice)
}

func (slice Domains) Less(i, j int) bool {
	return slice[i].OrderID < slice[j].OrderID
}

func (slice Domains) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice DomainsByName) Len() int {
	return len(slice)
}

func (slice DomainsByName) Less(i, j int) bool {
	return slice[i].Description < slice[j].Description
}

func (slice DomainsByName) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type Domain struct {
	Endtime      time.Time
	CreationTime time.Time
	CreationDT   time.Time
	Timestamp    time.Time

	OrderID     int64
	CustomerID  int64
	EntityID    int64
	TypeID      int64
	Description string
	Status      string
	TypeKey     string
	TypeName    string

	Autorenew    bool
	ResellerLock bool
	CustomerLock bool
	TransferLock bool
}

// DomainsFor customer, starting on a specified page.
// Up to 500 records are returned. Compare Records and MaxRecords to tell
// if another page exists.
func (c *Client) DomainsFor(customer string, page int) (*DomainList, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = "api/domains/search.json"
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("customer-id", customer)
	q.Set("no-of-records", "500")
	if page == 0 {
		page = 1
	}
	q.Set("page-no", fmt.Sprintf("%d", page))
	u.RawQuery = q.Encode()

	res, err := c.getResponse(u.String())
	if err != nil {
		return nil, err
	}

	list := *res
	dl := DomainList{
		Count:      atoi(fmt.Sprintf("%v", list["recsonpage"])),
		MaxRecords: atoi(fmt.Sprintf("%v", list["recsindb"])),
	}
	delete(list, "recsonpage")
	delete(list, "recsindb")

	for _, dom := range list {
		d := parseDomain(dom)
		if d != nil {
			dl.Domains = append(dl.Domains, d)
		}
	}
	return &dl, nil
}

func parseDomain(in interface{}) *Domain {
	data := in.(map[string]interface{})
	return &Domain{
		Endtime:      time.Unix(atoi(data["orders.endtime"].(string)), 0),
		CreationTime: time.Unix(atoi(data["orders.creationtime"].(string)), 0),
		CreationDT:   time.Unix(atoi(data["orders.creationdt"].(string)), 0),
		Timestamp:    parseDate(data["orders.timestamp"].(string)),
		OrderID:      atoi(data["orders.orderid"].(string)),
		CustomerID:   atoi(data["entity.customerid"].(string)),
		EntityID:     atoi(data["entity.entityid"].(string)),
		TypeID:       atoi(data["entity.entitytypeid"].(string)),
		Description:  data["entity.description"].(string),
		Status:       data["entity.currentstatus"].(string),
		TypeKey:      data["entitytype.entitytypekey"].(string),
		TypeName:     data["entitytype.entitytypename"].(string),
		Autorenew:    parseBool(data["orders.autorenew"].(string)),
		ResellerLock: parseBool(data["orders.resellerlock"].(string)),
		CustomerLock: parseBool(data["orders.customerlock"].(string)),
		TransferLock: parseBool(data["orders.transferlock"].(string)),
	}
}
