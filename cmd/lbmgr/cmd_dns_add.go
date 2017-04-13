package main

import "errors"

// DNSAddCmd arguments.
type DNSAddCmd struct {
	Args struct {
		Domain  string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
		Type    string `required:"true" positional-arg-name:"TYPE" description:"Record type."`
		Address string `required:"true" positional-arg-name:"IP" description:"IP address."`
		Host    string `positional-arg-name:"HOST" description:"Host name."`
		TTL     int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
	} `positional-args:"true"`
}

// Execute the add command.
func (cmd *DNSAddCmd) Execute(args []string) error {
	var err error
	switch cmd.Args.Type {
	case "a", "A":
		err = client.AddDNSA(cmd.Args.Domain, cmd.Args.Host, cmd.Args.Address, cmd.Args.TTL, false)
	case "aaaa", "AAAA":
		err = client.AddDNSA(cmd.Args.Domain, cmd.Args.Host, cmd.Args.Address, cmd.Args.TTL, true)
	default:
		return errors.New("Unknown record type '" + cmd.Args.Type + "'")
	}

	if err != nil {
		return err
	}

	pr("Record added.")
	return nil
}
