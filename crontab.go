package goutils

import "time"

//Crontab run function f on every period seconds
func Crontab(period time.Duration, f func()) {
	ticker := time.NewTicker(period * time.Second)
	go func() {
		for range ticker.C {
			f()
		}
	}()
}
