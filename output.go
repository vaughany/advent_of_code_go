package main

import (
	"fmt"
	"time"
)

type TimeInfoType int

const (
	InfoTypeSetup TimeInfoType = iota
	InfoTypeOne
	InfoTypeTwo
	InfoTypeBoth
	InfoTypeEverything
	InfoTypeWholeRun
)

func (cfg *config) title(year, day int) {
	if day == 0 {
		sendToLog(cfg.colours.title, fmt.Sprintf("Advent of Code %d.", year))
		return
	}

	sendToLog(cfg.colours.title, fmt.Sprintf("Advent of Code %d, Day %d.", year, day))
}

func (cfg *config) answerPart1(output int) {
	answer("One", output)
}

func (cfg *config) answerPart2(output int) {
	answer("Two", output)
}

func answer(part string, output int) {
	fmt.Printf("Part %s: %d\n", part, output)
}

func (cfg *config) info(text string) {
	sendToLog(cfg.colours.info, text)
}

func (cfg *config) timeInfo(timeType TimeInfoType, timeDuration time.Duration) {
	if !cfg.timing {
		return
	}

	var output string

	switch timeType {
	case InfoTypeSetup:
		output = "Setup took "
	case InfoTypeOne:
		output = "Part One took "
	case InfoTypeTwo:
		output = "Part Two took "
	case InfoTypeBoth:
		output = "Both Parts took "
	case InfoTypeEverything:
		output = "Everything took "
	case InfoTypeWholeRun:
		output = "\nThe whole run took "
	}

	sendToLog(cfg.colours.timing, fmt.Sprintf("%s%s", output, timeDuration))
}

func sendToLog(colour, contents string) {
	fmt.Printf("%s%s%s\n", colour, contents, "\u001b[0m")
}

// func AnswerString(part int, output interface{}) {
// 	fmt.Printf("Part %s: %s\n", getPartText(part), output)
// }
