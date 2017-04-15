package providers

import (
	"fmt"
	"log"

	"github.com/digitalocean/godo"
)

// Volume describes the state of a storage volume
// on Digital Ocean
type Volume struct {
	ID string
}

// ListVolumes lists volumes registed with Digital Ocean service.
func (digitalocean DigitalOcean) ListVolumes() {
	client, err := DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	volumes, _, err := client.client.Storage.ListVolumes(client.context, &godo.ListVolumeParams{})

	if err != nil {
		log.Fatal(err)
	}

	for _, volume := range volumes {
		fmt.Printf("Volume: %s\n", volume.ID)
	}
}

// GetVolumeByName retrieves a volume by a given name.
func (digitalocean DigitalOcean) GetVolumeByName(name string) (*godo.Volume, error) {
	doc, err := DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	volume, _, err := doc.client.Storage.GetVolume(doc.context, name)

	if err != nil {
		log.Fatal(err)
	}

	return volume, err
}

// CreateVolumeByName creates a new volume of block storage
// on digital ocean service.
func (digitalocean DigitalOcean) CreateVolumeByName(name string) (*godo.Volume, error) {
	doc, err := DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	createRequest := &godo.VolumeCreateRequest{
		Region:        "fra1",
		Name:          name,
		Description:   "Primary block storage for fillip.pro assets.",
		SizeGigaBytes: 10,
	}

	volume, _, err := doc.client.Storage.CreateVolume(doc.context, createRequest)

	if err != nil {
		log.Fatal(err)
	}

	return volume, err
}

// DeleteVolumeByID finds a named volume and deletes it from
// Digital Ocean.
func (digitalocean DigitalOcean) DeleteVolumeByID(id string) error {
	doc, err := DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	_, err = doc.client.Storage.DeleteVolume(doc.context, id)

	return err
}

// AttachVolumeToDroplet attaches a storage volume on Digital Ocean to a droplet.
func (digitalocean DigitalOcean) AttachVolumeToDroplet(volumeID string, dropletID int) error {
	doc, err := DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	_, _, err = doc.client.StorageActions.Attach(doc.context, volumeID, dropletID)

	return err
}

// DetachVolume detaches the volume from all clusters in Digital Ocean.
func (digitalocean DigitalOcean) DetachVolume(volumeID string) error {
	doc, err := DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	_, _, err = doc.client.StorageActions.Detach(doc.context, volumeID)

	return err
}
