package main

import (
	"regexp"
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day02(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day02part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day02part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2023-02-1: 2406
func (cfg *config) day02part1(instructions []string) (totalPossible int) {
	maxRegexes := []struct {
		regex *regexp.Regexp
		max   int
	}{
		{regex: regexp.MustCompile(`\d{1,2} red`), max: 12},
		{regex: regexp.MustCompile(`\d{1,2} green`), max: 13},
		{regex: regexp.MustCompile(`\d{1,2} blue`), max: 14},
	}

	possible := make(map[int]bool, len(instructions))
	for i := 0; i <= len(instructions)-1; i++ {
		possible[i] = true
	}

	for i, instruction := range instructions {
		for _, mr := range maxRegexes {
			matches := mr.regex.FindAllString(instruction, -1)
			for _, match := range matches {
				if day02getInt(match) > mr.max {
					possible[i] = false

					continue
				}
			}
		}
	}

	for i, p := range possible {
		if p == true {
			totalPossible += i + 1
		}
	}

	return
}

// 2023-02-2: 78375
func (cfg *config) day02part2(instructions []string) (sumOfPowers int) {
	reR := regexp.MustCompile(`\d{1,2} red`)
	reG := regexp.MustCompile(`\d{1,2} green`)
	reB := regexp.MustCompile(`\d{1,2} blue`)

	for _, instruction := range instructions {
		red := reR.FindAllString(instruction, -1)

		green := reG.FindAllString(instruction, -1)

		blue := reB.FindAllString(instruction, -1)

		sumOfPowers += day02getMatches(red) * day02getMatches(green) * day02getMatches(blue)
	}

	return
}

func day02getInt(in string) (out int) {
	parts := strings.Split(in, " ")

	out, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	return
}

func day02getMatches(in []string) (max int) {
	for _, match := range in {
		int := day02getInt(match)
		if int > max {
			max = int
		}
	}

	return
}
