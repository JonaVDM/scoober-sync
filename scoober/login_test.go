package scoober_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/jonavdm/scoober-sync/scoober"
)

func TestLogin(t *testing.T) {
	failResponse := `{
    "message": "not authorized",
    "response": {
        "header": {
            "success": false,
            "dateTime": "Thu Jan 14 16:05:04 GMT 2021",
            "errorCodes": [
                "2003",
                "The username or password is invalid"
            ]
        }
    },
    "status": 401
	}
	`
	successResponse := `{
    "userId": 30123,
    "firstName": "first name",
    "lastName": "last name",
    "email": "test@mail.com",
    "userName": "test@mail.com",
    "regionPrefix": "nl:groningen",
    "jobFilter": "FOOD_DELIVERY",
    "accessToken": "aabbcc",
    "accountName": null,
    "accountType": 1,
    "street": "Steentilstraat",
    "houseNumber": "42",
    "zipCode": "9711 GP",
    "city": "groningen",
    "phoneNumber": "0031612345678",
    "isWorking": false,
    "paused": false,
    "countryCode": "nl",
    "chain": null,
    "contractType": 0
	}`

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
