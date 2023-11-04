package main

import (
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day10(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day10part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day10part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-10-1: 15680
func (cfg *config) day10part1(instructions []string) int {
	var (
		out    int
		cycles []int
		x      = 1
	)

	cycles = append(cycles, x)

	for _, ins := range instructions {
		if strings.HasPrefix(ins, "noop") {
			cycles = append(cycles, x)

		} else {
			cycles = append(cycles, x)
			value, _ := strconv.Atoi(strings.Split(ins, " ")[1])
			x += value
			cycles = append(cycles, x)
		}
	}

	for _, x := range []int{20, 60, 100, 140, 180, 220} {
		out += cycles[x-1] * x
	}

	return out
}

// 2022-10-2: ZFBFHGUP
func (cfg *config) day10part2(instructions []string) int {
	var (
		cycles     []int
		x          = 1
		screen     []string
		screenLine string
	)

	cycles = append(cycles, x)

	for _, ins := range instructions {
		if strings.HasPrefix(ins, "noop") {
			cycles = append(cycles, x)

		} else {
			cycles = append(cycles, x)
			value, _ := strconv.Atoi(strings.Split(ins, " ")[1])
			x += value
			cycles = append(cycles, x)
		}
	}

	for j := 0; j < len(cycles); j++ {
		sprite := len(screenLine)

		if sprite == cycles[j] || sprite == cycles[j]+1 || sprite == cycles[j]-1 {
			screenLine += "#"
		} else {
			screenLine += " "
		}

		if len(screenLine) == 40 {
			screen = append(screen, screenLine)
			screenLine = ""
		}
	}

	output.Info(strings.Join(screen, "\n"))

	return 0
}
