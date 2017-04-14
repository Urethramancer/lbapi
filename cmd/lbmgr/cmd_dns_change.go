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
}

// DNSChangeACmd arguments.
type DNSChangeACmd struct {
	Args DNSChangeArgs `positional-args:"true"`
}

// DNSChangeACmd modifies an A record.
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

// DNSChangeAAAACmd modifies an AAAA record.
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
	Args DNSAddArgs `positional-args:"true"`
}

// DNSChangeCNAMECmd modifies a CNAME record.
func (cmd *DNSChangeCNAMECmd) Execute(args []string) error {
	return nil
}

// DNSChangeMXCmd arguments.
type DNSChangeMXCmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// DNSChangeMXCmd modifies an MX record.
func (cmd *DNSChangeMXCmd) Execute(args []string) error {
	return nil
}

// DNSChangeNSCmd arguments.
type DNSChangeNSCmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// DNSChangeNSCmd modifies an NS record.
func (cmd *DNSChangeNSCmd) Execute(args []string) error {
	return nil
}

// DNSChangeTXTCmd arguments.
type DNSChangeTXTCmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// DNSChangeTXTCmd modifies a TXT record.
func (cmd *DNSChangeTXTCmd) Execute(args []string) error {
	return nil
}

// DNSChangeSRVCmd arguments.
type DNSChangeSRVCmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// DNSChangeSRVCmd modifies a SRV record.
func (cmd *DNSChangeSRVCmd) Execute(args []string) error {
	return nil
}
