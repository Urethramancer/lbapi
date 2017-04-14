package main

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

// DNSAddACmd adds an A record.
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

// DNSAddAAAACmd adds an AAAA record.
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

// DNSAddCNAMECmd adds a CNAME record.
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

// DNSAddMXCmd adds an MX record.
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

// DNSAddNSCmd adds an NS record.
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

// DNSAddTXTCmd adds a TXT record.
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
		Host     string `positional-arg-name:"HOST" description:"Host name."`
		TTL      int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
		Priority uint   `positional-arg-name:"PRIORITY" description:"Priority of record. Default is 0 (most important)."`
		Port     uint   `positional-arg-name:"PORT" description:"Port number of the service."`
		Weight   uint   `positional-arg-name:"WEIGHT" description:"Relative weight for records with the same priority."`
	} `positional-args:"true"`
}

// DNSAddSRVCmd adds a SRV record.
func (cmd *DNSAddSRVCmd) Execute(args []string) error {
	err := client.AddSRV(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, cmd.Args.TTL, cmd.Args.Priority, cmd.Args.Port, cmd.Args.Weight)
	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}
