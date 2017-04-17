package common

import "os"

// Exists returns true if a file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
