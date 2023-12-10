package main

import (
	"slices"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day09(loader loaders.Loader) error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.Get2DInts(loader)
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day09part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day09part2(instructions))

	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2023-09-1: 1882395907
func (cfg *config) day09part1(instructions [][]int) (result int) {
	for _, instructionSlice := range instructions {
		instruction := make(map[int][]int)

		instruction[0] = append(instruction[0], instructionSlice...)

		j := 0
		for {
			for i := 0; i <= len(instruction[j])-1; i++ {
				if i < len(instruction[j])-1 {
					diff := instruction[j][i+1] - instruction[j][i]

					instruction[j+1] = append(instruction[j+1], diff)
				}
			}

			counts := make(map[int]int)
			for _, k := range instruction[j+1] {
				counts[k]++
			}
			if len(counts) == 1 && counts[0] > 0 {
				break
			}
			j++
		}

		for j := len(instruction) - 2; j >= 0; j-- {
			added := instruction[j][len(instruction[j])-1] + instruction[j+1][len(instruction[j+1])-1]
			instruction[j] = append(instruction[j], added)
		}

		result += instruction[0][len(instruction[0])-1]
	}

	return
}

// 2023-09-2: 1005
func (cfg *config) day09part2(instructions [][]int) (result int) {
	// Reverse the slices, and pass the instructions to day 1.
	for i := range instructions {
		slices.Reverse(instructions[i])
	}

	return cfg.day09part1(instructions)
}
