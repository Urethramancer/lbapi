package lbapi

import (
	"fmt"
	"net/url"
)

type CustomerList struct {
	// Records returned in this structure.
	Records int64
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
		Records:    atoi(list["recsonpage"].(string)),
		MaxRecords: atoi(list["recsindb"].(string)),
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
