package main

import (
	"os"

	"github.com/Urethramancer/lbapi"
)

var client *lbapi.Client

func main() {
	if !loadConfig() {
		pr("Couldn't load configuration file.")
		os.Exit(2)
	}

	client = lbapi.NewClient(cfg.API, cfg.ID, cfg.Key)
	if client == nil {
		pr("Error creating client structure for '%s'.", cfg.API)
		os.Exit(2)
	}

	parseFlags()
}
