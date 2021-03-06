package providers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/digitalocean/godo"
)

// Droplet contains the structure of a Digital Ocean instance.
type Droplet struct {
	ID   int
	Name string
	Tags []string
}

// CreateDroplet creates a cluster droplet.
func (digitalocean *DigitalOcean) CreateDroplet(droplet *Droplet) (*Droplet, error) {
	ignitionConfig, err := ioutil.ReadFile("/Users/pjlaszkowicz/Workspaces/Fillip/oki/src/gitlab.com/fillip/oki/configuration/do-cloud-config.ign")

	createRequest := &godo.DropletCreateRequest{
		Name:              droplet.Name,
		Region:            "fra1",
		Size:              "512mb",
		PrivateNetworking: true,
		IPv6:              true,
		Tags:              droplet.Tags,
		UserData:          string(ignitionConfig),
		Image: godo.DropletCreateImage{
			Slug: "coreos-stable",
		},
	}

	client, err := DigitalOceanClient()

	if err != nil {

	}

	dropletResponse, _, err := client.client.Droplets.Create(client.context, createRequest)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Droplet status is %s\n", dropletResponse.Status)

	for dropletResponse.Status != "active" {
		time.Sleep(3000 * time.Millisecond)

		dropletResponse, _, err = client.client.Droplets.Get(client.context, dropletResponse.ID)

		fmt.Printf("Waiting for %d to come online...\n", dropletResponse.ID)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Droplet %d is online!\n", dropletResponse.ID)

	droplet.ID = dropletResponse.ID

	return droplet, err
}

// DeleteDroplet deletes a droplet from Digital Ocean based on
// the droplet's ID.
func (digitalocean *DigitalOcean) DeleteDroplet(droplet Droplet) (Droplet, error) {
	if droplet.ID == 0 {
		return droplet, errors.New("No ID provided for deletion")
	}

	client, err := DigitalOceanClient()

	if err != nil {

	}

	_, err = client.client.Droplets.Delete(client.context, droplet.ID)

	return droplet, err
}

// DeleteDropletsByTag deletes all droplets with a matching tag.
func (digitalocean *DigitalOcean) DeleteDropletsByTag(tag string) error {
	client, err := DigitalOceanClient()

	if err != nil {

	}

	_, err = client.client.Droplets.DeleteByTag(client.context, tag)

	for {
		time.Sleep(3000 * time.Millisecond)
		droplets, err := digitalocean.ListDropletsByTag(tag)

		if err != nil {

		}

		if len(droplets) > 0 {
			fmt.Printf("%d %s droplets found. Deleting...\n", len(droplets), tag)
			_, err = client.client.Droplets.DeleteByTag(client.context, tag)
		} else {
			break
		}
	}

	return err
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
