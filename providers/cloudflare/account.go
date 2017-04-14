package providers

import (
	"fmt"
	"log"
)

// Account lists the account details with CloudFlare.
func (cloudflare *Cloudflare) Account() {
	client, err := CloudflareClient()

	// Fetch user details on the account
	u, err := client.client.UserDetails()

	if err != nil {
		log.Fatal(err)
	}

	// Print user details
	fmt.Println(u)
}
