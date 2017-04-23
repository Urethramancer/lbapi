package main

import (
	"fmt"

	"github.com/Urethramancer/lbapi/common"
)

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func prn(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}

func okColour(b bool) string {
	if b {
		return common.ANSI_GREEN
	}

	return common.ANSI_RED
}
