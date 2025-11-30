package main

import (
	"fmt"
	"strconv"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day10(loader loaders.Loader) error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetString(loader)
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

// 2015-10-1:
func (cfg *config) day01part1(instruction string) int {
	iterations := 40

	input := []byte(instruction)
	for i := 1; i <= iterations; i++ {
		input = lookAndSay(input)

		if cfg.debug {
			fmt.Println("iteration", i, "len", len(input))
		}
	}

	return len(input)
}

// 2015-10-2:
func (cfg *config) day01part2(instruction string) int {
	iterations := 50

	input := []byte(instruction)
	for i := 1; i <= iterations; i++ {
		input = lookAndSay(input)

		if cfg.debug {
			fmt.Println("iteration", i, "len", len(input))
		}
	}

	return len(input)
}

func lookAndSay(in []byte) []byte {
	out := make([]byte, 0, len(in)*2)

	for i := 0; i < len(in); {
		j := i + 1
		for j < len(in) && in[j] == in[i] {
			j++
		}

		count := j - i
		// Append decimal count and the digit itself
		out = strconv.AppendInt(out, int64(count), 10)
		out = append(out, in[i])
		i = j
	}

	return out
}
