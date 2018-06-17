package main

import (
	"time"

	"github.com/joostaarts/GolangHue/pkg/bridgediscovery"
)

func main() {
	bridgediscovery.StartDiscovery()
	sleep()
}

func sleep() {
	for {
		time.Sleep(time.Second)
	}
}
