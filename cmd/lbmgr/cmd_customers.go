package main

import (
	"fmt"
	"sort"

	"github.com/Urethramancer/lbapi"
	"github.com/ryanuber/columnize"
)

// CustomerCmd is a place to hang the Execute command.
type CustomerCmd struct {
	List   CustomerListCmd   `command:"list" description:"List all customers."`
	Show   CustomerShowCmd   `command:"show" description:"Show details for one specific customer ID."`
	Search CustomerSearchCmd `command:"search" description:"Search for customers."`
}

// CustomerListCmd arguments.
type CustomerListCmd struct {
	Verbose bool `short:"v" long:"verbose" description:"Show detailed information."`
}

// Execute the list command.
func (cmd *CustomerListCmd) Execute(args []string) error {
	var cl *lbapi.CustomerList
	var err error
	var count int64
	page := 1
	var everybody lbapi.Customers
	for {
		cl, err = client.Customers(page)
		if err != nil {
			return err
		}

		everybody = append(everybody, cl.Customers...)
		page++
		count += cl.Count
		if count >= cl.MaxRecords {
			break
		}
	}

	sort.Sort(everybody)
	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	var s []string
	if cmd.Verbose {
		s = []string{fmt.Sprintln("ID\tName\tE-mail (username)\tCompany\tCity\tCountry\tStatus\tReceipts\tWebsites")}
		for _, c := range everybody {
			s = append(s, fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\n", c.ID, c.Name, c.Email, c.Company, c.City, c.Country, c.Status, c.TotalReceipts, c.Websites))
		}
		res := columnize.Format(s, cc)
		pr(res)
	} else {
		s = []string{fmt.Sprintln("ID\tName\tE-mail (username)\tStatus")}
		for _, c := range everybody {
			s = append(s, fmt.Sprintf("%d\t%s\t%s\t%s\n", c.ID, c.Name, c.Email, c.Status))
		}
		res := columnize.Format(s, cc)
		pr(res)
	}
	return nil
}

// CustomerShowCmd arguments.
type CustomerShowCmd struct {
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"Customer ID."`
	} `positional-args:"true"`
}

// Execute the show command.
func (cmd *CustomerShowCmd) Execute(args []string) error {
	cust := client.CustomerByID
	return nil
}

// CustomerSearchCmd arguments.
type CustomerSearchCmd struct {
}

// Execute the search command.
func (cmd *CustomerSearchCmd) Execute(args []string) error {
	return nil
}
