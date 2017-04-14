package providers

import (
	"log"

	cf "github.com/cloudflare/cloudflare-go"
)

// ListDNSRecords returns a list of DNS records for a given zone.
func (cloudflare *Cloudflare) ListDNSRecords(zone cf.Zone) ([]cf.DNSRecord, error) {
	client, err := CloudflareClient()
	// Fetch all records for a zone
	recs, err := client.client.DNSRecords(zone.ID, cf.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}

	return recs, err
}
