package main

import (
	"encoding/json"
	"os"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/app/competition"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
	"github.com/joho/godotenv"
)

// Получение имен файлов, откуда читать данные и куда выводить, из .env файла
// Парсинг json файла с характеристиками соревнования и вызов функции Battle
func main() {
	godotenv.Load()
	configName := os.Getenv("CONFIG_NAME")
	eventsName := os.Getenv("EVENTS_NAME")
	outputName := os.Getenv("OUTPUT_LOGS_NAME")
	resultName := os.Getenv("RESULT_TABLE_NAME")
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
