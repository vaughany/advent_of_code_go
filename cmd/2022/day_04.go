package main

import (
	"fmt"
	"strconv"
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
	part1, part2 := cfg.day04bothParts(instructions)
	output.AnswerPart1(part1)
	output.AnswerPart2(part2)

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-04-1: 573
// 2022-04-2: 867
func (cfg *config) day04bothParts(instructions []string) (int, int) {
	var (
		contained, overlapping int
		s1s, s1e, s2s, s2e     int
	)

	for _, ins := range instructions {
		sections := strings.Split(ins, ",")

		s1 := strings.Split(sections[0], "-")
		s2 := strings.Split(sections[1], "-")

		s1s, _ = strconv.Atoi(s1[0])
		s1e, _ = strconv.Atoi(s1[1])
		s2s, _ = strconv.Atoi(s2[0])
		s2e, _ = strconv.Atoi(s2[1])

		if cfg.debug {
			fmt.Println(s1s, s1e, s2s, s2e)
		}

		if s1s <= s2s && s1e >= s2e || s2s <= s1s && s2e >= s1e {
			contained++
		}

		if (s1s >= s2s && s1s <= s2e || s1e >= s2s && s1e <= s2e) || (s1s <= s2s && s1e >= s2e || s2s <= s1s && s2e >= s1e) {
			overlapping++
		}
	}

	return contained, overlapping
}
