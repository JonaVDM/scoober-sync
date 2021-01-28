package main

import (
	"log"

	"github.com/jonavdm/scoober-sync/internal/sync"
)

func main() {
	changes, err := sync.Sync()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(changes)
}
