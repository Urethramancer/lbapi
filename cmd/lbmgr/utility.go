package main

import (
	"fmt"
	"os"
)

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func qpr(f string, v ...interface{}) {
	// if opts.Quiet {
	// return
	// }
	fmt.Printf(f+"\n", v...)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// createdir is used when the existence of a directory is absolutely required.
func createdir(dir string) string {
	if !exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			pr("Fatal: Couldn't create directory %s", dir)
			os.Exit(2)
		}
	}
	return dir
}
