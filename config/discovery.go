package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.com/fillip/oki/log"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// Discovery provides the features necessary for setting
// up discovery services, like etcd.
type Discovery struct {
}

// DiscoveryURL returns a new URL token for an etcd cluster.
func (discovery Discovery) DiscoveryURL() (string, error) {
	request, err := http.NewRequest("GET", "https://discovery.etcd.io/new?size=1", nil)

	r, err := myClient.Do(request)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	discoveryURL := string(body)

	return discoveryURL, err
}

func getJSON(url string, target interface{}) error {
	request, err := http.NewRequest("GET", url, nil)

	r, err := myClient.Do(request)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	readableBody := string(body)

	log.Infoln(readableBody)

	return json.NewDecoder(r.Body).Decode(target)
}
