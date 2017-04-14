package providers

import (
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// ListZones lists the zones registered with CloudFlare.
func (cloudflare *Cloudflare) ListZones() ([]cloudflare.Zone, error) {
	client, err := CloudflareClient()

	api := client.client

	zones, err := api.ListZones()

	if err != nil {
		log.Fatal(err)
	}

	return zones, err

	/*// Fetch the zone ID
	id, err := api.ZoneIDByName("fillip.pro") // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zone)

	// Fetch all records for a zone
	recs, err := api.DNSRecords(id, cloudflare.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range recs {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}*/
}

// GetZone returns a specified zone.
func (cloudflare *Cloudflare) GetZone(zoneName string) (cloudflare.Zone, error) {
	client, err := CloudflareClient()
	if err != nil {
		log.Fatal(err)
	}

	id, err := client.client.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatal(err)
	}

	zone, err := client.client.ZoneDetails(id)
	if err != nil {
		log.Fatal(err)
	}

	return zone, err
}
