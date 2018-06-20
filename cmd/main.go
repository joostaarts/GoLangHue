package main

import (
	"time"

	"github.com/joostaarts/GolangHue/pkg/bridgediscovery"
	"github.com/joostaarts/GolangHue/service/webserver"
)

func main() {
	go webserver.Startup()
	bridgediscovery.StartDiscovery()
	sleep()
}

func sleep() {
	for {
		time.Sleep(time.Second)
	}
}
