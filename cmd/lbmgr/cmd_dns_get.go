package main

import (
	"errors"
	"fmt"
	"sort"

	"github.com/Urethramancer/lbapi"
	"github.com/ryanuber/columnize"
)

// DNSGetCmd arguments.
type DNSGetCmd struct {
	A     DNSGetACmd     `command:"a" description:"Add an A record." alias:"A"`
	AAAA  DNSGetAAAACmd  `command:"aaaa" description:"Get an AAAA record." alias:"AAAA"`
	CNAME DNSGetCNAMECmd `command:"cname" description:"Get a CNAME record." alias:"CNAME"`
	MX    DNSGetMXCmd    `command:"mx" description:"Get an MX record." alias:"MX"`
	NS    DNSGetNSCmd    `command:"ns" description:"Get an NS record." alias:"NS"`
	TXT   DNSGetTXTCmd   `command:"txt" description:"Get a TXT record." alias:"TXT"`
	SRV   DNSGetSRVCmd   `command:"srv" description:"Get a SRV record." alias:"SRV"`
}

// DNSGetACmd arguments.
type DNSGetACmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute A record fetch.
func (cmd *DNSGetACmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "A")
}

// DNSGetAAAACmd arguments.
type DNSGetAAAACmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute AAAA record fetch.
func (cmd *DNSGetAAAACmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "AAAA")
}

// DNSGetCNAMECmd arguments.
type DNSGetCNAMECmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute CNAME record fetch.
func (cmd *DNSGetCNAMECmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "CNAME")
}

// DNSGetMXCmd arguments.
type DNSGetMXCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute MX record fetch.
func (cmd *DNSGetMXCmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "MX")
}

// DNSGetNSCmd arguments.
type DNSGetNSCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute NS record fetch.
func (cmd *DNSGetNSCmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "NS")
}

// DNSGetTXTCmd arguments.
type DNSGetTXTCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute TXT record fetch.
func (cmd *DNSGetTXTCmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "TXT")
}

// DNSGetSRVCmd arguments.
type DNSGetSRVCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute SRV record fetch.
func (cmd *DNSGetSRVCmd) Execute(args []string) error {
	return getRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "SRV")
}

// Execute DNS get command.
func getRecord(domain, value, host, t string) error {
	var err error
	var count int64
	var everything lbapi.DNSRecords
	page := 1
	for {
		var recs *lbapi.DNSRecordList
		recs, err = client.GetDNSRecords(domain, host, value, t, page)
		if err != nil {
			return err
		}

		if recs.Count == 0 {
			return errors.New("no records matching arguments")
		}

		everything = append(everything, recs.Records...)
		page++
		count += recs.Count
		if count >= recs.MaxRecords {
			break
		}
	}

	sort.Sort(everything)
	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	var s []string
	var pri bool
	switch t {
	case "mx", "MX", "srv", "SRV":
		pri = true
	}
	if pri {
		s = []string{fmt.Sprintln("Host\tAddress\tType\tTTL\tPriority\tStatus")}
	} else {
		s = []string{fmt.Sprintln("Host\tAddress\tType\tTTL\tStatus")}
	}
	for _, r := range everything {
		if pri {
			s = append(s, fmt.Sprintf("%s\t%s\t%s\t%d\t%d\t%s", r.Host, r.Address, r.Type, r.TTL, r.Priority, r.Status))
		} else {
			s = append(s, fmt.Sprintf("%s\t%s\t%s\t%d\t%s", r.Host, r.Address, r.Type, r.TTL, r.Status))
		}
	}
	res := columnize.Format(s, cc)
	pr(res)
	return nil
}
