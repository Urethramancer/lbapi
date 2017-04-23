package main

import (
	"fmt"
	"sort"

	"github.com/Urethramancer/countries"
	"github.com/Urethramancer/lbapi"
	"github.com/Urethramancer/lbapi/common"
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
		ID int64 `required:"true" positional-arg-name:"ID" description:"Customer ID."`
	} `positional-args:"true"`
}

// Execute the show command.
func (cmd *CustomerShowCmd) Execute(args []string) error {
	cust, err := client.CustomerByID(cmd.Args.ID)
	if err != nil {
		return err
	}

	var reseller string
	if cust.ParentReseller == cfg.ID {
		reseller = "you"
	}

	pr("%s (%d) - "+okColour(cust.Status == "Active")+"%s"+common.ANSI_NORMAL+"\nSigned up: %v", cust.Name, cust.ID, cust.Status, cust.Created)
	pr("E-mail: %s  Phone: %s", cust.Email, cust.Phone)

	pr(common.ANSI_YELLOW + "Address:" + common.ANSI_NORMAL)
	pr("%s\n", countries.FormatAddress(cust.Address, "", cust.Zip, cust.City, cust.State, cust.Country))
	pr("Two-factor enabled: "+okColour(cust.Twofactor)+"%v"+common.ANSI_NORMAL, cust.Twofactor)
	pr("Parent reseller: %d (%s)", cust.ParentReseller, reseller)
	pr("Total receipts: $%s", cust.TotalReceipts)
	return nil
}

// CustomerSearchCmd arguments.
type CustomerSearchCmd struct {
}

// Execute the search command.
func (cmd *CustomerSearchCmd) Execute(args []string) error {
	return nil
}
