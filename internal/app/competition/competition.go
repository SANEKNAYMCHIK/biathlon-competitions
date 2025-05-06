package competition

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/SANEKNAYMCHIK/biathlon-competitions/internal/competitionsettings"
)

var actions = map[int]string{
	1:  "The competitor(%v) registered",
	2:  "The start time was set by a draw to %v",
	3:  "The competitor(%v) is on the start line",
	4:  "The competitor(%v) has started",
	5:  "The competitor(%v) is on the firing range(%v)",
	6:  "The target(%v) has been hit by competitor(%v)",
	7:  "The competitor(%v) left the firing range",
	8:  "The competitor(%v) entered the penalty laps",
	9:  "The competitor(%v) left the penalty laps",
	10: "The competitor(%v) ended the main lap",
	11: "The competitor(%v) can`t continue: %v",
	32: "The competitor(%v) is disqualified",
	33: "The competitor(%v) has finished",
}

func Battle(settings *competitionsettings.CompetitionValues, eventsName string, outputName string) {
	file, err := os.Open(eventsName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	out, err := os.Create(outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer out.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		eventVals, err := parseLog(line)
		if err != nil {
			fmt.Printf("Incorrect log: %s\n", err)
			continue
		}
		fmt.Println(eventVals.EventID)
		fmt.Println(eventVals.EventID)
		fmt.Println(eventVals.CompetitorID)
		fmt.Println(eventVals.Extra)
		fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
