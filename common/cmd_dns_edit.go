package common

// DNSEditCmd arguments.
type DNSEditCmd struct {
	A     DNSEditACmd     `command:"a" description:"Edit an A record." alias:"A"`
	AAAA  DNSEditAAAACmd  `command:"aaaa" description:"Edit an AAAA record." alias:"AAAA"`
	CNAME DNSEditCNAMECmd `command:"cname" description:"Edit a CNAME record." alias:"CNAME"`
	MX    DNSEditMXCmd    `command:"mx" description:"Edit an MX record." alias:"MX"`
	NS    DNSEditNSCmd    `command:"ns" description:"Edit an NS record." alias:"NS"`
	TXT   DNSEditTXTCmd   `command:"txt" description:"Edit a TXT record." alias:"TXT"`
	SRV   DNSEditSRVCmd   `command:"srv" description:"Edit a SRV record." alias:"SRV"`
	SOA   DNSEditSOACmd   `command:"soa" description:"Edit the SOA record." alias:"SOA"`
}

// DNSEditACmd arguments.
type DNSEditACmd struct {
	Args DNSEditArgs `positional-args:"true"`
}

// Execute A record modification.
func (cmd *DNSEditACmd) Execute(args []string) error {
	err := client.EditARecord(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, false)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSEditAAAACmd arguments.
type DNSEditAAAACmd struct {
	Args DNSEditArgs `positional-args:"true"`
}

// Execute AAAA record modification.
func (cmd *DNSEditAAAACmd) Execute(args []string) error {
	err := client.EditARecord(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, true)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSEditCNAMECmd arguments.
type DNSEditCNAMECmd struct {
	Args DNSEditArgs `positional-args:"true"`
}

// Execute CNAME record modification.
func (cmd *DNSEditCNAMECmd) Execute(args []string) error {
	err := client.EditCNAME(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSEditMXCmd arguments.
type DNSEditMXCmd struct {
	Args DNSEditArgsPri `positional-args:"true"`
}

// Execute MX record modification.
func (cmd *DNSEditMXCmd) Execute(args []string) error {
	err := client.EditMX(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSEditNSCmd arguments.
type DNSEditNSCmd struct {
	Args DNSEditArgs `positional-args:"true"`
}

// Execute NS record modification.
func (cmd *DNSEditNSCmd) Execute(args []string) error {
	err := client.EditNS(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSEditTXTCmd arguments.
type DNSEditTXTCmd struct {
	Args DNSEditArgs `positional-args:"true"`
}

// Execute TXT record modification.
func (cmd *DNSEditTXTCmd) Execute(args []string) error {
	err := client.EditTXT(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSEditSRVCmd arguments.
type DNSEditSRVCmd struct {
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
func (cmd *DNSEditSRVCmd) Execute(args []string) error {
	return client.EditSRV(cmd.Args.Domain, cmd.Args.Old, cmd.Args.New, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority, cmd.Args.Port, cmd.Args.Weight)
}

// DNSEditSOACmd arguments.
type DNSEditSOACmd struct {
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
func (cmd *DNSEditSOACmd) Execute(args []string) error {
	return client.EditSOA(cmd.Args.Domain, cmd.Args.Person, cmd.Args.Refresh, cmd.Args.Retry, cmd.Args.Expire, cmd.Args.TTL)
}
