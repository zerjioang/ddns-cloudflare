package datatypes

import "os"

// Payload wraps all required data to successfully execute a CF API DDNS A record update
type Payload struct {
	// Zone is Clouflare DNS zone
	// example: example.com
	Zone string
	// dnsrecord is the A record which will be updated
	// example: api.example.com
	DNSRecord string
	// cloudflare token string used for client authentication
	// note: this value ust be kept secret/safe/secure
	Token string
}

// NewPayload creates a new payload based on user requested data from ENV vars
func NewPayload() *Payload {
	return &Payload{
		Zone:      os.Getenv("X_CF_AGENT_ZONE"),
		DNSRecord: os.Getenv("X_CF_AGENT_DNS_A_RECORD"),
		Token:     os.Getenv("X_CF_AGENT_TOKEN"),
	}
}
