package main

import (
	"log"

	"github.com/zerjioang/ddns-cloudflare/cmd"
)

// main is the main entry point of the service
func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
