package scoober

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type loginResponse struct {
	Token string `json:"accessToken"`
}

var token string

// Login signs you into your scoober account and gives back the
// token, the token is also stored in the state of the application
func Login(email string, password string) (string, error) {
	fmt.Println(email)
	fmt.Println(password)

	postBody, err := json.Marshal(map[string]string{
		"userName": email,
		"password": password,
	})

	reqBody := bytes.NewBuffer(postBody)

	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://shiftplanning-api.scoober.com/login", "application/json", reqBody)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sb := string(body)
	data := loginResponse{}
	err = json.Unmarshal([]byte(sb), &data)
	if err != nil {
		return "", err
	}

	return data.Token, nil
}
