package common

import (
	"errors"
	"fmt"
	"sort"

	"github.com/Urethramancer/lbapi"
	"github.com/ryanuber/columnize"
)

// DNSGetCmd arguments.
type DNSGetCmd struct {
	A     DNSGetACmd     `command:"a" description:"Get A records for a domain." alias:"A"`
	AAAA  DNSGetAAAACmd  `command:"aaaa" description:"Get AAAA records for a domain." alias:"AAAA"`
	CNAME DNSGetCNAMECmd `command:"cname" description:"Get CNAME records for a domain." alias:"CNAME"`
	MX    DNSGetMXCmd    `command:"mx" description:"Get MX records for a domain." alias:"MX"`
	NS    DNSGetNSCmd    `command:"ns" description:"Get NS records for a domain." alias:"NS"`
	TXT   DNSGetTXTCmd   `command:"txt" description:"Get TXT records for a domain." alias:"TXT"`
	SRV   DNSGetSRVCmd   `command:"srv" description:"Get SRV records for a domain." alias:"SRV"`
	All   DNSGetAllCmd   `command:"all" description:"Get all records for a domain."`
}

// DNSGetACmd arguments.
type DNSGetACmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute A record fetch.
func (cmd *DNSGetACmd) Execute(args []string) error {
	return printRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "A")
}

// DNSGetAAAACmd arguments.
type DNSGetAAAACmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute AAAA record fetch.
func (cmd *DNSGetAAAACmd) Execute(args []string) error {
	return printRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "AAAA")
}

// DNSGetCNAMECmd arguments.
type DNSGetCNAMECmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute CNAME record fetch.
func (cmd *DNSGetCNAMECmd) Execute(args []string) error {
	return printRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "CNAME")
}

// DNSGetMXCmd arguments.
type DNSGetMXCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute MX record fetch.
func (cmd *DNSGetMXCmd) Execute(args []string) error {
	return printRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "MX")
}

// DNSGetNSCmd arguments.
type DNSGetNSCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute NS record fetch.
func (cmd *DNSGetNSCmd) Execute(args []string) error {
	return printRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "NS")
}

// DNSGetTXTCmd arguments.
type DNSGetTXTCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute TXT record fetch.
func (cmd *DNSGetTXTCmd) Execute(args []string) error {
	return printRecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "TXT")
}

// Print the requested records of a type.
func printRecord(domain, value, host, t string) error {
	var err error

	dns, err := getRecords(domain, value, host, t)
	if err != nil {
		return err
	}

	sort.Sort(dns)
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
		s = []string{fmt.Sprintln("Host\tAddress\tTTL\tStatus\tPriority")}
	} else {
		s = []string{fmt.Sprintln("Host\tAddress\tTTL\tStatus")}
	}
	for _, r := range dns {
		if pri {
			s = append(s, fmt.Sprintf("%s\t%s\t%d\t%s\t%d", r.Host, r.Value, r.TTL, r.Status, r.Priority))
		} else {
			s = append(s, fmt.Sprintf("%s\t%s\t%d\t%s", r.Host, r.Value, r.TTL, r.Status))
		}
	}
	res := columnize.Format(s, cc)
	pr(res)
	return nil
}

func getRecords(domain, value, host, t string) (lbapi.DNSRecords, error) {
	var err error
	var records lbapi.DNSRecords
	page := 1
	var count int64
	for {
		var recs *lbapi.DNSRecordList
		recs, err = client.GetDNSRecords(domain, host, value, t, page)
		if err != nil {
			return nil, err
		}

		if recs.Count == 0 {
			return nil, errors.New("no records matching arguments")
		}

		records = append(records, recs.Records...)
		page++
		count += recs.Count
		if count >= recs.MaxRecords {
			break
		}
	}
	return records, nil
}

// DNSGetAllCmd arguments.
type DNSGetAllCmd struct {
	Args struct {
		Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	} `positional-args:"true"`
}

// Execute a complete record fetch for the domain.
func (cmd *DNSGetAllCmd) Execute(args []string) error {
	var a, aaaa, cname, mx, ns, txt, srv lbapi.DNSRecords
	var err error
	prn("Fetching A records…")
	a, err = getRecords(cmd.Args.Domain, "", "", "A")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(a)
	}

	prn("Fetching AAAA records…")
	aaaa, err = getRecords(cmd.Args.Domain, "", "", "AAAA")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(aaaa)
	}

	prn("Fetching CNAME records…")
	cname, err = getRecords(cmd.Args.Domain, "", "", "CNAME")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(cname)
	}

	prn("Fetching MX records…")
	mx, err = getRecords(cmd.Args.Domain, "", "", "MX")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(mx)
	}

	prn("Fetching NS records…")
	ns, err = getRecords(cmd.Args.Domain, "", "", "NS")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(ns)
	}

	prn("Fetching TXT records…")
	txt, err = getRecords(cmd.Args.Domain, "", "", "TXT")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(txt)
	}

	prn("Fetching SRV records…")
	srv, err = getRecords(cmd.Args.Domain, "", "", "SRV")
	if err != nil {
		pr("none found.")
	} else {
		pr("done.")
		sort.Sort(srv)
	}

	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	var s []string
	var res string
	base1 := []string{fmt.Sprintln("Host\tAddress\tTTL\tStatus")}
	base2 := []string{fmt.Sprintln("Host\tAddress\tTTL\tStatus\tPriority")}

	if len(a) > 0 {
		for _, r := range a {
			s = append(base1, fmt.Sprintf("%s\t%s\t%d\t%s", r.Host, r.Value, r.TTL, r.Status))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"A records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	if len(aaaa) > 0 {
		for _, r := range aaaa {
			s = append(base1, fmt.Sprintf("%s\t%s\t%d\t%s", r.Host, r.Value, r.TTL, r.Status))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"AAAA records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	if len(cname) > 0 {
		for _, r := range cname {
			s = append(base1, fmt.Sprintf("%s\t%s\t%d\t%s", r.Host, r.Value, r.TTL, r.Status))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"CNAME records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	if len(mx) > 0 {
		for _, r := range mx {
			s = append(base2, fmt.Sprintf("%s\t%s\t%d\t%s\t%d", r.Host, r.Value, r.TTL, r.Status, r.Priority))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"MX records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	if len(ns) > 0 {
		for _, r := range ns {
			s = append(base2, fmt.Sprintf("%s\t%s\t%d\t%s\t%d", r.Host, r.Value, r.TTL, r.Status, r.Priority))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"NS records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	if len(txt) > 0 {
		for _, r := range txt {
			s = append(base1, fmt.Sprintf("%s\t%s\t%d\t%s", r.Host, r.Value, r.TTL, r.Status))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"TXT records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	if len(txt) > 0 {
		for _, r := range txt {
			s = append(base1, fmt.Sprintf("%s\t%s\t%d\t%s", r.Host, r.Value, r.TTL, r.Status))
		}
		res = columnize.Format(s, cc)
		pr(ANSI_YELLOW+"TXT records:"+ANSI_NORMAL+"\n%s\n", res)
	}

	return nil
}

// DNSGetSRVCmd arguments.
type DNSGetSRVCmd struct {
	Args DNSGetArgs `positional-args:"true"`
}

// Execute SRV record fetch.
func (cmd *DNSGetSRVCmd) Execute(args []string) error {
	srv, err := getRecords(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, "SRV")
	if err != nil {
		return err
	}

	if len(srv) == 0 {
		return errors.New("no records found")
	}

	s := []string{fmt.Sprintln("Service name\tHost name\tPort\tWeight\tStatus")}
	for _, r := range srv {
		s = append(s, fmt.Sprintf("%s\t%s\t%d\t%d\t%s", r.Host, r.Value, r.Port, r.Weight, r.Status))
	}

	cc := columnize.DefaultConfig()
	cc.Delim = "\t"
	cc.Glue = "  "
	res := columnize.Format(s, cc)
	pr(res)
	return nil
}
