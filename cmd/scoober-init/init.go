package main

import (
	"fmt"
	"log"

	"github.com/jonavdm/scoober-sync/internal/config"
)

func main() {
	fmt.Println("Scoober Init")

	err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("saved the config")
}
