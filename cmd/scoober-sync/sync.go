package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jonavdm/scoober-sync/internal/config"
	"github.com/jonavdm/scoober-sync/internal/gotify"
	"github.com/jonavdm/scoober-sync/internal/sync"
)

func main() {
	changes, err := sync.Sync()
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var messages []string

	for _, added := range changes.Added {
		messages = append(messages, fmt.Sprintf("Added %s", added))
	}

	for _, deleted := range changes.Deleted {
		messages = append(messages, fmt.Sprintf("Removed %s", deleted))
	}

	for _, updated := range changes.Updated {
		messages = append(messages, fmt.Sprintf("Updated %s", updated))
	}

	msg := strings.Join(messages, "\n")

	if msg == "" {
		log.Println("No changed have been made")
		return
	}

	log.Println(msg)

	gf := gotify.Gotify{
		URL:   conf.GotifyURL,
		Token: conf.GotifyToken,
	}

	if err := gf.Send("Updated Schedule", msg, 5); err != nil {
		log.Print(err)
	}

	time.Sleep(time.Second * 10)
}
