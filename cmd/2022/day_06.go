package main

import (
	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day06(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetString(loader)
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

// 2022-06-1: 1175
func (cfg *config) day06part1(instructions string) int {
	var (
		window = 4
		out    int
	)

	for j := 0; j <= len(instructions)-window; j++ {
		slice := instructions[j : j+window]

		chars := make(map[byte]int, window)
		for i := 0; i < window; i++ {
			chars[slice[i]]++
		}

		if len(chars) == window {
			out = j + window

			break
		}
	}

	return out
}

// 2022-06-2: 3217
func (cfg *config) day06part2(instructions string) int {
	var (
		window = 14
		out    int
	)

	for j := 0; j <= len(instructions)-window; j++ {
		slice := instructions[j : j+window]

		chars := make(map[byte]int, window)
		for i := 0; i < window; i++ {
			chars[slice[i]]++
		}

		if len(chars) == window {
			out = j + window

			break
		}
	}

	return out
}
