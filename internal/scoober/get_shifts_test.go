package scoober_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/jonavdm/scoober-sync/internal/scoober"
)

func TestGetShifts(t *testing.T) {
	successResponse, _ := ioutil.ReadFile("testdata/get_shifts.json")
	failTo, _ := ioutil.ReadFile("testdata/get_shifts_missing_to.json")
	scoober := Scoober{
		Client: &http.Client{},
		Token:  "aabbcc",
	}

	t.Run("Fail when the server is not reachable", func(t *testing.T) {
		shifts, err := scoober.GetShifts("2021-01-01", "2021-01-10")
		scoober.BaseURL = ""

		if err == nil {
			t.Fatal("The error is empty")
		}

		if shifts != nil {
			t.Fatal("The shifts list is not empty")
		}
	})

	t.Run("It request with the right URL", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/api/users/plannings" {
				log.Print(req.URL.String())
				t.Fatal("API-Request path is not right")
			}

			if req.URL.RawQuery != "fromDate=2021-01-01&toDate=2021-01-01" {
				t.Fatal("Query is not correctly formed")
			}

			rw.Write(successResponse)
		}))
		defer server.Close()

		scoober.BaseURL = server.URL
		scoober.GetShifts("2021-01-01", "2021-01-01")

		scoober.BaseURL = ""
	})

	t.Run("Fail when the token is not set", func(t *testing.T) {
		tmpToken := scoober.Token
		scoober.Token = ""

		_, err := scoober.GetShifts("2021-01-01", "2021-01-01")

		if err == nil {
			t.Fatal("The error is not correct")
		}

		scoober.Token = tmpToken
	})

	t.Run("Fail when either of the args is empty", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write(failTo)
		}))
		defer server.Close()

		scoober.BaseURL = server.URL
		shifts, err := scoober.GetShifts("2021-01-01", "")
		if shifts != nil {
			t.Fatal("Shifts is not empty")
		}

		if err == nil {
			t.Fatal("The error is not correct")
		}

		scoober.BaseURL = ""
	})
}
