package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"os"

	"github.com/Urethramancer/cross"
)

const (
	program = "lbproxy"
)

// Version is injected via build flags from git tags.
var Version = "undefined"

// Config is used to parse the main configuration file.
type Config struct {
	// APIPath is the path before each endpoint.
	APIPath string `json:"apipath"`
	// CleanupInterval is how often to clear out stale authentication data.
	CleanupInterval string `json:"cleanupinterval"`
	// Logs settings.
	Logs `json:"logs"`
	// Security settings.
	Security `json:"security"`
	// Web settings.
	Web `json:"web"`
	// LogicBoxes credentials.
	LogicBoxes `json:"logicboxes"`
	// InfoDump for site information.
	InfoDump `json:"infodump"`
}

// LogicBoxes API configuration.
type LogicBoxes struct {
	API string `json:"api,omitempty"`
	ID  int64  `json:"resellerid"`
	Key string `json:"key"`
}

// Logs hold the names of log files to output to.
type Logs struct {
	// Info is for general messages.
	Info string `json:"info,omitempty"`
	// Error is for serious warnings.
	Error string `json:"error,omitempty"`
}

// Security defines password security level and many SSL settings.
type Security struct {
	Certificate string `json:"certificate,omitempty"`
	Key         string `json:"key,omitempty"`
	// SessionLength is the time in the form of number+unit before a session is deleted.
	// Example: 60m, 3600s, 1h are all equal.
	//
	// This is refreshed when used.
	SessionLength string `json:"sessionlength"`
	// sessionLength is the converted SessionLength.
	sessionlength time.Duration
	// SSL enables the secure web server instead. A certificate pair is required for it to work.
	SSL bool `json:"ssl,omitempty"`
}

// Web defines addresses, ports and domains to route.
type Web struct {
	// Address is the IP address to bind to. Valid entries are any reachable address,
	// or 0.0.0.0 to bind to all public addresses, or even 127.0.0.1 if you rely on
	// a proxy server to expose it to the world.
	Address string `json:"address"`
	// Port to bind HTTP to.
	Port string `json:"port,omitempty"`
	// Domain is the FQDN for the server, and is mostly used for display purposes.
	Domain string `json:"domain"`
	url    string
}

// InfoDump is just that - a collection of information relating to the service.
// The /info command returns this to interested clients.
// Note: May be entirely empty if no data is defined.
type InfoDump struct {
	// Website for more information about the service.
	Website string `json:"website,omitempty"`
	// Other information of interest. Or just a big, fat ad.
	Other []string `json:"other,omitempty"`
	// Defaults for DNS records, if any. The counterpart for lbdns may be more useful.
	Defaults DomainDefaults `json:"defaults,omitempty"`
}

// DomainDefaults is a collection of DNS records to create for a domain if the
// client asks for a setup from scratch.
type DomainDefaults struct {
	TTL int64
	// MS defaults are handy if all domains share a mail host.
	// TODO: Figure out something for SPF records.
	//MX ??? `json:"mx"`
	// NS records are set up by default initially, but disappear after the user nukes them.
	// NS ??? `json:"ns"`
	// TXT records may be useful for something.
	// TXT ??? `json:"txt"`
	// SRV records can be handy when setting up a cluster of chat services
	// announcing their capabilities in SRV records (like some Microsoft services do).
	// SRV ??? `json:"srv"`
}

var cfg Config

func init() {
	parseFlags()
	loadConfig()
	initChannels()
}

func defaultConfig() Config {
	return Config{
		CleanupInterval: "10m",
		LogicBoxes: LogicBoxes{
			API: "",
			ID:  0,
			Key: "your key here",
		},
		Logs: Logs{
			Info:  "info.log",
			Error: "error.log",
		},
		Security: Security{
			Certificate:   "./cert.pem",
			Key:           "./key.pem",
			SessionLength: "10m",
		},
		Web: Web{
			Address: "127.0.0.1",
			Port:    "11000",
			Domain:  "localhost",
		},
	}
}

func loadConfig() {
	var err error

	name := opt.Config
	if name == "" {
		name, err = cross.GetConfigName(program, "lbproxy.json")
		if err != nil {
			crit("Error getting configuration: %s", err.Error())
			os.Exit(2)
		}
	}

	if !cross.Exists(name) {
		cfg = defaultConfig()
		var data []byte
		data, err = json.Marshal(cfg)
		check(err)
		err = ioutil.WriteFile(name, data, 0600)
		if err != nil {
			crit("Error saving default configuration %s: %s", name, err.Error())
		}
		syslog("Created default configuration.")
		return
	}

	data, err := ioutil.ReadFile(name)
	check(err)

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		crit("Error loading %s: %s", name, err.Error())
	}

	syslog("Loaded config from %s", name)
	cfg.Security.sessionlength = getTime(cfg.Security.SessionLength)
}
