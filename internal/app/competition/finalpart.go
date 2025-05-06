package competition

import (
	"fmt"
	"math"
	"os"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitor"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/timehelpers"
)

// Предобработка данных, отдельно будет функция вывода, отдельно предпосчет всех скоростей и сохранение в нужном формате
func preprocessingData(resData *[]competitor.CompetitorResult, settings *competitionsettings.CompetitionValues) {
	for key, value := range competitor.AllCompetitors {
		newRes := competitor.NewCompetitorResult(value.AllTime, value.ExtraInfo, key)
		if newRes.ExtraInfo != "" {
			newRes.AllTime = math.MaxUint32
		}
		lapsTime := make([]string, settings.Laps)
		for i := 0; i < settings.Laps; i++ {
			lapsTime[i] = ""
		}
		lapsSpeed := make([]string, settings.Laps)
		for i := 0; i < settings.Laps; i++ {
			lapsSpeed[i] = ""
		}
		for i := 0; i < len(value.Laps); i++ {
			var speedPerLap float64
			if value.Laps[i] != 0 {
				speedPerLap = float64(settings.LapLen*1000) / float64(value.Laps[i])
			} else {
				speedPerLap = 0.000
			}
			lapsTime[i] = timehelpers.MilliToTime(value.Laps[i])
			lapsSpeed[i] = timehelpers.SpeedToTime(speedPerLap)
		}
		newRes.LapsTime = lapsTime
		newRes.LapsSpeed = lapsSpeed
		newRes.PenaltyTime = timehelpers.MilliToTime(value.PenaltyTime)
		var penaltySpeed float64
		if value.PenaltyTime != 0 {
			penaltySpeed = float64(value.PenaltyAmount*uint32(settings.PenaltyLen)*1000) / float64(value.PenaltyTime)
		} else {
			penaltySpeed = 0.000
		}
		newRes.PenaltySpeed = timehelpers.SpeedToTime(penaltySpeed)
		newRes.ShotsResult = fmt.Sprintf("%d/%d", value.Hits, value.Shots)
		*resData = append(*resData, *newRes)
	}
}

// Конвертация слайсов времени, потраченного на круг и средней скорости за круг в одну строку
func valsToTuples(timeVals []string, speedVals []string) string {
	result := ""
	for i := 0; i < len(timeVals); i++ {
		result += "{"
		result += timeVals[i]
		if timeVals[i] != "" {
			result += ", "
		} else {
			result += ","
		}
		result += speedVals[i]
		result += "}"
		if i != len(timeVals)-1 {
			result += ", "
		}
	}
	return result
}

// Запись итоговых значений в результирующий файл
func writeResults(resData *[]competitor.CompetitorResult, out *os.File) {
	for i := 0; i < len(*resData); i++ {
		var firstArg string
		if (*resData)[i].ExtraInfo != "" {
			firstArg = (*resData)[i].ExtraInfo
		} else {
			firstArg = timehelpers.MilliToTime((*resData)[i].AllTime)
		}
		dataPerLaps := valsToTuples((*resData)[i].LapsTime, (*resData)[i].LapsSpeed)
		fmt.Fprintf(out, "[%s] %d [%s] {%s, %s} %s\n",
			firstArg, (*resData)[i].ID, dataPerLaps, (*resData)[i].PenaltyTime,
			(*resData)[i].PenaltySpeed, (*resData)[i].ShotsResult)
	}
}
