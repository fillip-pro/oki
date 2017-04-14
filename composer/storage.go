package composer

import (
	"fmt"
	"log"

	do "gitlab.com/fillip/oki/providers/do"
)

// Storage provides details of current storage state for `Oki`.
type Storage struct {
}

// Volume describes a generic storage volume.
type Volume struct {
	ID   string
	Name string
	Size int64
}

// NewStorage creates a new storage object.
func NewStorage() (*Storage, error) {
	storage := &Storage{}

	return storage, nil
}

// CreatePrimaryStorage creates primary storage volumes suitable for
// project.
func (storage Storage) CreatePrimaryStorage() (*Volume, error) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	volume, err := doc.CreateVolumeByName("eu-volume-01-fillip-pro")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Volume '%s' created!", volume.ID)

	response := &Volume{
		ID:   volume.ID,
		Name: volume.Name,
		Size: volume.SizeGigaBytes,
	}

	return response, nil
}

// DestroyPrimaryStorage deletes the primary registered storage
// volume.
func (storage Storage) DestroyPrimaryStorage(volume *Volume) (*Volume, error) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	err = doc.DeleteVolumeByID(volume.ID)

	return volume, err
}
