package api

import (
	"context"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

var (
	// Most API calls require a Context
	ctx = context.Background()
)

// triggerDDNSUpdate returns CLoudFLare zone identifier for given zone name
func triggerDDNSUpdate(token string, zone string, aName string, currIp string) error {
	// Construct a new API object
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		return err
	}
	// Fetch the zone ID
	// Assuming example.com exists in your Cloudflare account already
	zoneId, err := api.ZoneIDByName(zone)
	if err != nil {
		return err
	}
	log.Println("Cloudflare ZONE ID: ", zoneId)
	rr, err := api.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{
		Name: aName + "." + zone,
		Type: "A",
	})
	if err != nil {
		return err
	}
	log.Println("Cloudflare DNS RECORD found: ", len(rr), "records")
	if len(rr) == 0 {
		return errors.New("Could not get a valid DNS A RECORD for given configuration")
	}
	dnsa := rr[0]
	if currIp != dnsa.Content {
		log.Println("Last registered IP found with value : ", dnsa.Content)
		log.Println("New device IP value will be set to  : ", currIp)
		dnsa.Content = currIp
		updateErr := api.UpdateDNSRecord(ctx, zoneId, dnsa.ID, dnsa)
		if updateErr != nil {
			log.Println("failed to update DNS A record due to: ", err)
			return err
		}
		log.Println("Update successful")
		return nil
	}
	log.Println("No need to update DNS A Record")
	return nil
}
