package main

// DNSChangeCmd arguments.
type DNSChangeCmd struct {
	A     DNSChangeACmd     `command:"a" description:"Modify an A record." alias:"A"`
	AAAA  DNSChangeAAAACmd  `command:"aaaa" description:"Modify an AAAA record." alias:"AAAA"`
	CNAME DNSChangeCNAMECmd `command:"cname" description:"Modify a CNAME record." alias:"CNAME"`
	MX    DNSChangeMXCmd    `command:"mx" description:"Modify an MX record." alias:"MX"`
	NS    DNSChangeNSCmd    `command:"ns" description:"Modify an NS record." alias:"NS"`
	TXT   DNSChangeTXTCmd   `command:"txt" description:"Modify a TXT record." alias:"TXT"`
	SRV   DNSChangeSRVCmd   `command:"srv" description:"Modify a SRV record." alias:"SRV"`
	SOA   DNSChangeSOACmd   `command:"soa" description:"Modify the SOA record." alias:"SOA"`
}

// DNSChangeACmd arguments.
type DNSChangeACmd struct {
	Args DNSChangeArgs `positional-args:"true"`
}

// Execute A record modification.
func (cmd *DNSChangeACmd) Execute(args []string) error {
	err := client.ChangeARecord(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, false)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSChangeAAAACmd arguments.
type DNSChangeAAAACmd struct {
	Args DNSChangeArgs `positional-args:"true"`
}

// Execute AAAA record modification.
func (cmd *DNSChangeAAAACmd) Execute(args []string) error {
	err := client.ChangeARecord(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, true)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSChangeCNAMECmd arguments.
type DNSChangeCNAMECmd struct {
	Args DNSChangeArgs `positional-args:"true"`
}

// Execute CNAME record modification.
func (cmd *DNSChangeCNAMECmd) Execute(args []string) error {
	err := client.ChangeCNAME(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSChangeMXCmd arguments.
type DNSChangeMXCmd struct {
	Args DNSChangeArgsPri `positional-args:"true"`
}

// Execute MX record modification.
func (cmd *DNSChangeMXCmd) Execute(args []string) error {
	err := client.ChangeMX(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSChangeNSCmd arguments.
type DNSChangeNSCmd struct {
	Args DNSChangeArgs `positional-args:"true"`
}

// Execute NS record modification.
func (cmd *DNSChangeNSCmd) Execute(args []string) error {
	err := client.ChangeNS(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSChangeTXTCmd arguments.
type DNSChangeTXTCmd struct {
	Args DNSChangeArgs `positional-args:"true"`
}

// Execute TXT record modification.
func (cmd *DNSChangeTXTCmd) Execute(args []string) error {
	err := client.ChangeTXT(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSChangeSRVCmd arguments.
type DNSChangeSRVCmd struct {
	Args struct {
		Domain   string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
		Old      string `required:"true" positional-arg-name:"OLDIP" description:"IP address to change from."`
		New      string `required:"true" positional-arg-name:"NEWIP" description:"New IP address."`
		Host     string `required:"true" positional-arg-name:"HOST" description:"A fully qualified service name in the form '_<service name>._<protocol>.<domain.tld>'."`
		TTL      int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
		Priority uint   `positional-arg-name:"PRIORITY" description:"Priority of record. Default is 0 (most important)."`
		Port     uint   `positional-arg-name:"PORT" description:"Port number of the service."`
		Weight   uint   `positional-arg-name:"WEIGHT" description:"Relative weight for records with the same priority."`
	} `positional-args:"true"`
}

// Execute SRV record modification.
func (cmd *DNSChangeSRVCmd) Execute(args []string) error {
	return client.ChangeSRV(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority, cmd.Args.Port, cmd.Args.Weight)
}

// DNSChangeSOACmd arguments.
type DNSChangeSOACmd struct {
	Args struct {
		Domain  string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
		Person  string `required:"true" positional-arg-name:"PERSON" description:"Responsible person's e-mail."`
		Refresh int64  `positional-arg-name:"REFRESH" description:"Seconds after which the secondary DNS server checks the primary DNS server for zone changes. Minimum 7200 (2 hours)."`
		Retry   int64  `positional-arg-name:"RETRY" description:"Seconds between retries for failed refreshes. Minimum 7200 (2 hours)."`
		Expire  int64  `positional-arg-name:"EXPIRE" description:"Upper limit in seconds before a zone is no longer authoritative. Minimum 172800 (48 hours)."`
		TTL     int64  `positional-arg-name:"TTL" description:"Time to live, or the number of seconds the record needs to be cached by DNS servers. Minimum 14400 (4 hours)."`
	}
}

// Execute SOA (Start of Authority) record modification.
func (cmd *DNSChangeSOACmd) Execute(args []string) error {
	return client.ChangeSOA(cmd.Args.Domain, cmd.Args.Person, cmd.Args.Refresh, cmd.Args.Retry, cmd.Args.Expire, cmd.Args.TTL)
}
