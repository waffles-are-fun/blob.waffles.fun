package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Token string `json:"token"`
}

var fn string
var path string

func setup() error {
	path, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	path = fmt.Sprintf("%s/waffles", path)
	err = os.MkdirAll(path, 0700)
	if err != nil {
		return err
	}

	fn = fmt.Sprintf("%s/config.json", path)

	return nil
}

func (c *Config) Save() error {
	err := setup()
	if err != nil {
		return err
	}

	conf, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fn, conf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() (*Config, error) {

	var conf Config

	err := setup()
	if err != nil {
		return nil, err
	}

	buf, err := os.ReadFile(fn)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &conf, nil
		}
		return nil, err
	}

	err = json.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
