package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opt struct {
	Version bool   `short:"V" long:"version" description:"Print version and exit."`
	Config  string `short:"C" long:"config" description:"Path to configuration file." value-name:"CONFIG"  default:"ream.json"`
}

func parseFlags() {
	_, err := flags.Parse(&opt)
	if err != nil {
		os.Exit(0)
	}

	if opt.Version {
		pr("%s %s", program, Version)
		os.Exit(0)
	}
}
