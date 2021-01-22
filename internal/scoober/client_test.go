package scoober_test

import (
	"testing"

	. "github.com/jonavdm/scoober-sync/internal/scoober"
)

func TestNewScoober(t *testing.T) {
	scoober := NewScoober()

	if scoober.Token != "" {
		t.Error("Token has to be empty")
	}

	if scoober.BaseURL != "https://shiftplanning-api.scoober.com" {
		t.Error("The url is not right")
	}
}
