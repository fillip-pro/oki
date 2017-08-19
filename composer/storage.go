package composer

import (
	"gitlab.com/fillip/oki/log"
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

// Compose composes a new `Storage` configuration for use.
func (storageComposer StorageComposer) Compose() {
	_, err := storageComposer.CreatePrimaryStorage()

	if err != nil {
		log.Error(err)
	}

	/*
		detachStorage(storage)
			fmt.Printf("%s storage detached\n", storage.Name)

			storage, err = destroyPrimaryStorage(storage)
			fmt.Printf("%s storage deleted.\n", storage.Name)*/

}

// CreatePrimaryStorage creates primary storage volumes suitable for
// project.
func (storageComposer StorageComposer) CreatePrimaryStorage() (*Storage, error) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Error(err)
	}

	volume, err := doc.CreateVolumeByName("eu-volume-1-fillip-pro")

	if err != nil {
		log.Error(err)
	}

	log.Infof("Volume '%s' created!\n", volume.ID)

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
		log.Error(err)
	}

	err = doc.DeleteVolumeByID(storage.ID)

	return storage, err
}

func (storageComposer StorageComposer) detachStorage(storage *Storage) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Error(err)
	}

	err = doc.DetachVolume(storage.ID)
}
