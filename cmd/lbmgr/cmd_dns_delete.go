package main

// DNSDeleteCmd arguments.
type DNSDeleteCmd struct {
	Args DNSArgs `positional-args:"true"`
}

// Execute the delete command.
func (cmd *DNSDeleteCmd) Execute(args []string) error {
	return nil
}
