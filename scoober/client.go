package scoober

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Scoober is the main scoober account object
type Scoober struct {
	Token string
}

// NewScoober lets you login to your account
func NewScoober(email string, password string) (*Scoober, error) {
	postBody, err := json.Marshal(map[string]string{
		"userName": email,
		"password": password,
	})

	reqBody := bytes.NewBuffer(postBody)

	if err != nil {
		return &Scoober{}, err
	}

	resp, err := http.Post("https://shiftplanning-api.scoober.com/login", "application/json", reqBody)
	if err != nil {
		return &Scoober{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Scoober{}, err
	}

	type loginResponse struct {
		Token string `json:"accessToken"`
	}

	sb := string(body)
	data := loginResponse{}
	err = json.Unmarshal([]byte(sb), &data)
	if err != nil {
		return &Scoober{}, err
	}

	return &Scoober{
		data.Token,
	}, nil
}
