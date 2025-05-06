package main

import (
	"encoding/json"
	"os"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/app/competition"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
)

const configName = "config.json"
const eventsName = "events"
const outputName = "output"
const resultName = "resulting table"

// Парсинг json файла с характеристиками соревнования и вызов функции Battle
func main() {
	data, err := os.ReadFile(configName)
	if err != nil {
		panic(err)
	}
	var settings competitionsettings.CompetitionValues
	err = json.Unmarshal(data, &settings)
	if err != nil {
		panic(err)
	}
	competition.Battle(&settings, eventsName, outputName, resultName)
}
