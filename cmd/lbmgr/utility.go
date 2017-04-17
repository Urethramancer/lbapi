package main

import "fmt"

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func prn(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}

func okColour(b bool) string {
	if b {
		return ANSI_GREEN
	}

	return ANSI_RED
}
