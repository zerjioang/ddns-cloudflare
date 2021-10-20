package api

import (
	"log"
	"time"

	"github.com/zerjioang/cf-agent/datatypes"
)

// Start executes the IP update method
func Start() error {
	defer measureTime(time.Now(), "cloudflare ddns updater")
	log.Println("Updating device IP. Please wait...")
	p := datatypes.NewPayload()
	log.Println("Requesting IP check for: ", p.Zone+"."+p.DNSRecord)
	log.Println("Reading current device IP. Please wait...")
	currIp := externalIP()
	log.Println("Readed IP: ", currIp)
	log.Println("Connecting with Cloudflare services...")
	return triggerDDNSUpdate(p.Token, p.Zone, "console", currIp)
}

func measureTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
