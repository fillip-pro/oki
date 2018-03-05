package main

import (
	"fmt"
	"sync"

	"gitlab.com/fillip/oki/composer"
	log "gitlab.com/fillip/oki/log"
	cloudflare "gitlab.com/fillip/oki/providers/cloudflare"
	do "gitlab.com/fillip/oki/providers/do"
)

func main() {
	log.Infoln("Welcome to Oki")

	//configuration.CloudConfig()
	generateCluster()
	providerStatus()
}

func providerStatus() {
	doStatus()
	cfStatus()
}

func doStatus() {
	do, _ := do.DigitalOceanClient()
	account, err := do.Account()

	if account != nil {
		log.Infof("Email: %s\n", account.Email)
		log.Infof("Droplet Limit: %d\n", account.DropletLimit)
		log.Infof("Floating IP Limit: %d\n", account.FloatingIPLimit)
	}

	if err != nil {
		log.Error(err)
	}

	keys, err := do.ListKeys()

	if keys != nil {
		for _, key := range keys {
			log.Infof("Found key: %s (%s)\n", key.Name, key.Fingerprint)
		}
	} else {
		log.Infoln("No keys found")
	}

	if err != nil {
		log.Error(err)
	}

	droplets, err := do.ListDropletsByTag("cluster")

	if droplets != nil && len(droplets) > 0 {
		for _, droplet := range droplets {
			log.Infof("Found droplet: %s\n", droplet.Name)
		}
	} else {
		log.Infoln("No droplets found")
	}
}

func cfStatus() {
	cloudflare, err := cloudflare.CloudflareClient()
	zones, err := cloudflare.ListZones()

	if zones != nil && len(zones) > 0 {

		for _, zone := range zones {
			log.Infof("Found zone: %s\n", zone.Name)
		}
	} else {
		log.Infoln("No zones found")
	}

	zone, err := cloudflare.GetZone("fillip.pro")

	if err != nil {
		log.Error(err)
	}

	dns, err := cloudflare.ListDNSRecords(zone)

	for _, r := range dns {
		log.Infof("%s: %s\n", r.Name, r.Content)
	}
}

func generateCluster() {
	messages := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		storageComposer, _ := composer.NewStorageComposer()
		storageComposer.Compose()
		messages <- 1
	}()

	go func() {
		defer wg.Done()
		clusterComposer, _ := composer.NewClusterComposer()
		clusterComposer.Compose()
		messages <- 2
	}()

	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
