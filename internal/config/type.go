package config

import "golang.org/x/oauth2"

// Config holds all the applications config
type Config struct {
	ScooberToken string
	GoogleToken  *oauth2.Token
}
