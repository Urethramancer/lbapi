package main

// DNSChangeCmd arguments.
type DNSChangeCmd struct {
	Args DNSAddArgs `positional-args:"true"`
}

// Execute the change command.
func (cmd *DNSChangeCmd) Execute(args []string) error {
	return nil
}
