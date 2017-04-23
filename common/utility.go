package common

import (
	"fmt"
	"os"
)

// Exists returns true if a file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func prn(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}
