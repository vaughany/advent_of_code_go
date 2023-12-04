package main

import (
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day04(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day04part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day04part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2023-04-1: 25651
func (cfg *config) day04part1(instructions []string) int {
	var score int
	for _, instruction := range instructions {
		split := strings.Split(instruction, ":")
		split2 := strings.Split(split[1], "|")
		winning := strings.Split(split2[0], " ")
		mine := strings.Split(split2[1], " ")

		var matches int

		for _, w := range winning {
			if strings.Trim(w, " ") == "" {
				continue
			}

			for _, m := range mine {
				if strings.Trim(m, " ") == "" {
					continue
				}

				if w == m {
					matches++
				}
			}
		}

		if matches > 0 {
			s := 1
			for i := 1; i < matches; i++ {
				s *= 2
			}
			score += s
		}
	}

	return score
}

// 2023-04-2:
func (cfg *config) day04part2(instructions []string) int {
	return 0
}
