package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Manifest map[string]string

var MalformedManifestError = errors.New("Manifest is malformed")

func LoadManifest() (Manifest, error) {

	manifest := make(Manifest)
	buf, err := os.ReadFile("manifest.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return manifest, nil
		} else {
			return nil, err
		}
	}
	entries := strings.Split(string(buf), "\n")
	for _, entry := range entries {
		items := strings.Split(entry, " ")
		if len(items) != 2 {
			return nil, MalformedManifestError
		}
		manifest[items[0]] = items[1]
	}
	return manifest, nil
}

func (m *Manifest) Save() error {
	var entries []string
	for key, value := range *m {
		entries = append(entries, fmt.Sprintf("%s %s", key, value))
	}
	data := strings.Join(entries[:], "\n")

	err := ioutil.WriteFile("manifest.txt", []byte(data), 0600)
	if err != nil {
		return err
	}

	return nil
}
