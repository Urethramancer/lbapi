package lbapi

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// DomainList is what client software gets.
// It's not guaranteed to hold all records, so check Count against MaxRecords.
type DomainList struct {
	// Count of records returned in this structure.
	Count int64
	// MaxRecords can be used to calculate pagination.
	MaxRecords int64
	// Domains for the specified search query.
	Domains Domains
}

// Domains is a sortable list of domains by order ID.
type Domains []*Domain

// DomainsByName uses the name instead.
type DomainsByName []*Domain

// Len is the number of domains.
func (slice Domains) Len() int {
	return len(slice)
}

// Less compares order IDs for sorting.
func (slice Domains) Less(i, j int) bool {
	return slice[i].OrderID < slice[j].OrderID
}

// Swap does just that.
func (slice Domains) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Len is the number of domains.
func (slice DomainsByName) Len() int {
	return len(slice)
}

// Less compares domain names for sorting.
func (slice DomainsByName) Less(i, j int) bool {
	return slice[i].Description < slice[j].Description
}

// Swap does just that.
func (slice DomainsByName) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Domain is the parsed structure from the messy JSON returned by LogicBoxes.
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

func (c *Client) Domain(name string) (*Domain, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = API_DOMAINS_SEARCH
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("no-of-records", "10")
	q.Set("page-no", "1")
	q.Set("domain-name", name)
	u.RawQuery = q.Encode()

	res, err := GetResponse(c.Client, u.String())
	if err != nil {
		return nil, err
	}

	list := *res
	in, ok := list["1"]
	if !ok {
		return nil, errors.New("no domain matching that name.")
	}
	domain := parseDomain(in)
	return domain, nil
}

// DomainsFor customer, starting on a specified page.
// Up to 500 records are returned. Compare Count and MaxRecords to tell
// if another page exists.
func (c *Client) DomainsFor(customer string, page int) (*DomainList, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = API_DOMAINS_SEARCH
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

	res, err := GetResponse(c.Client, u.String())
	if err != nil {
		return nil, err
	}

	list := *res
	dl := DomainList{
		Count:      Atoi(fmt.Sprintf("%v", list["recsonpage"])),
		MaxRecords: Atoi(fmt.Sprintf("%v", list["recsindb"])),
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
		Endtime:      time.Unix(Atoi(data["orders.endtime"].(string)), 0),
		CreationTime: time.Unix(Atoi(data["orders.creationtime"].(string)), 0),
		CreationDT:   time.Unix(Atoi(data["orders.creationdt"].(string)), 0),
		Timestamp:    ParseDate(data["orders.timestamp"].(string)),
		OrderID:      Atoi(data["orders.orderid"].(string)),
		CustomerID:   Atoi(data["entity.customerid"].(string)),
		EntityID:     Atoi(data["entity.entityid"].(string)),
		TypeID:       Atoi(data["entity.entitytypeid"].(string)),
		Description:  data["entity.description"].(string),
		Status:       data["entity.currentstatus"].(string),
		TypeKey:      data["entitytype.entitytypekey"].(string),
		TypeName:     data["entitytype.entitytypename"].(string),
		Autorenew:    ParseBool(data["orders.autorenew"].(string)),
		ResellerLock: ParseBool(data["orders.resellerlock"].(string)),
		CustomerLock: ParseBool(data["orders.customerlock"].(string)),
		TransferLock: ParseBool(data["orders.transferlock"].(string)),
	}
}
