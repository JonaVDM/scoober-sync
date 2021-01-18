package main

import (
	"flag"
	"log"
	"os"

	"github.com/jonavdm/scoober-sync/scoober"
)

func main() {
	email := flag.String("email", os.Getenv("SCOOBER_EMAIL"), "The email of your scoober account")
	password := flag.String("password", os.Getenv("SCOOBER_PASSWORD"), "The password of your scoober account")
	flag.Parse()

	if *email == "" || *password == "" {
		log.Fatal("Email or password not defined")
	}

	Scoober := scoober.NewScoober()
	err := Scoober.Login(*email, *password)
	if err != nil {
		log.Fatal(err)
	}

	shifts, err := Scoober.GetShifts("2021-01-01", "2021-01-15")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(shifts[0].Date)
}
