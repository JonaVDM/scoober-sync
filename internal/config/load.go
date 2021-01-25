package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load will load in the config file and return a config object
func Load() (*Config, error) {
	folder := os.Getenv("SCOOBER_CONFIG")

	data, err := ioutil.ReadFile(folder + "/config.json")
	if err != nil {
		return nil, err
	}

	conf := Config{}
	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
