package competition

import (
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitor"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/timehelpers"
)

// Обработка входящих событий
// Имитация соревнования, запись всех необходимых значений в структуру Competitor в map по ID участника
func CheckLog(eventVals Event, settings *competitionsettings.CompetitionValues) *Event {
	switch eventVals.EventID {
	case 1:
		competitor.AllCompetitors[eventVals.CompetitorID] = competitor.NewCompetitor()
	case 2:
		competitor.AllCompetitors[eventVals.CompetitorID].ScheduledStart = timehelpers.ToTime(eventVals.Extra).ToMilli()
	case 4:
		competitor.AllCompetitors[eventVals.CompetitorID].ActualStart = eventVals.CurrentTime.ToMilli()
		competitor.AllCompetitors[eventVals.CompetitorID].PrevStart = eventVals.CurrentTime.ToMilli()
		if (competitor.AllCompetitors[eventVals.CompetitorID].ActualStart - competitor.AllCompetitors[eventVals.CompetitorID].ScheduledStart) > settings.StartDelta.ToMilli() {
			competitor.AllCompetitors[eventVals.CompetitorID].ExtraInfo = "NotStarted"
			return &Event{
				CurrentTime:  eventVals.CurrentTime,
				EventID:      32,
				CompetitorID: eventVals.CompetitorID,
			}
		}
	case 5:
		competitor.AllCompetitors[eventVals.CompetitorID].Shots += 5
	case 6:
		competitor.AllCompetitors[eventVals.CompetitorID].Hits += 1
	case 8:
		competitor.AllCompetitors[eventVals.CompetitorID].PenaltyStart = eventVals.CurrentTime.ToMilli()
	case 9:
		penaltyPerLap := eventVals.CurrentTime.ToMilli() - competitor.AllCompetitors[eventVals.CompetitorID].PenaltyStart
		competitor.AllCompetitors[eventVals.CompetitorID].PenaltyTime += penaltyPerLap
		competitor.AllCompetitors[eventVals.CompetitorID].PenaltyAmount += 1
	case 10:
		timePerLap := eventVals.CurrentTime.ToMilli() - competitor.AllCompetitors[eventVals.CompetitorID].PrevStart
		competitor.AllCompetitors[eventVals.CompetitorID].Laps = append(competitor.AllCompetitors[eventVals.CompetitorID].Laps, timePerLap)
		if settings.Laps == len(competitor.AllCompetitors[eventVals.CompetitorID].Laps) {
			diff := eventVals.CurrentTime.ToMilli() - competitor.AllCompetitors[eventVals.CompetitorID].ScheduledStart
			competitor.AllCompetitors[eventVals.CompetitorID].AllTime = diff
			return &Event{
				CurrentTime:  eventVals.CurrentTime,
				EventID:      33,
				CompetitorID: eventVals.CompetitorID,
			}
		}
	default:
		competitor.AllCompetitors[eventVals.CompetitorID].ExtraInfo = "NotFinished"
	}
	return nil
}
