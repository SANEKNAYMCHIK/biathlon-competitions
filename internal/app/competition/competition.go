package competition

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitor"
)

// Шаблоны ответов на события
var actions = map[int]string{
	1:  "%s The competitor(%v) registered\n",
	2:  "%s The start time for the competitor(%v) was set by a draw to %v\n",
	3:  "%s The competitor(%v) is on the start line\n",
	4:  "%s The competitor(%v) has started\n",
	5:  "%s The competitor(%v) is on the firing range(%v)\n",
	6:  "%s The target(%v) has been hit by competitor(%v)\n",
	7:  "%s The competitor(%v) left the firing range\n",
	8:  "%s The competitor(%v) entered the penalty laps\n",
	9:  "%s The competitor(%v) left the penalty laps\n",
	10: "%s The competitor(%v) ended the main lap\n",
	11: "%s The competitor(%v) can`t continue: %v\n",
	32: "%s The competitor(%v) is disqualified\n",
	33: "%s The competitor(%v) has finished\n",
}

// Вывод шаблонных ответов в форматированном варианте
func writeOutputLog(eventVals *Event, out *os.File) {
	if eventVals.Extra != "" {
		if eventVals.EventID == 6 {
			fmt.Fprintf(out, actions[eventVals.EventID], eventVals.CurrentTime,
				eventVals.Extra, eventVals.CompetitorID)
		} else {
			fmt.Fprintf(out, actions[eventVals.EventID], eventVals.CurrentTime,
				eventVals.CompetitorID, eventVals.Extra)
		}
	} else {
		fmt.Fprintf(out, actions[eventVals.EventID], eventVals.CurrentTime,
			eventVals.CompetitorID)
	}
}

// Чтение логов соревнования, вывод в обработанном формате
// Запись необходимых значений, вывод исходящих событий, результатов
func Battle(settings *competitionsettings.CompetitionValues, eventsName string, outputName string, resultName string) {
	file, err := os.Open(eventsName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	out, err := os.Create(outputName)
	if err != nil {
		log.Fatal(err)
	}
	res, err := os.Create(resultName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer out.Close()
	defer res.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		eventVals, err := parseLog(line)
		if err != nil {
			fmt.Printf("Incorrect log: %s\n", err)
			continue
		}
		writeOutputLog(&eventVals, out)
		response := CheckLog(eventVals, settings)
		if response != nil {
			writeOutputLog(response, out)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	var resData []competitor.CompetitorResult
	preprocessingData(&resData, settings)
	sort.Slice(resData, func(i, j int) bool {
		if resData[i].AllTime != resData[j].AllTime {
			return resData[i].AllTime < resData[j].AllTime
		}
		return resData[i].ID < resData[j].ID
	})
	writeResults(&resData, res)
}
