package main

import (
	"fmt"
	"log"

	do "gitlab.com/fillip/oki/providers/do"
)

// Storage provides details of current storage state for `Oki`.
type Storage struct {
}

// NewStorage creates a new storage object.
func NewStorage() (*Storage, error) {
	storage := &Storage{}

	return storage, nil
}

// CreatePrimaryStorage creates primary storage volumes suitable for
// project.
func (storage Storage) CreatePrimaryStorage() {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	volume, err := doc.CreateVolumeByName("eu-volume-01-fillip-pro")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Volume '%s' created!", volume.ID)
}
