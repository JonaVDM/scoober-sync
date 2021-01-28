package main

import (
	"log"

	"github.com/jonavdm/scoober-sync/internal/sync"
)

func main() {
	if err := sync.Sync(); err != nil {
		log.Fatal(err)
	}

	// email := flag.String("email", os.Getenv("SCOOBER_EMAIL"), "The email of your scoober account")
	// password := flag.String("password", os.Getenv("SCOOBER_PASSWORD"), "The password of your scoober account")
	// flag.Parse()

	// if *email == "" || *password == "" {
	// 	log.Fatal("Email or password not defined")
	// }

	// client, err := googleapi.GetClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = calendar.New(client)
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Calendar client: %v", err)
	// }

	// Scoober := scoober.NewScoober()
	// err = Scoober.Login(*email, *password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
