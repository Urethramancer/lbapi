package main

// DNSDeleteCmd arguments.
type DNSDeleteCmd struct {
	A     DNSDeleteACmd     `command:"a" description:"Delete an A record." alias:"A"`
	AAAA  DNSDeleteAAAACmd  `command:"aaaa" description:"Delete an AAAA record." alias:"AAAA"`
	CNAME DNSDeleteCNAMECmd `command:"cname" description:"Delete a CNAME record." alias:"CNAME"`
	MX    DNSDeleteMXCmd    `command:"mx" description:"Delete an MX record." alias:"MX"`
	NS    DNSDeleteNSCmd    `command:"ns" description:"Delete an NS record." alias:"NS"`
	TXT   DNSDeleteTXTCmd   `command:"txt" description:"Delete a TXT record." alias:"TXT"`
	SRV   DNSDeleteSRVCmd   `command:"srv" description:"Delete a SRV record." alias:"SRV"`
}

// DNSDeleteACmd arguments.
type DNSDeleteACmd struct {
	Args DNSDeleteArgs `positional-args:"true"`
}

// DNSDeleteACmd deletes an A record.
func (cmd *DNSDeleteACmd) Execute(args []string) error {
	err := client.DeleteARecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, false)
	if err != nil {
		return err
	}

	pr("Record deleted.")
	return nil
}

// DNSDeleteAAAACmd arguments.
type DNSDeleteAAAACmd struct {
	Args DNSDeleteArgs `positional-args:"true"`
}

// DNSDeleteAAAACmd deletes an AAAA record.
func (cmd *DNSDeleteAAAACmd) Execute(args []string) error {
	err := client.DeleteARecord(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host, true)
	if err != nil {
		return err
	}

	pr("Record deleted.")
	return nil
}

// DNSDeleteCNAMECmd arguments.
type DNSDeleteCNAMECmd struct {
	Args DNSDeleteArgsAll `positional-args:"true"`
}

// DNSDeleteCNAMECmd deletes a CNAME record.
func (cmd *DNSDeleteCNAMECmd) Execute(args []string) error {
	err := client.DeleteCNAME(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host)
	if err != nil {
		return err
	}

	pr("Record deleted.")
	return nil
}

// DNSDeleteMXCmd arguments.
type DNSDeleteMXCmd struct {
	Args DNSDeleteArgsAll `positional-args:"true"`
}

// DNSDeleteMXCmd deletes a MX record.
func (cmd *DNSDeleteMXCmd) Execute(args []string) error {
	err := client.DeleteMX(cmd.Args.Domain, cmd.Args.Value, cmd.Args.Host)
	if err != nil {
		return err
	}

	pr("Record deleted.")
	return nil
}

// DNSDeleteNSCmd arguments.
type DNSDeleteNSCmd struct {
	Args DNSDeleteArgs `positional-args:"true"`
}

// DNSDeleteTXTCmd arguments.
type DNSDeleteTXTCmd struct {
	Args DNSDeleteArgs `positional-args:"true"`
}

// DNSDeleteSRVCmd arguments.
type DNSDeleteSRVCmd struct {
	Args struct {
		Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
		Value  string `required:"true" positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME/MX etc."`
		Host   string `positional-arg-name:"HOST" description:"Host name."`
		Port   int    `positional-arg-name:"PORT" description:"Port number of the service."`
		Weight int    `positional-arg-name:"WEIGHT" description:"Weight of the service."`
	} `positional-args:"true"`
}
