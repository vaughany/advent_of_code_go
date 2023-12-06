package main

import (
	"regexp"
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day06(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day06part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day06part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2023-06-1: 303600
func (cfg *config) day06part1(instructions []string) (result int) {
	re := regexp.MustCompile(`\d{1,4}`)

	timesStr := re.FindAllString(instructions[0], -1)
	distancesStr := re.FindAllString(instructions[1], -1)

	var times, distances []int
	for i := range timesStr {
		t, _ := strconv.Atoi(timesStr[i])
		times = append(times, t)

		t, _ = strconv.Atoi(distancesStr[i])
		distances = append(distances, t)
	}

	var allWins []int
	for i, time := range times {
		var wins int
		for j := 1; j <= time-1; j++ {
			if ((time - j) * j) > distances[i] {
				wins++
			}
		}

		allWins = append(allWins, wins)
	}

	for _, win := range allWins {
		if result == 0 {
			result = win
		} else {
			result *= win
		}
	}

	return
}

// 2023-06-2: 23654842
func (cfg *config) day06part2(instructions []string) (wins int) {
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(instructions[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(instructions[1], ":")[1], " ", ""))

	for j := 1; j <= time-1; j++ {
		if ((time - j) * j) > distance {
			wins++
		}
	}

	return
}
