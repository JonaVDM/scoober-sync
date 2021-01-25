package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jonavdm/scoober-sync/internal/scoober"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Setup will run through all the setup steps
func Setup() error {
	gtk, err := setupGoogle()
	if err != nil {
		return err
	}

	scb, err := setupScoober()
	if err != nil {
		return err
	}

	config := Config{
		ScooberToken: scb,
		GoogleToken:  gtk,
	}

	err = config.Save()
	if err != nil {
		return err
	}

	return nil
}

func setupGoogle() (*oauth2.Token, error) {
	path := os.Getenv("SCOOBER_CONFIG")
	b, err := ioutil.ReadFile(path + "/credentials.json")
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		return nil, err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	tok, err := config.Exchange(context.TODO(), code)
	return tok, err
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
