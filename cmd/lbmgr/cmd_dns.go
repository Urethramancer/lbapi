package main

// DNSCmd arguments and sub-commands.
type DNSCmd struct {
	Status DNSStatusCmd `command:"status" description:"Shows DNS management status for a domain by order ID, and activates it if not yet enabled."`
	Get    DNSGetCmd    `command:"get" description:"Get a DNS record."`
	Add    DNSAddCmd    `command:"get" description:"Add a DNS record."`
	Delete DNSDeleteCmd `command:"delete" description:"Delete a DNS record." alias:"del" alias:"rm"`
	Change DNSChangeCmd `command:"change" description:"Edit a DNS record." alias:"ch" alias:"edit" alias:"modify" alias:"mod"`
}

// DNSArgs are the default arguments of most DNS sub-commands.
type DNSArgs struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Type   string `required:"true" positional-arg-name:"TYPE" description:"Record type."`
}

// DNSStatusCmd arguments.
type DNSStatusCmd struct {
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"Domain order ID."`
	} `positional-args:"true"`
}

// Execute the status command.
func (cmd *DNSStatusCmd) Execute(args []string) error {
	if client.DNSActive(cmd.Args.ID) {
		pr("DNS is active.")
	} else {
		pr("Non-existent or not yet activated.")
	}
	return nil
}

func getRecordType(in string) string {
	switch in {
	case "a", "A":
		return "A"
	case "mx", "MX":
		return "MX"
	case "cname", "CNAME":
		return "CNAME"
	case "txt", "TXT":
		return "TXT"
	case "ns", "NS":
		return "NS"
	case "srv", "SRV":
		return "SRV"
	case "aaaa", "AAAA":
		return "AAAA"
	default:
		return ""
	}
}
