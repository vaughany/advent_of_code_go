package output

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

type colourInfoType string

const (
	colourTitle  colourInfoType = "\u001b[32m"
	colourInfo   colourInfoType = "\u001b[33m"
	colourTiming colourInfoType = "\u001b[36m"
	colourReset  colourInfoType = "\u001b[0m"
	colourNYI    colourInfoType = "\u001b[90m"
)

func sendToLog(colour colourInfoType, contents string) {
	fmt.Printf("%s%s%s\n", colour, contents, colourReset)
}

func answer(part string, output any) {
	fmt.Printf("Part %s: %v\n", part, output)
}

func Title(year int) {
	sendToLog(colourTitle, fmt.Sprintf("Advent of Code %d.", year))
}

func Subtitle(day int) {
	sendToLog(colourTitle, fmt.Sprintf("Day %d.", day))
}

func AnswerPart1(output any) {
	answer("One", output)
}

func AnswerPart2(output any) {
	answer("Two", output)
}

func Info(text string) {
	sendToLog(colourInfo, text)
}

func NotYetImplemented() {
	sendToLog(colourNYI, "Not yet implemented")
}

func TimeInfo(timeType TimeInfoType, timeDuration time.Duration) {
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

	sendToLog(colourTiming, fmt.Sprintf("%s%s", output, timeDuration))
}
