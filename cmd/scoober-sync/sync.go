package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/jonavdm/scoober-sync/internal/config"
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
		return
	}

	log.Println(msg)

	if conf.PushoverApp != "" {
		push := pushover.New(conf.PushoverApp)
		recipient := pushover.NewRecipient(conf.PushoverUser)

		message := pushover.NewMessage(msg)

		_, err := push.SendMessage(message, recipient)
		if err != nil {
			log.Print(err)
		}
	}

	time.Sleep(time.Second * 10)
}
