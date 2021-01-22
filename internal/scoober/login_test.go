package scoober_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/jonavdm/scoober-sync/internal/scoober"
)

func TestLogin(t *testing.T) {
	failResponse, _ := ioutil.ReadFile("testdata/login_fail.json")
	successResponse, _ := ioutil.ReadFile("testdata/login.json")

	t.Run("Fail when the server is not reachable", func(t *testing.T) {
		scoober := Scoober{
			BaseURL: "",
			Client:  &http.Client{},
		}

		err := scoober.Login("test@mail.com", "password")

		if err == nil {
			t.Fatal("The error is empty")
		}
	})

	t.Run("Signed in with valid info", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/login" {
				t.Fatal("API-Request path is not right")
			}

			rw.Write([]byte(successResponse))
		}))
		defer server.Close()

		scoober := Scoober{
			BaseURL: server.URL,
			Client:  &http.Client{},
		}

		err := scoober.Login("test@mail.com", "password")

		if err != nil {
			t.Fatal("The error is not empty")
		}

		if scoober.Token == "" {
			t.Fatal("The token was not loaded")
		}
	})

	t.Run("Not sign in with in-valid info", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/login" {
				t.Fatal("API-Request path is not right")
			}

			rw.Write([]byte(failResponse))
		}))
		defer server.Close()

		scoober := Scoober{
			BaseURL: server.URL,
			Client:  &http.Client{},
		}

		err := scoober.Login("test@mail.com", "password")

		if err == nil {
			t.Fatal("The error is empty")
		}

		if scoober.Token != "" {
			t.Fatal("The token was not loaded")
		}
	})
}
