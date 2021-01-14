package scoober

import (
	"net/http"
)

// Scoober is the main scoober account object
type Scoober struct {
	Token   string
	Client  *http.Client
	BaseURL string
}

// NewScoober creates a new scoober object
func NewScoober() *Scoober {
	return &Scoober{
		Client:  &http.Client{},
		BaseURL: "https://shiftplanning-api.scoober.com",
	}
}
