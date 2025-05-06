package competition

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/timehelpers"
)

// Структура, хранящая в себе все нужные значения из входящих событий
type Event struct {
	CurrentTime  timehelpers.FullTime
	EventID      int
	CompetitorID int
	Extra        string
}

// Парсинг входящих логов и возврат структуры Event с записанными данными
func parseLog(line string) (Event, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return Event{}, fmt.Errorf("Not enough parameteres per line")
	}
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
		event.Extra = parts[3]
	}
	return event, nil
}
