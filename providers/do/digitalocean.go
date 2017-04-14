package providers

import (
	"context"

	"github.com/digitalocean/godo"
)

// DigitalOcean providers client API access for Digital Ocean services.
type DigitalOcean struct {
	client  *godo.Client
	context context.Context
}
