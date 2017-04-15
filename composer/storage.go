package composer

import (
	"fmt"
	"log"

	do "gitlab.com/fillip/oki/providers/do"
)

// StorageComposer provides details of current storage state for `Oki`.
type StorageComposer struct {
}

// Storage describes a generic storage volume.
type Storage struct {
	ID   string
	Name string
	Size int64
}

// NewStorageComposer creates a new storage object.
func NewStorageComposer() (*StorageComposer, error) {
	storageComposer := &StorageComposer{}

	return storageComposer, nil
}

// CreatePrimaryStorage creates primary storage volumes suitable for
// project.
func (storageComposer StorageComposer) CreatePrimaryStorage() (*Storage, error) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	volume, err := doc.CreateVolumeByName("eu-volume-1-fillip-pro")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Volume '%s' created!\n", volume.ID)

	response := &Storage{
		ID:   volume.ID,
		Name: volume.Name,
		Size: volume.SizeGigaBytes,
	}

	return response, nil
}

// DestroyPrimaryStorage deletes the primary registered storage
// volume.
func (storageComposer StorageComposer) DestroyPrimaryStorage(storage *Storage) (*Storage, error) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	err = doc.DeleteVolumeByID(storage.ID)

	return storage, err
}

func (storageComposer StorageComposer) DetachStorage(storage *Storage) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	err = doc.DetachVolume(storage.ID)
}
