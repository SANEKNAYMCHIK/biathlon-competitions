package timehelpers

import (
	"fmt"
	"time"
)

type FullTime struct {
	time.Time
}

func (t FullTime) String() string {
	return fmt.Sprintf("[%s]", t.Format("15:04:05.000"))
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
