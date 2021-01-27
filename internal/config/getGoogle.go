package config

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// GetGoogleConfig returns the google config object with config loaded
func GetGoogleConfig() (*oauth2.Config, error) {
	path, err := getConfigFolder()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(path + "/credentials.json")
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(b, calendar.CalendarScope)
}
