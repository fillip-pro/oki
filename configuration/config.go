package configuration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// CloudConfig produces a base cloud-config.yml based on the initial template.
func CloudConfig() {
	inPath, err := filepath.Abs("configuration/cloud-config.yml")
	outPath, err := filepath.Abs("configuration/new-cloud-config.yml")

	read, err := ioutil.ReadFile(inPath)
	if err != nil {
		panic(err)
	}

	newContents := string(read)

	systemdPath, err := filepath.Abs("configuration/systemd")
	files, _ := ioutil.ReadDir(systemdPath)
	for _, file := range files {
		systemdFile, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", systemdPath, file.Name()))
		newContents = strings.Replace(newContents, fmt.Sprintf("{{%%%s}}", file.Name()), string(systemdFile), -1)
	}

	discovery := &Discovery{}
	discoveryURL, err := discovery.DiscoveryURL()
	newContents = strings.Replace(newContents, "{{$etcd-discovery-token}}", discoveryURL, -1)

	err = ioutil.WriteFile(outPath, []byte(newContents), 0644)
	if err != nil {
		panic(err)
	}
}

func etcdDiscovery() string {
	return "bobsack"
}
