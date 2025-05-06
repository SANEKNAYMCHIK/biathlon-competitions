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

func ToTime(s string) *FullTime {
	parsedTime, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		panic(err)
	}
	return &FullTime{parsedTime}
}

func (t FullTime) ToMilli() uint32 {
	var res uint32 = 0
	res += uint32(t.Hour() * 60 * 60 * 1000)
	res += uint32(t.Minute() * 60 * 1000)
	res += uint32(t.Second() * 1000)
	res += uint32(t.Nanosecond() / 1e6)
	return res
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
