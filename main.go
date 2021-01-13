package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jonavdm/scoober-sync/scoober"
)

func main() {
	fmt.Println("hi")

	email := os.Getenv("SCOOBER_EMAIL")
	password := os.Getenv("SCOOBER_PASSWORD")

	token, err := scoober.Login(email, password)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(token)
}
