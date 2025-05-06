package competition

import (
	"fmt"
	"math"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitor"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/timehelpers"
)

func preprocessingData(resData *[]competitor.CompetitorResult, settings *competitionsettings.CompetitionValues) {
	for key, value := range competitor.AllCompetitors {
		newRes := competitor.NewCompetitorResult(value.AllTime, value.ExtraInfo, key)
		if newRes.ExtraInfo != "" {
			newRes.AllTime = math.MaxUint32
		}
		lapsTime := make([]string, settings.Laps)
		lapsSpeed := make([]string, settings.Laps)
		for i := 0; i < len(value.Laps); i++ {
			fmt.Println(value.Laps[i])
			speedPerLap := float64(settings.LapLen*1000) / float64(value.Laps[i])
			lapsTime[i] = timehelpers.MilliToTime(value.Laps[i])
			lapsSpeed[i] = timehelpers.SpeedToTime(speedPerLap)
			fmt.Println(lapsTime[i])
			fmt.Println(lapsSpeed[i])
		}
		newRes.LapsTime = lapsTime
		newRes.LapsSpeed = lapsTime
		newRes.PenaltyTime = timehelpers.MilliToTime(value.PenaltyTime)
		fmt.Println(newRes.PenaltyTime)
		penaltySpeed := float64(value.PenaltyAmount*uint32(settings.PenaltyLen)*1000) / float64(value.PenaltyTime*1000)
		newRes.PenaltySpeed = timehelpers.SpeedToTime(penaltySpeed)
		fmt.Println(newRes.PenaltySpeed)
		newRes.ShotsResult = fmt.Sprintf("%d/%d", value.Hits, value.Shots)
		fmt.Println(newRes.ShotsResult)
		*resData = append(*resData, *newRes)
	}
}

func writeResults(resData *[]competitor.CompetitorResult) {

}
