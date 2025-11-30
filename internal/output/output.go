package output

import (
	"fmt"
	"time"
)

// TimeInfoType represents the type of timing information to display.
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
	colourTitle  colourInfoType = "\u001b[32m" // Green
	colourInfo   colourInfoType = "\u001b[33m" // Yellow
	colourTiming colourInfoType = "\u001b[36m" // Cyan
	colourReset  colourInfoType = "\u001b[0m"  // Reset
	colourNYI    colourInfoType = "\u001b[90m" // Dark gray
)

func sendToLog(colour colourInfoType, contents string) {
	fmt.Printf("%s%s%s\n", colour, contents, colourReset)
}

func sendToLogF(colour colourInfoType, format string, args ...any) {
	fmt.Printf("%s", colour)
	fmt.Printf(format, args...)
	fmt.Printf("%s\n", colourReset)
}

func answer(part string, output any) {
	fmt.Printf("Part %s: %v\n", part, output)
}

// Title displays the Advent of Code year title in green.
func Title(year int) {
	sendToLogF(colourTitle, "Advent of Code %d.", year)
}

// Subtitle displays the day subtitle in green.
func Subtitle(day int) {
	sendToLogF(colourTitle, "Day %d.", day)
}

// AnswerPart1 displays the answer for part one.
func AnswerPart1(output any) {
	answer("One", output)
}

// AnswerPart2 displays the answer for part two.
func AnswerPart2(output any) {
	answer("Two", output)
}

// Info displays informational text in yellow.
func Info(text string) {
	sendToLog(colourInfo, text)
}

// NotYetImplemented displays a "Not yet implemented" message in dark gray.
func NotYetImplemented() {
	sendToLog(colourNYI, "Not yet implemented")
}

// TimeInfo displays timing information based on the timeType in cyan.
// If an unknown timeType is provided, it defaults to "Timing: ".
func TimeInfo(timeType TimeInfoType, timeDuration time.Duration) {
	var prefix string

	switch timeType {
	case InfoTypeSetup:
		prefix = "Setup took "
	case InfoTypeOne:
		prefix = "Part One took "
	case InfoTypeTwo:
		prefix = "Part Two took "
	case InfoTypeBoth:
		prefix = "Both Parts took "
	case InfoTypeEverything:
		prefix = "Everything took "
	case InfoTypeWholeRun:
		prefix = "\nThe whole run took "
	default:
		prefix = "Timing: "
	}

	sendToLogF(colourTiming, "%s%s", prefix, timeDuration)
}
