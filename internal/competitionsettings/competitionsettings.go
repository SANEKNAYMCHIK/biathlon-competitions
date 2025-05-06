package competitionsettings

import "time"

type FullTime struct {
	time.Time
}

type CompetitionValues struct {
	Laps        int      `json:"laps"`
	LapLen      int      `json:"lapLen"`
	PenaltyLen  int      `json:"penaltyLen"`
	FiringLines int      `json:"firingLines"`
	Start       FullTime `json:"start"`
	StartDelta  FullTime `json:"startDelta"`
}

func (t *FullTime) UnmarshalJSON(data []byte) error {
	strTime := string(data[1 : len(data)-1])
	parsedTime, err := time.Parse(time.TimeOnly, strTime)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
