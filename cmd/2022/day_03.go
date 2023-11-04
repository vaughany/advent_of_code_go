package main

import (
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day03(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day03part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day03part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-03-1: 8018
func (cfg *config) day03part1(instructions []string) int {
	var score int

	for _, ins := range instructions {
		half1, half2 := ins[:len(ins)/2], ins[len(ins)/2:]

		for _, letter := range half1 {
			if strings.Contains(half2, string(letter)) {
				asciiOfRune := int(letter)
				if asciiOfRune >= 97 && asciiOfRune <= 122 {
					score += asciiOfRune - 96
				} else if asciiOfRune >= 65 && asciiOfRune <= 90 {
					score += asciiOfRune - 38
				}
				break
			}
		}
	}

	return score
}

// 2022-03-2: 2518
func (cfg *config) day03part2(instructions []string) int {
	var (
		elfGroup   = make([]string, 3)
		elfInGroup int
		score      int
	)

	for _, ins := range instructions {
		elfGroup[elfInGroup] = ins

		elfInGroup++
		if elfInGroup == 3 {
			elfInGroup = 0

			for _, letter := range elfGroup[0] {
				if strings.Contains(elfGroup[1], string(letter)) {
					if strings.Contains(elfGroup[2], string(letter)) {
						asciiOfRune := int(letter)
						if asciiOfRune >= 97 && asciiOfRune <= 122 {
							score += asciiOfRune - 96
						} else if asciiOfRune >= 65 && asciiOfRune <= 90 {
							score += asciiOfRune - 38
						}

						break
					}
				}
			}
		}
	}

	return score
}
