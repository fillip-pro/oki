package composer_test

import (
	"log"
	"testing"

	"gitlab.com/fillip/oki/composer"
)

var volume = &composer.Volume{}

func TestStoragePrimaryStorage(t *testing.T) {
	t.Run("Primary Storage", func(t *testing.T) {
		t.Run("Activate Primary Storage", isPrimaryStorageActivatingTest)
	})

	if volume != nil {
		storage, err := composer.NewStorage()

		if err != nil {
			log.Fatal(err)
		}

		storage.DestroyPrimaryStorage(volume)
	}
}

func isPrimaryStorageActivatingTest(t *testing.T) {
	storage, err := composer.NewStorage()

	if err != nil {
		t.Fatal(err)
	}

	volume, err := storage.CreatePrimaryStorage()

	if err != nil {
		t.Fatal(err)
	}

	if volume != nil {
		t.Logf("%s storage volume created.", volume.ID)
	}
}
