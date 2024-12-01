package main

import (
	"math"
	"slices"
	"strconv"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day01(loader loaders.Loader) error {
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
	output.AnswerPart1(cfg.day01part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day01part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2024-01-1: 2192892
func (cfg *config) day01part1(instructions []string) (total int) {
	lefts, rights := cfg.day01ProcessInput(instructions)

	slices.Sort(lefts)
	slices.Sort(rights)

	for j := 0; j < len(lefts); j++ {
		total += int(math.Abs(float64(lefts[j]) - float64(rights[j])))
	}

	return
}

// 2024-01-2: 22962826
func (cfg *config) day01part2(instructions []string) (total int) {
	lefts, rights := cfg.day01ProcessInput(instructions)

	for _, left := range lefts {
		count := 0
		for _, right := range rights {
			if left == right {
				count++
			}
		}

		total += left * count
	}

	return total
}

func (cfg *config) day01ProcessInput(instructions []string) (lefts, rights []int) {
	// Different slice indexes depending on sample or real puzzle input.
	// Better than a regex here 'cos the input digits are a consistent length and separation throughout the input files.
	leftEnd, rightStart := 1, 4
	if !cfg.sample {
		leftEnd, rightStart = 5, 8
	}

	for _, ins := range instructions {
		left, err := strconv.Atoi(ins[:leftEnd])
		if err != nil {
			panic(err)
		}
		lefts = append(lefts, left)

		right, err := strconv.Atoi(ins[rightStart:])
		if err != nil {
			panic(err)
		}
		rights = append(rights, right)
	}

	return
}
