package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Save will store the config in a config file in the config folder
func (c *Config) Save() error {
	folder := os.Getenv("SCOOBER_CONFIG")
	if folder == "" {
		var err error
		folder, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(folder+"/config.json", data, 0644)

	return err
}
