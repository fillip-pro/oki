package composer

import (
	"time"

	"gitlab.com/fillip/oki/log"
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

// Compose composes a new `Cluster` piece across the infrastructure
// providers.
func (clusterComposer ClusterComposer) Compose() {
	cluster, _ := clusterComposer.createPrimaryCluster()

	if cluster != nil {
		clusterComposer.attachStorageToCluster(cluster)
		log.Infof("Storage attached to %s cluster\n", cluster.Name)
		//_ = clusterComposer.destroyPrimaryCluster()
		//fmt.Printf("%s cluster deleted.\n", cluster.Name)
	}
}

// CreatePrimaryCluster creates a primary cluster for the service.
func (clusterComposer ClusterComposer) createPrimaryCluster() (*Cluster, error) {
	droplet := &do.Droplet{
		Name: "eu-cluster-1.fillip.pro",
		Tags: []string{"cluster"},
	}

	droplet, err := clients.DigitalOcean.CreateDroplet(droplet)

	log.Infof("Cluster '%d' created!\n", droplet.ID)

	if err != nil {
		log.Error(err)
	}

	cluster := &Cluster{
		ID:   droplet.ID,
		Name: droplet.Name,
		Tags: droplet.Tags,
	}

	return cluster, err
}

// DestroyPrimaryCluster destroys the primary cluster.
func (clusterComposer ClusterComposer) destroyPrimaryCluster() error {
	err := clients.DigitalOcean.DeleteDropletsByTag("cluster")

	return err
}

// AttachStorageToCluster attaches volume storage to a cluster.
func (clusterComposer ClusterComposer) attachStorageToCluster(cluster *Cluster) {
	var id string

	for {
		volumes, err := clients.DigitalOcean.ListVolumes()

		if len(volumes) > 0 && err == nil {
			for _, volume := range volumes {
				if volume.Name == "eu-volume-1-fillip-pro" {
					id = volume.ID
					break
				}
			}
		}

		if len(id) > 0 {
			break
		}

		time.Sleep(3000 * time.Millisecond)
	}

	clients.DigitalOcean.AttachVolumeToDroplet(id, cluster.ID)
}
