package main

import (
	"log"
	"os"

	"github.com/joostaarts/GolangHue/pkg/bridgediscovery"
)

func main() {
	bridgediscovery.StartDiscovery()
	log.Println("press a key")
	waitForKeyPress()
}

func waitForKeyPress() {
	b := make([]byte, 10)
	if _, err := os.Stdin.Read(b); err != nil {
		log.Fatal(err)
	}
}
