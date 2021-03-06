package composer_test

import (
	"log"
	"testing"

	"gitlab.com/fillip/oki/composer"
)

var storage = &composer.StorageComposer{}

func TestStoragePrimaryStorage(t *testing.T) {
	t.Run("Primary Storage", func(t *testing.T) {
		t.Run("Activate Primary Storage", PrimaryStorageActivationTest)
	})

	if storage != nil {
		volume, err := storage.CreatePrimaryStorage()

		if err != nil {
			log.Fatal(err)
		}

		storage.DestroyPrimaryStorage(volume)
	}
}

func PrimaryStorageActivationTest(t *testing.T) {
	storage, err := composer.NewStorageComposer()

	if err != nil {
		t.Fatal(err)
	}

	volume, err := storage.CreatePrimaryStorage()

	if err != nil {
		t.Fatal(err)
	}

	if volume != nil && len(volume.ID) != 0 {
		t.Logf("%s storage volume created.", volume.ID)
	} else {
		t.Errorf("Could not validate successful primary storage volume creation.")
		t.FailNow()
	}
}
