package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jonavdm/scoober-sync/internal/config"
)

// SendWebhook sends a webhook to the discord server
func SendWebhook(message string) error {
	conf, err := config.Load()
	if err != nil {
		return err
	}

	if conf.DiscordHook == "" {
		return errors.New("Webhook URL not provided")
	}

	data, err := json.Marshal(map[string]string{
		"content": message,
	})
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(data)

	http.Post(conf.DiscordHook, "application/json", body)

	return nil
}
