package main

import (
	"fmt"

	"gitlab.com/fillip/oki/providers/do"
)

func main() {
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
			fmt.Printf("Found droplet: %s", droplet.Name)
		}
	} else {
		fmt.Println("No droplets found")
	}
}
