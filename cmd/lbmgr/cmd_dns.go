package main

// DNSCmd holds the DNS sub-commands.
type DNSCmd struct {
	Status DNSStatusCmd `command:"status" description:"Shows DNS management status for a domain by order ID, and activates it if not yet enabled."`
	Get    DNSGetCmd    `command:"get" description:"Get a DNS record."`
}

type DNSArgs struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Type   string `required:"true" positional-arg-name:"TYPE" description:"Record type."`
}

type DNSStatusCmd struct {
	Verbose bool `short:"v" long:"verbose" description:"Show detailed information."`
	Args    struct {
		ID string `required:"true" positional-arg-name:"ID" description:"Domain order ID."`
	} `positional-args:"true"`
}

func (cmd *DNSStatusCmd) Execute(args []string) error {
	if client.DNSActive(cmd.Args.ID) {
		pr("DNS is active.")
	} else {
		pr("Non-existent or not yet activated.")
	}
	return nil
}
