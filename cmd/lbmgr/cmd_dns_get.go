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
	Args DNSArgs `positional-args:"true"`
}

// Execute DNS get command.
func (cmd *DNSGetCmd) Execute(args []string) error {
	t := getRecordType(cmd.Args.Type)
	if t == "" {
		return errors.New("unknown record type '" + cmd.Args.Type + "'")
	}

	var err error
	var count int64
	var everything lbapi.DNSRecords
	page := 1
	for {
		var recs *lbapi.DNSRecordList
		recs, err = client.GetDNSRecords(cmd.Args.Domain, t, page)
		if err != nil {
			return err
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
	s := []string{fmt.Sprintln("Host\tAddress\tType\tTTL\tStatus")}
	for _, r := range everything {
		s = append(s, fmt.Sprintf("%s\t%s\t%s\t%d\t%s", r.Host, r.Address, r.Type, r.TTL, r.Status))
	}
	res := columnize.Format(s, cc)
	pr(res)
	return nil
}
