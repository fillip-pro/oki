package providers

import (
	"log"
	"os"

	cf "github.com/cloudflare/cloudflare-go"
)

// CloudflareClient returns a new API client for CloudFlare.
func CloudflareClient() (*Cloudflare, error) {
	cloudflare := &Cloudflare{
		client: nil,
	}

	// Construct a new API object
	client, err := cf.New(os.Getenv("CF_KEY"), os.Getenv("CF_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	cloudflare.client = client

	return cloudflare, err
}
