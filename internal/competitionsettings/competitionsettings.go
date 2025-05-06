package competitionsettings

import "github.com/SANEKNAYMCHIK/biathlon-competitions/internal/timehelpers"

type CompetitionValues struct {
	Laps        int                  `json:"laps"`
	LapLen      int                  `json:"lapLen"`
	PenaltyLen  int                  `json:"penaltyLen"`
	FiringLines int                  `json:"firingLines"`
	Start       timehelpers.FullTime `json:"start"`
	StartDelta  timehelpers.FullTime `json:"startDelta"`
}
