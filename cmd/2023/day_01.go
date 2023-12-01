package main

import (
	"regexp"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day01(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day01part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day01part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2023-01-1: 54990
func (cfg *config) day01part1(instructions []string) (totalCal int) {
	re := regexp.MustCompile(`(\d)`)

	for _, instruction := range instructions {
		matches := re.FindAllString(instruction, -1)

		totalCal += getInteger(matches[0])*10 + getInteger(matches[len(matches)-1])
	}

	return
}

// 2023-01-2: 54473
func (cfg *config) day01part2(instructions []string) (totalCal int) {
	re := regexp.MustCompile(`(\d|oneight|twone|eightwo|one|two|three|four|five|six|seven|eight|nine)`)

	for _, instruction := range instructions {
		var thisCal int

		matches := re.FindAllString(instruction, -1)
		first, last := matches[0], matches[len(matches)-1]

		switch first {
		case "oneight":
			thisCal = 10
		case "twone":
			thisCal = 20
		case "eightwo":
			thisCal = 80
		default:
			thisCal = 10 * getInteger(first)
		}

		switch last {
		case "oneight":
			thisCal += 8
		case "twone":
			thisCal++
		case "eightwo":
			thisCal += 2
		default:
			thisCal += getInteger(last)
		}

		totalCal += thisCal
	}

	return
}

func getInteger(in string) int {
	switch in {
	case "one", "1":
		return 1
	case "two", "2":
		return 2
	case "three", "3":
		return 3
	case "four", "4":
		return 4
	case "five", "5":
		return 5
	case "six", "6":
		return 6
	case "seven", "7":
		return 7
	case "eight", "8":
		return 8
	case "nine", "9":
		return 9
	case "0":
		panic("found a zero")
	default:
		panic("found something else")
	}
}
