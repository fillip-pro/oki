package composer_test

import (
	"log"
	"testing"

	"gitlab.com/fillip/oki/composer"
)

var cluster = &composer.ClusterComposer{}

func TestPrimaryCluster(t *testing.T) {
	t.Run("Primary Cluster", func(t *testing.T) {
		t.Run("Activate Primary Cluster", PrimaryClusterActivationTest)
	})

	if cluster != nil {
		clusterComposer, err := composer.NewClusterComposer()

		if err != nil {
			log.Fatal(err)
		}

		clusterComposer.DestroyPrimaryCluster()
	}
}

func PrimaryClusterActivationTest(t *testing.T) {
	clusterComposer, err := composer.NewClusterComposer()

	if err != nil {
		t.Fatal(err)
	}

	cluster, err := clusterComposer.CreatePrimaryCluster()

	if err != nil {
		t.Fatal(err)
	}

	if cluster != nil && cluster.ID != 0 {
		t.Logf("%d cluster created.", cluster.ID)
	} else {
		t.Errorf("Could not validate successful primary cluster creation.")
		t.FailNow()
	}
}
