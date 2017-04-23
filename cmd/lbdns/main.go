//go:generate ../build
package main

import (
	"os"

	"github.com/Urethramancer/lbapi"
	"github.com/Urethramancer/lbapi/common"
)

var client *lbapi.ProxyClient

func main() {
	if !loadConfig() {
		pr("Couldn't load configuration file.")
		os.Exit(2)
	}

	client = lbapi.NewProxyClient(cfg.API, cfg.Username, cfg.Password)
	if client == nil {
		pr("Error creating client structure for '%s'.", cfg.API)
		os.Exit(2)
	}

	common.SetClient(client)
	parseFlags()
}
