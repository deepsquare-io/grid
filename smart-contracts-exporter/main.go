package main

import (
	"log"

	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln("couldn't execute command", err)
	}
}
