package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
)

func main() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var settings competitionsettings.CompetitionValues
	err = json.Unmarshal(data, &settings)
	if err != nil {
		panic(err)
	}
	fmt.Println(settings.Laps)
	fmt.Println(settings.LapLen)
	fmt.Println(settings.PenaltyLen)
	fmt.Println(settings.FiringLines)
	fmt.Println(settings.Start.Hour())
	fmt.Println(settings.Start.Nanosecond() / 1e6)
	fmt.Println(settings.StartDelta.Minute())
	fmt.Println(settings.StartDelta.Nanosecond() / 1e6)
}
