package main

import (
	"math"
	"strconv"
	"strings"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day02(loader loaders.Loader) error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day02part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day02part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2024-02-1: 510
func (cfg *config) day02part1(instructions []string) (safe int) {
	for _, ins := range instructions {
		report := day2ProcessInput(ins)

		if day2AllIncOrDec(report) && day2DiffOneToThree(report) {
			safe++
		}
	}

	return
}

// 2024-02-2: 553
func (cfg *config) day02part2(instructions []string) (safe int) {
	for _, ins := range instructions {
		report := day2ProcessInput(ins)

		if !day2AllIncOrDec(report) || !day2DiffOneToThree(report) {
			if day2CheckLevels(report) {
				safe++
			}
		} else {
			safe++
		}
	}

	return
}

func day2ProcessInput(in string) (out []int) {
	for _, s := range strings.Split(in, ` `) {
		int, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		out = append(out, int)
	}

	return
}

func day2AllIncOrDec(in []int) bool {
	greater, lesser := 0, 0

	for i := 0; i < len(in)-1; i++ {
		if in[i] > in[i+1] {
			lesser++
		} else if in[i] < in[i+1] {
			greater++
		} else {
			return false
		}
	}

	if greater > 0 && lesser > 0 {
		return false
	}

	return true
}

func day2DiffOneToThree(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		diff := math.Abs(float64(in[i]) - float64(in[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func day2CheckLevels(in []int) bool {
	for i := 0; i < len(in); i++ {
		var report []int

		// Make a new slice with the item at element i missing.
		for j := 0; j < len(in); j++ {
			if i != j {
				report = append(report, in[j])
			}
		}

		if day2AllIncOrDec(report) && day2DiffOneToThree(report) {
			return true
		}
	}

	return false
}
