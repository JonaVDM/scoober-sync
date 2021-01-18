package scoober

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetShifts will make a request and return the shift for
// the period in time, with great detail
// Both the starting and end time must be formatted as
// yyyy-mm-dd. I don't actually check this in here beacause
// I am very lazy, and just didn't want to write code to
// check it. I assume you know what you are doing tho.
func (s *Scoober) GetShifts(start, end string) ([]Shift, error) {
	// TODO: Add a check for the start and end date
	if start == "" || end == "" {
		return nil, errors.New("The start and end date must not be empty")
	}

	if s.Token == "" {
		return nil, errors.New("Client is not logged in")
	}

	req, err := http.NewRequest("GET", s.BaseURL+"/api/users/plannings", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("fromDate", start)
	q.Add("toDate", end)
	req.Header.Add("accessToken", s.Token)
	req.URL.RawQuery = q.Encode()

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := []Shift{}
	sb := string(body)
	err = json.Unmarshal([]byte(sb), &data)

	return data, nil
}
