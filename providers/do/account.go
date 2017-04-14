package providers

import (
	"github.com/digitalocean/godo"
)

// Account provides the `godo.Account` information from Digital Ocean.
func (digitalocean *DigitalOcean) Account() (*godo.Account, error) {
	client, err := DigitalOceanClient()

	account, _, err := client.client.Account.Get(client.context)

	return account, err
}

// ListKeys lists the keys registered with Digital Ocean
func (digitalocean *DigitalOcean) ListKeys() ([]godo.Key, error) {
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	client, err := DigitalOceanClient()

	keys, _, err := client.client.Keys.List(client.context, opt)

	return keys, err
}
