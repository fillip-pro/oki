package providers

import (
	cloudflare "github.com/cloudflare/cloudflare-go"
)

// Cloudflare providers client API access for Cloudflare services.
type Cloudflare struct {
	client *cloudflare.API
}
