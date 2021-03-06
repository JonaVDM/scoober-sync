package gotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Gotify holds the server data needed to send stuff
type Gotify struct {
	URL   string
	Token string
}

// Send will make a good attemt to send a notication, might return an error tho
func (g *Gotify) Send(title, message string, priority int) error {
	if g.URL == "" || g.Token == "" {
		return errors.New("Gotify Url or token undefined")
	}

	values := map[string]interface{}{
		"title":    title,
		"message":  message,
		"priority": priority,
	}

	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://"+g.URL+"/message", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Add("X-Gotify-Key", g.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Gotify Returned status code %d", res.StatusCode)
	}

	return nil
}
