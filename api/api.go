package api

import (
	"log"
	"time"

	"github.com/zerjioang/ddns-cloudflare/datatypes"
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

// Monitor start a foreground monitor that updates every
// 10 minutes the IP of the device
func Monitor() error {
	// make an initial update and then deploy the ticker
	log.Println("Trying to update device IP DNS Record")
	if err := Start(); err != nil {
		log.Println("failed to make initial IP update", err)
		return err
	}

	ticker := time.NewTicker(10 * time.Minute)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			case _, _ = <-ticker.C:
				log.Println("Trying to update device IP DNS Record")
				if err := Start(); err != nil {
					log.Println("failed to update device IP due to", err)
				}
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- struct{}{}
	log.Println("monitor stopped")
	return nil
}

func measureTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
