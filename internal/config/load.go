package config

import (
	"encoding/json"
	"io/ioutil"
)

// Load will load in the config file and return a config object
func Load() (*Config, error) {
	folder, err := getConfigFolder()
	if err != nil {
		return nil, err
	}

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
