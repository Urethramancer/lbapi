package lbapi

import (
	"fmt"
	"net/url"
	"time"

	"github.com/Urethramancer/countries"
)

// CustomerList is what client software gets.
// It's not guaranteed to hold all records, so check Count against MaxRecords.
type CustomerList struct {
	// Count of records returned in this structure.
	Count int64
	// MaxRecords is the total available for this query.
	MaxRecords int64
	// Domains for the specified search query.
	Customers Customers
}

type Customers []*Customer

func (slice Customers) Len() int {
	return len(slice)
}

func (slice Customers) Less(i, j int) bool {
	return slice[i].ID < slice[j].ID
}

func (slice Customers) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Customer is the simple overview returned by bulk search of customers.
type Customer struct {
	ID            int64
	Email         string
	Name          string
	Company       string
	City          string
	Country       string
	Status        string
	TotalReceipts string

	Websites int64
}

// CustomerDetails contains the personal details returned by single-user fetch.
type CustomerDetails struct {
	Created        time.Time
	ID             int64
	ParentReseller int64

	Name          string
	Email         string
	Phone         string
	Address       string
	Zip           string
	City          string
	State         string
	Country       string
	Language      string
	PIN           string
	Status        string
	TotalReceipts string

	Twofactor bool
}

func (c *Client) CustomerByID(cid int64) (*CustomerDetails, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = API_CUSTOMERS_DETAILS_BY_ID
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
	q.Set("customer-id", fmt.Sprintf("%d", cid))
	u.RawQuery = q.Encode()

	res, err := c.getResponse(u.String())
	if err != nil {
		return nil, err
	}

	in := *res
	cust := parseCustomerDetails(in)
	return cust, nil
}

func parseCustomerDetails(in interface{}) *CustomerDetails {
	data := in.(maplist)
	phone := "+" + data["telnocc"].(string) + " " + data["telno"].(string)

	return &CustomerDetails{
		Created:        time.Unix(atoi(data["creationdt"].(string)), 0),
		ID:             atoi(data["customerid"].(string)),
		ParentReseller: atoi(data["resellerid"].(string)),
		Name:           data["name"].(string),
		Email:          data["useremail"].(string),
		Phone:          phone,
		Address:        data["address1"].(string),
		Zip:            data["zip"].(string),
		City:           data["city"].(string),
		State:          data["state"].(string),
		Country:        countries.Country(data["country"].(string)),
		Language:       data["langpref"].(string),
		PIN:            data["pin"].(string),
		Status:         data["customerstatus"].(string),
		TotalReceipts:  data["totalreceipts"].(string),
		Twofactor:      parseBool(data["twofactorauth_enabled"].(string)),
	}
}

func (c *Client) Customers(page int) (*CustomerList, error) {
	var err error
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = "api/customers/search.json"
	q := u.Query()
	q.Set("auth-userid", c.ID)
	q.Set("api-key", c.Key)
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
	cl := CustomerList{
		Count:      atoi(fmt.Sprintf("%v", list["recsonpage"])),
		MaxRecords: atoi(fmt.Sprintf("%v", list["recsindb"])),
	}
	delete(list, "recsonpage")
	delete(list, "recsindb")

	for _, customers := range list {
		cust := parseCustomer(customers)
		if cust != nil {
			cl.Customers = append(cl.Customers, cust)
		}
	}
	return &cl, nil
}

func parseCustomer(in interface{}) *Customer {
	data := in.(map[string]interface{})
	return &Customer{
		ID:            atoi(data["customer.customerid"].(string)),
		Email:         data["customer.username"].(string),
		Name:          data["customer.name"].(string),
		Company:       data["customer.company"].(string),
		City:          data["customer.city"].(string),
		Country:       data["customer.country"].(string),
		Status:        data["customer.customerstatus"].(string),
		TotalReceipts: data["customer.totalreceipts"].(string),
		Websites:      atoi(data["customer.websitecount"].(string)),
	}
}
