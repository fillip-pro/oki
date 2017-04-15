package composer

import (
	"fmt"
	"log"

	do "gitlab.com/fillip/oki/providers/do"
)

// ClusterComposer manages the composition of clusters.
type ClusterComposer struct {
}

// Cluster contains all of the data
// related to a given cluster state.
type Cluster struct {
	ID   int
	Name string
	Tags []string
}

// NewClusterComposer creates a new cluster composer.
func NewClusterComposer() (*ClusterComposer, error) {
	clusterComposer := &ClusterComposer{}

	return clusterComposer, nil
}

// CreatePrimaryCluster creates a primary cluster for the service.
func (clusterComposer ClusterComposer) CreatePrimaryCluster() (*Cluster, error) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	droplet := &do.Droplet{
		Name: "eu-cluster-1.fillip.pro",
		Tags: []string{"cluster"},
	}

	droplet, err = doc.CreateDroplet(droplet)

	fmt.Printf("Cluster '%d' created!\n", droplet.ID)

	if err != nil {
		log.Fatal(err)
	}

	cluster := &Cluster{
		ID:   droplet.ID,
		Name: droplet.Name,
		Tags: droplet.Tags,
	}

	return cluster, err
}

// DestroyPrimaryCluster destroys the primary cluster.
func (clusterComposer ClusterComposer) DestroyPrimaryCluster() error {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	err = doc.DeleteDropletsByTag("cluster")

	return err
}

// AttachStorageToCluster attaches volume storage to a cluster.
func (clusterComposer ClusterComposer) AttachStorageToCluster(storage *Storage, cluster *Cluster) {
	doc, err := do.DigitalOceanClient()

	if err != nil {
		log.Fatal(err)
	}

	doc.AttachVolumeToDroplet(storage.ID, cluster.ID)
}
