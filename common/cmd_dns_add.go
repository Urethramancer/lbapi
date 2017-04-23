package common

// DNSAddCmd arguments.
type DNSAddCmd struct {
	A     DNSAddACmd     `command:"a" description:"Add an A record." alias:"A"`
	AAAA  DNSAddAAAACmd  `command:"aaaa" description:"Add an AAAA record." alias:"AAAA"`
	CNAME DNSAddCNAMECmd `command:"cname" description:"Add a CNAME record." alias:"CNAME"`
	MX    DNSAddMXCmd    `command:"mx" description:"Add an MX record." alias:"MX"`
	NS    DNSAddNSCmd    `command:"ns" description:"Add an NS record." alias:"NS"`
	TXT   DNSAddTXTCmd   `command:"txt" description:"Add a TXT record." alias:"TXT"`
	SRV   DNSAddSRVCmd   `command:"srv" description:"Add a SRV record." alias:"SRV"`
}

// DNSAddACmd arguments.
type DNSAddACmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// Execute A record creation.
func (cmd *DNSAddACmd) Execute(args []string) error {
	err := client.AddARecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, false)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSAddAAAACmd arguments.
type DNSAddAAAACmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// Execute AAAA record creation.
func (cmd *DNSAddAAAACmd) Execute(args []string) error {
	err := client.AddARecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, true)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSAddCNAMECmd arguments.
type DNSAddCNAMECmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// Execute CNAME record creation.
func (cmd *DNSAddCNAMECmd) Execute(args []string) error {
	err := client.AddCNAME(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSAddMXCmd arguments.
type DNSAddMXCmd struct {
	Args DNSAddArgsPri `positional-args:"true"`
}

// Execute MX record creation.
func (cmd *DNSAddMXCmd) Execute(args []string) error {
	err := client.AddMX(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSAddNSCmd arguments.
type DNSAddNSCmd struct {
	Args DNSAddArgsPri `positional-args:"true"`
}

// Execute NS record creation.
func (cmd *DNSAddNSCmd) Execute(args []string) error {
	err := client.AddNS(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSAddTXTCmd arguments.
type DNSAddTXTCmd struct {
	Args DNSAddArgsPri `positional-args:"true"`
}

// Execute TXT record creation.
func (cmd *DNSAddTXTCmd) Execute(args []string) error {
	err := client.AddTXT(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}

// DNSAddSRVCmd arguments.
type DNSAddSRVCmd struct {
	Args struct {
		Domain   string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
		Value    string `required:"true" positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME."`
		Host     string `required:"true" positional-arg-name:"HOST" description:"A fully qualified service name in the form '_<service name>._<protocol>.<domain.tld>'."`
		TTL      int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
		Priority uint16 `positional-arg-name:"PRIORITY" description:"Priority of record. Default is 0 (most important)."`
		Port     uint16 `positional-arg-name:"PORT" description:"Port number of the service."`
		Weight   uint16 `positional-arg-name:"WEIGHT" description:"Relative weight for records with the same priority."`
	} `positional-args:"true"`
}

// Execute SRV record creation.
func (cmd *DNSAddSRVCmd) Execute(args []string) error {
	err := client.AddSRV(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority, cmd.Args.Port, cmd.Args.Weight)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}
