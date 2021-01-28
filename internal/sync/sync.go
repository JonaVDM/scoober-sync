package sync

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jonavdm/scoober-sync/internal/config"
	"github.com/jonavdm/scoober-sync/internal/scoober"
	googleCal "google.golang.org/api/calendar/v3"
)

// Sync will sync scoober with google
func Sync() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}

	now := time.Now()

	dayOfWeek := int(now.Weekday())
	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	monday := now.AddDate(0, 0, -dayOfWeek+1)
	sunday := now.AddDate(0, 0, 7-dayOfWeek+7)

	scb := scoober.NewScoober()
	scb.Token = conf.ScooberToken

	shifts, err := scb.GetShifts(monday.Format("2006-01-02"), sunday.Format("2006-01-02"))
	if err != nil {
		return err
	}

	fmt.Println(shifts)

	google, err := config.GetGoogleConfig()
	if err != nil {
		return err
	}

	calendar, err := googleCal.New(google.Client(context.Background(), conf.GoogleToken))
	if err != nil {
		return err
	}

	events, err := calendar.Events.List(conf.CalendarID).ShowDeleted(false).
		SingleEvents(true).TimeMin(monday.Format(time.RFC3339)).
		TimeMax(sunday.Format(time.RFC3339)).OrderBy("startTime").Do()

	if err != nil {
		return err
	}

	for _, event := range events.Items {
		id := strings.Split(event.Description, "\n")[0]
		if id == "" {
			continue
		}

		shift := findShift(&shifts, id)
		if shift == nil {
			if err := calendar.Events.Delete(conf.CalendarID, event.Id).Do(); err != nil {
				return err
			}
		}
	}

	for _, shift := range shifts {
		event := findEvent(events, shift.ID)

		if event == nil {
			calendar.Events.Insert(
				conf.CalendarID,
				&googleCal.Event{
					Description: shift.ID,
					Summary:     "Work Thuisbezorgd",
					Start:       &googleCal.EventDateTime{DateTime: shift.From},
					End:         &googleCal.EventDateTime{DateTime: shift.To},
				},
			).Do()
		}
	}

	return nil
}

func findEvent(events *googleCal.Events, shiftID string) *googleCal.Event {
	for _, event := range events.Items {
		id := strings.Split(event.Description, "\n")[0]
		if id == "" {
			continue
		}

		if id == shiftID {
			return event
		}
	}

	return nil
}

func findShift(shifts *[]scoober.Shift, id string) *scoober.Shift {
	for _, shift := range *shifts {
		if shift.ID == id {
			return &shift
		}
	}

	return nil
}
