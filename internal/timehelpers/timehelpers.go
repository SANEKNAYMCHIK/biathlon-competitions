package timehelpers

import (
	"fmt"
	"math"
	"time"
)

// Свой встроенный тип данных времени
type FullTime struct {
	time.Time
}

// Реализация интерфейса Stringer для форматированного вывода времени
func (t FullTime) String() string {
	return fmt.Sprintf("[%s]", t.Format("15:04:05.000"))
}

// Функция, переводящая строку со временем в тип FullTime
func ToTime(s string) *FullTime {
	parsedTime, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		panic(err)
	}
	return &FullTime{parsedTime}
}

// Перевод из миллисекунд в строку формата HH:MM:SS.sss
func MilliToTime(val uint32) string {
	hours := val / 3600000
	val -= hours * 3600000
	minutes := val / 60000
	val -= minutes * 60000
	seconds := val / 1000
	val -= seconds * 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, val)
}

// Форматирование вывода скорости
func SpeedToTime(val float64) string {
	return fmt.Sprintf("%.3f", math.Trunc(val*1000)/1000)
}

// Метод, переводящий FullTime в миллисекунды
func (t FullTime) ToMilli() uint32 {
	var res uint32 = 0
	res += uint32(t.Hour() * 60 * 60 * 1000)
	res += uint32(t.Minute() * 60 * 1000)
	res += uint32(t.Second() * 1000)
	res += uint32(t.Nanosecond() / 1e6)
	return res
}

// Переопределение функции Unmarshal для типа FullTime
func (t *FullTime) UnmarshalJSON(data []byte) error {
	strTime := string(data[1 : len(data)-1])
	parsedTime, err := time.Parse(time.TimeOnly, strTime)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
