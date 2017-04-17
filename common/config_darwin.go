package common

import (
	"os"
	"os/user"
	"path/filepath"
)

// GetConfigName gets the correct full path of the configuration file for command line tools.
func GetConfigName(program, filename string) (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(u.HomeDir, "Library", "Application Support", program)
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(dir, filename), nil
}

// GetServerConfigName gets the correct full path of the configuration file for servers.
func GetServerConfigName(program, filename string) (string, error) {
	return GetConfigName(program, filename)
}
