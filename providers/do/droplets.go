package providers

import (
	"github.com/digitalocean/godo"
)

// CreateDroplet creates a cluster droplet.
func (digitalocean *DigitalOcean) CreateDroplet() {
	createRequest := &godo.DropletCreateRequest{
		Name:              "eu-cluster-1.fillip.pro",
		Region:            "ams3",
		Size:              "512mb",
		PrivateNetworking: true,
		IPv6:              true,
		Tags:              []string{"cluster"},
		Image: godo.DropletCreateImage{
			Slug: "coreos-stable",
		},
	}

	client, err := DigitalOceanClient()

	if err != nil {

	}

	_, _, err = client.client.Droplets.Create(client.context, createRequest)

	if err != nil {

	}
}

// ListDropletsByTag lists droplets by a given tag.
func (digitalocean *DigitalOcean) ListDropletsByTag(tag string) ([]godo.Droplet, error) {
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	client, err := DigitalOceanClient()

	if err != nil {

	}

	droplets, _, err := client.client.Droplets.ListByTag(client.context, tag, opt)

	return droplets, err
}
