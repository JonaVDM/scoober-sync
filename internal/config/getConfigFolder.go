package config

import "os"

func getConfigFolder() (string, error) {
	folder := os.Getenv("SCOOBER_CONFIG")

	if folder == "" {
		return os.Getwd()
	}

	return folder, nil
}
