package main

import (
	"fmt"

	"sort"

	"github.com/Urethramancer/lbapi"
	"github.com/ryanuber/columnize"
)

// DomainCmd holds the domain sub-commands.
type DomainCmd struct {
	List DomainListCmd `command:"list" description:"List all domains."`
	For  DomainForCmd  `command:"for" description:"List all domains belonging to a specified customer ID."`
}

type DomainListCmd struct {
	Verbose bool `short:"v" long:"verbose" description:"Show detailed information."`
}

func (cmd *DomainListCmd) Execute(args []string) error {
	return nil
}

type DomainForCmd struct {
	Verbose bool `short:"v" long:"verbose" description:"Show detailed information."`
	Short   bool `short:"s" long:"short" description:"List only the domain names, one per line. Overrides -v."`
	Args    struct {
		ID string `required:"true" positional-arg-name:"ID" description:"Customer ID."`
	} `positional-args:"true"`
}

func (cmd *DomainForCmd) Execute(args []string) error {
	var dl *lbapi.DomainList
	var err error
	var count int64
	var everything []*lbapi.Domain
	page := 1
	for {
		dl, err = client.DomainsFor(cmd.Args.ID, page)
		if err != nil {
			return err
		}

		everything = append(everything, dl.Domains...)
		page++
		count += dl.Records
		if count >= dl.MaxRecords {
			break
		}
	}

	if cmd.Short {
		for _, d := range dl.Domains {
			pr(d.Description)
		}
		return nil
	}

	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	if cmd.Verbose {
	} else {
		var s []string
		for _, d := range dl.Domains {
			s = append(s, fmt.Sprintf("%s\t%s\t%v", d.OrderID, d.Description, d.Endtime))
		}
		sort.Strings(s)
		s = append([]string{fmt.Sprintln("Order ID\tDomain\tExpires")}, s...)
		res := columnize.Format(s, cc)
		pr(res)
	}
	return nil
}
