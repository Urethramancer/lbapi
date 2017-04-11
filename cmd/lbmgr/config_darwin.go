package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func getConfigName() string {
	u, err := user.Current()
	if err != nil {
		pr("Fatal: %s", err.Error())
		os.Exit(2)
	}

	dir := filepath.Join(u.HomeDir, "Library", "Application Support", program)
	if !exists(dir) {
		createdir(dir)
	}
	return filepath.Join(dir, configname)
}
