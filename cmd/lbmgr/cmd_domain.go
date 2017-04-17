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
	Show DomainShowCmd `command:"show" description:"Show details for one domain."`
	For  DomainForCmd  `command:"for" description:"List all domains belonging to a specified customer ID."`
}

// DomainListCmd arguments.
type DomainListCmd struct {
}

// Execute the list command.
func (cmd *DomainListCmd) Execute(args []string) error {
	return nil
}

// DomainForCmd arguments.
type DomainForCmd struct {
	Verbose bool `short:"v" long:"verbose" description:"Show detailed information."`
	Short   bool `short:"s" long:"short" description:"List only the domain names, one per line. Overrides -v."`
	Name    bool `short:"n" long:"name" description:"Sort by name instead of ID."`
	Args    struct {
		ID string `required:"true" positional-arg-name:"ID" description:"Customer ID."`
	} `positional-args:"true"`
}

// Execute the "show domains for" command.
func (cmd *DomainForCmd) Execute(args []string) error {
	var dl *lbapi.DomainList
	var err error
	var count int64
	var everything lbapi.Domains
	var everyname lbapi.DomainsByName
	page := 1
	for {
		dl, err = client.DomainsFor(cmd.Args.ID, page)
		if err != nil {
			return err
		}

		if cmd.Name || cmd.Short {
			everyname = append(everyname, dl.Domains...)
		} else {
			everything = append(everything, dl.Domains...)
		}
		page++
		count += dl.Count
		if count >= dl.MaxRecords {
			break
		}
	}

	if cmd.Name || cmd.Short {
		sort.Sort(everyname)
	} else {
		sort.Sort(everything)
	}

	if cmd.Short {
		for _, d := range everyname {
			pr(d.Description)
		}
		return nil
	}

	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	var s []string
	if cmd.Verbose {
	} else {
		if cmd.Name {
			s = []string{fmt.Sprintln("Domain\tOrder ID\tExpires")}
			for _, d := range everyname {
				s = append(s, fmt.Sprintf("%s\t%d\t%v", d.Description, d.OrderID, d.Endtime))
			}
		} else {
			s = []string{fmt.Sprintln("Order ID\tDomain\tExpires")}
			for _, d := range everything {
				s = append(s, fmt.Sprintf("%d\t%s\t%v", d.OrderID, d.Description, d.Endtime))
			}
		}
		res := columnize.Format(s, cc)
		pr(res)
	}
	return nil
}
