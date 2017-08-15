package providers

import (
	"sync"

	cf "gitlab.com/fillip/oki/providers/cloudflare"
	do "gitlab.com/fillip/oki/providers/do"
)

var (
	providerClients *Clients
	once            sync.Once
)

// Clients provides client components for calling provider APIs.
type Clients struct {
	DigitalOcean *do.DigitalOcean
	Cloudflare   *cf.Cloudflare
}

// Providers initializes provider clients for service building.
func Providers() *Clients {
	once.Do(func() {
		providerClients = &Clients{}
		providerClients.DigitalOcean, _ = do.DigitalOceanClient()
		providerClients.Cloudflare, _ = cf.CloudflareClient()
	})

	return providerClients
}
