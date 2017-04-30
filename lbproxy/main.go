//go:generate ../../cmd/build
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Urethramancer/lbapi"
)

var client *lbapi.Client

func main() {
	client = lbapi.NewClient(cfg.LogicBoxes.API, cfg.LogicBoxes.ID, cfg.LogicBoxes.Key)
	if client == nil {
		crit("Error setting up access for '%s'", cfg.LogicBoxes.API)
	}

	openLogs()
	defer closeLogs()
	startJanitor()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Channels.janitorquit <- true
		Channels.mainquit <- true
	}()

	initWeb()
	<-Channels.mainquit
	syslog("Quit signal received. Shutting down.")
	go stopServers()
	time.Sleep(time.Millisecond * 500)
	return
}
