package main

import "errors"

// DNSDeleteCmd arguments.
type DNSDeleteCmd struct {
	Args struct {
		Domain  string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
		Type    string `required:"true" positional-arg-name:"TYPE" description:"Record type."`
		Address string `required:"true" positional-arg-name:"IP" description:"IP address."`
		Host    string `positional-arg-name:"HOST" description:"Host name."`
	} `positional-args:"true"`
}

// Execute the delete command.
func (cmd *DNSDeleteCmd) Execute(args []string) error {
	var err error
	switch cmd.Args.Type {
	case "a", "A":
		err = client.DeleteDNSRecord(cmd.Args.Domain, cmd.Args.Host, cmd.Args.Address, false)
	case "aaaa", "AAAA":
		err = client.DeleteDNSRecord(cmd.Args.Domain, cmd.Args.Host, cmd.Args.Address, true)
	case "cname", "CNAME":
	case "mx", "MX":
	case "ns", "NS":
	case "txt", "TXT":
	case "srv", "SRV":
	default:
		return errors.New("Unknown record type '" + cmd.Args.Type + "'")
	}

	if err != nil {
		return err
	}

	pr("Record deleted.")

	return nil
}
