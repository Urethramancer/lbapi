package main

import (
	"os"

	"github.com/Urethramancer/lbapi/common"
	flags "github.com/jessevdk/go-flags"
)

const (
	program = "lbdns"
)

// Version is filled in from git tags if using the supplied build script.
var Version = "undefined"

var opt struct {
	Ver VersionCmd    `command:"version" description:"Print program version and exit." alias:"ver"`
	DNS common.DNSCmd `command:"dns" description:"Manage DNS for domains."`
}

func parseFlags() {
	_, err := flags.Parse(&opt)
	if err != nil {
		os.Exit(0)
	}
}

// VersionCmd is a place to hang the Execute command.
type VersionCmd struct {
}

// Execute prints the version to the console and exits.
func (cmd *VersionCmd) Execute(args []string) error {
	pr("%s %s", program, Version)
	return nil
}
