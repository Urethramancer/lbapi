// +build linux dragonfly freebsd netbsd openbsd solaris

package common

import (
	"os"
	"os/user"
	"path/filepath"
)

// GetConfigName gets the correct full path of the configuration file.
func GetConfigName(program, filename string) (string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(u.HomeDir, "."+program)
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(dir, filename)
}

// GetServerConfigName gets the correct full path of the configuration file for servers.
func GetServerConfigName(program, filename string) (string, error) {
	dir := filepath.Join("/etc", program)
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(dir, filename)
}
