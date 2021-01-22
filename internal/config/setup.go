package config

import (
	"fmt"

	"github.com/jonavdm/scoober-sync/internal/scoober"
)

// Setup will run through all the setup steps
func Setup() error {
	scb, err := setupScoober()
	if err != nil {
		return err
	}

	config := Config{
		ScooberToken: scb,
	}

	err = config.Save()
	if err != nil {
		return err
	}

	return nil
}

func setupGoogle() {

}

func setupCalendar() {

}

func setupScoober() (string, error) {
	fmt.Println("Scoober Sign in")
	// Get email
	fmt.Println("Enter email:")
	var email string
	if _, err := fmt.Scan(&email); err != nil {
		return "", err
	}

	fmt.Println("Enter password:")
	var password string
	if _, err := fmt.Scan(&password); err != nil {
		return "", err
	}

	fmt.Println("email:", email)
	fmt.Println("password:", password)

	client := scoober.NewScoober()
	err := client.Login(email, password)
	return client.Token, err
}

func setupTiming() {

}
