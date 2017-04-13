package main

import (
	"encoding/json"
	"io/ioutil"

	"os"

	"github.com/Urethramancer/lbapi"
)

// Config holds the reseller settings, and optionally, the
// URL of a different API to be used.
type Config struct {
	API string `json:"api,omitempty"`
	ID  string `json:"resellerid"`
	Key string `json:"apikey"`
}

var cfg Config
var cfgpath string

const (
	configname = "config.json"
)

func init() {
	cfgpath = getConfigName()
}

func loadConfig() bool {
	if !exists(cfgpath) {
		pr("%s does not exist, creating.", cfgpath)
		cfg.API = lbapi.APIURL
		cfg.ID = "your reseller ID"
		cfg.Key = "your reseller API key"
		res, err := json.MarshalIndent(cfg, "", "\t")
		if err != nil {
			pr("Couldn't save default configuration: %s", err.Error())
			os.Exit(2)
		}

		f, err := os.OpenFile(cfgpath, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			os.Exit(2)
		}
		defer f.Close()

		_, err = f.Write(res)
		if err != nil {
			os.Exit(2)
		}

		_, err = f.WriteString("\n")
		if err != nil {
			os.Exit(2)
		}
	}

	data, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		pr("Error loading %s: %s", cfgpath, err)
		return false
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		pr("Error decoding %s: %s", cfgpath, err)
		return false
	}

	return true
}
