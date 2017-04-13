package main

// DNSAddCmd arguments.
type DNSAddCmd struct {
	Args DNSArgs `positional-args:"true"`
}

// Execute the add command.
func (cmd *DNSAddCmd) Execute(args []string) error {
	return nil
}
