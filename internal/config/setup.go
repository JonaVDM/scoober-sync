package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/jonavdm/scoober-sync/internal/scoober"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Setup will run through all the setup steps
func Setup() error {
	gcl, gtk, err := setupGoogle()
	if err != nil {
		return err
	}

	fmt.Print("\n\n")

	calID, err := setupCalendar(gcl)
	if err != nil {
		return err
	}

	fmt.Print("\n\n")

	scb, err := setupScoober()
	if err != nil {
		return err
	}

	fmt.Print("\n\n")

	config := Config{
		ScooberToken: scb,
		GoogleToken:  gtk,
		CalendarID:   calID,
	}

	err = config.Save()
	if err != nil {
		return err
	}

	return nil
}

func setupGoogle() (*http.Client, *oauth2.Token, error) {
	path := os.Getenv("SCOOBER_CONFIG")
	b, err := ioutil.ReadFile(path + "/credentials.json")
	if err != nil {
		return nil, nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, nil, err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, nil, err
	}

	tok, err := config.Exchange(context.TODO(), code)
	client := config.Client(context.Background(), tok)
	return client, tok, err
}

func setupCalendar(client *http.Client) (string, error) {
	cal, err := calendar.New(client)
	if err != nil {
		return "", err
	}

	list, err := cal.CalendarList.List().Do()
	if err != nil {
		return "", err
	}

	fmt.Println("Please type in the number for the calendar to use. " +
		"Note: this can and will delete events from the calendar!")
	for i, v := range list.Items {
		fmt.Printf(" - [%d] %v\n", i, v.Summary)
	}

	var index string
	if _, err := fmt.Scan(&index); err != nil {
		return "", err
	}

	i, err := strconv.Atoi(index)
	if err != nil {
		return "", err
	}

	selected := list.Items[i]
	fmt.Printf("Using calendar %s (%s)", selected.Summary, selected.Id)

	return selected.Id, nil
}

func setupScoober() (string, error) {
	fmt.Println("Scoober Sign in")

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

	client := scoober.NewScoober()
	err := client.Login(email, password)
	return client.Token, err
}

func setupTiming() {

}
