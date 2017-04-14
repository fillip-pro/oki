package main

import (
	"fmt"

	cloudflare "gitlab.com/fillip/oki/providers/cloudflare"
	do "gitlab.com/fillip/oki/providers/do"
)

func main() {
	do, _ := do.DigitalOceanClient()
	account, err := do.Account()

	if account != nil {
		fmt.Printf("Email: %s\n", account.Email)
		fmt.Printf("Droplet Limit: %d\n", account.DropletLimit)
		fmt.Printf("Floating IP Limit: %d\n", account.FloatingIPLimit)
	}

	if err != nil {
		fmt.Println(err)
	}

	keys, err := do.ListKeys()

	if keys != nil {
		for _, key := range keys {
			fmt.Printf("Found key: %s (%s)\n", key.Name, key.Fingerprint)
		}
	} else {
		fmt.Println("No keys found")
	}

	if err != nil {
		fmt.Println(err)
	}

	droplets, err := do.ListDropletsByTag("cluster")

	if droplets != nil && len(droplets) > 0 {
		for _, droplet := range droplets {
			fmt.Printf("Found droplet: %s\n", droplet.Name)
		}
	} else {
		fmt.Println("No droplets found")
	}

	cloudflare, err := cloudflare.CloudflareClient()
	zones, err := cloudflare.ListZones()

	if zones != nil && len(zones) > 0 {

		for _, zone := range zones {
			fmt.Printf("Found zone: %s\n", zone.Name)
		}
	} else {
		fmt.Println("No zones found")
	}

	zone, err := cloudflare.GetZone("fillip.pro")

	if err != nil {
		fmt.Println(err)
	}

	dns, err := cloudflare.ListDNSRecords(zone)

	for _, r := range dns {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}

	storage, _ := NewStorage()
	storage.CreatePrimaryStorage()

	do.ListVolumes()
}
