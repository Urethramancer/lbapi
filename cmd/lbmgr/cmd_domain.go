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
	var everything lbapi.Domains
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
		for _, d := range everything {
			pr(d.Description)
		}
		return nil
	}

	sort.Sort(everything)
	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	var s []string
	if cmd.Verbose {
	} else {
		s = []string{fmt.Sprintln("Order ID\tDomain\tExpires")}
		for _, d := range everything {
			s = append(s, fmt.Sprintf("%d\t%s\t%v", d.OrderID, d.Description, d.Endtime))
		}
		res := columnize.Format(s, cc)
		pr(res)
	}
	return nil
}
