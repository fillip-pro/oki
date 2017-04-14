package do

import (
	"github.com/digitalocean/godo"
)

// CreateDroplet creates a cluster droplet.
func CreateDroplet() {
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

	ctx, client, err := NewContext()

	if err != nil {

	}

	_, _, err = client.Droplets.Create(ctx, createRequest)

	if err != nil {

	}
}

// ListDropletsByTag lists droplets by a given tag.
func ListDropletsByTag(tag string) ([]godo.Droplet, error) {
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	ctx, client, err := NewContext()

	if err != nil {

	}

	droplets, _, err := client.Droplets.ListByTag(ctx, tag, opt)

	return droplets, err
}
