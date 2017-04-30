package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Urethramancer/lbapi/common"
)

// Config holds the login credentials and access URL.
type Config struct {
	API      string `json:"api"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var cfg Config
var cfgpath string

const (
	configname = "config.json"
)

func init() {
	var err error
	cfgpath, err = common.GetConfigName(program, configname)
	if err != nil {
		pr("Error getting configuration path: %s", err.Error())
	}
}

func GetCredentials(uq, pq string) (string, string) {
	prn("%s: ", uq)
	in := bufio.NewReader(os.Stdin)
	username, _ := in.ReadString('\n')
	username = strings.TrimSpace(username)
	prn("%s: ", pq)
	password, _ := in.ReadString('\n')
	password = strings.TrimSpace(password)
	return username, password
}

func loadConfig() bool {
	if !common.Exists(cfgpath) {
		pr("%s does not exist, creating.", cfgpath)
		cfg.API = "http://localhost:11000"
		// cfg.Username = "user@example.com"
		// cfg.Password = "secret"
		cfg.Username, cfg.Password = GetCredentials("Username", "Password")

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
