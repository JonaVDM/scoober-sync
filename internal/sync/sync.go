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
func Sync() (*Log, error) {
	conf, err := config.Load()
	if err != nil {
		return nil, err
	}

	monday, sunday := getDates()

	scb := scoober.NewScoober()
	scb.Token = conf.ScooberToken

	shifts, err := scb.GetShifts(monday.Format("2006-01-02"), sunday.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	google, err := config.GetGoogleConfig()
	if err != nil {
		return nil, err
	}

	calendar, err := googleCal.New(google.Client(context.Background(), conf.GoogleToken))
	if err != nil {
		return nil, err
	}

	events, err := calendar.Events.List(conf.CalendarID).ShowDeleted(false).
		SingleEvents(true).TimeMin(monday.Format(time.RFC3339)).
		TimeMax(sunday.Format(time.RFC3339)).OrderBy("startTime").Do()

	if err != nil {
		return nil, err
	}

	log := newLog()

	for _, event := range events.Items {
		id := strings.Split(event.Description, "\n")[0]
		if id == "" {
			continue
		}

		shift := findShift(&shifts, id)
		if shift == nil {
			if err := calendar.Events.Delete(conf.CalendarID, event.Id).Do(); err != nil {
				return nil, err
			}

			sStart, _ := time.Parse(time.RFC3339, event.Start.DateTime)
			sEnd, _ := time.Parse(time.RFC3339, event.End.DateTime)

			log.Deleted = append(log.Deleted, fmt.Sprintf("%s - %s", sStart.Format(time.Stamp), sEnd.Format(time.Stamp)))
		}
	}

	for _, shift := range shifts {
		event := findEvent(events, shift.ID)

		sStart, _ := time.Parse(time.RFC3339, shift.FromWithTimeZone)
		sEnd, _ := time.Parse(time.RFC3339, shift.ToWithTimeZone)

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

			log.Added = append(log.Added, fmt.Sprintf("%s - %s", sStart.Format(time.Stamp), sEnd.Format(time.Stamp)))

			continue
		}

		eStart, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		eEnd, _ := time.Parse(time.RFC3339, event.End.DateTime)

		if eStart != sStart || eEnd != sEnd {
			event.Start = &googleCal.EventDateTime{DateTime: shift.From}
			event.End = &googleCal.EventDateTime{DateTime: shift.To}
			calendar.Events.Update(conf.CalendarID, event.Id, event).Do()

			log.Updated = append(log.Updated, fmt.Sprintf("%s - %s", sStart.Format(time.Stamp), sEnd.Format(time.Stamp)))
		}
	}

	return log, err
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

func getDates() (time.Time, time.Time) {
	now := time.Now()

	dayOfWeek := int(now.Weekday())
	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	monday := now.AddDate(0, 0, -dayOfWeek+1)
	sunday := now.AddDate(0, 0, 7-dayOfWeek+7)

	monday = time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
	sunday = time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 23, 59, 59, 0, sunday.Location())

	return monday, sunday
}
