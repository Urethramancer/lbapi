// The janitor cleans up expired tokens.
package main

import (
	"time"
)

func startJanitor() {
	syslog("Starting janitor with interval %s.", cfg.CleanupInterval)
	ticker := time.NewTicker(getTime(cfg.CleanupInterval))
	go func() {
		for {
			select {
			case <-ticker.C:
				clearTokens()
			case <-Channels.janitorquit:
				syslog("Shutting down janitor.")
				ticker.Stop()
				return
			}
		}
	}()
}
