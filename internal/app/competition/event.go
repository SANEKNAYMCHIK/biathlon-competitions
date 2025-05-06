package competition

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/timehelpers"
)

type Event struct {
	CurrentTime  timehelpers.FullTime
	EventID      int
	CompetitorID int
	Extra        any
}

func parseLog(line string) (Event, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return Event{}, fmt.Errorf("Not enough parameteres per line")
	}
	// tStr := parts[0]
	t, err := time.Parse(time.TimeOnly, parts[0][1:len(parts[0])-1])
	if err != nil {
		return Event{}, fmt.Errorf("Error of parsing data: %v", err)
	}
	eventID, err := strconv.Atoi(parts[1])
	if err != nil {
		return Event{}, fmt.Errorf("Error of parsing eventID: %v", err)
	}
	competitorID, err := strconv.Atoi(parts[2])
	if err != nil {
		return Event{}, fmt.Errorf("Error of parsing competitorID: %v", err)
	}
	event := Event{
		CurrentTime:  timehelpers.FullTime{t},
		EventID:      eventID,
		CompetitorID: competitorID,
	}
	if len(parts) > 3 {
		if extraInt, err := strconv.Atoi(parts[3]); err == nil {
			event.Extra = extraInt
			return event, nil
		}
		event.Extra = parts[3]
	}
	return event, nil
}
