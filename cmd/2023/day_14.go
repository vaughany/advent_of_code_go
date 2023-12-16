package main

import (
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day14(loader loaders.Loader) error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStringsGrid(loader)
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day14part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day14part2(instructions))

	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2023-14-1: 108813
func (cfg *config) day14part1(instructions [][]string) (result int) {
	width := len(instructions[0])
	height := len(instructions)

	rock := "O"
	empty := "."

	for {
		var moves int

		for y := 0; y <= height-1; y++ {
			for x := 0; x <= width-1; x++ {
				if instructions[y][x] == rock && y > 0 && instructions[y-1][x] == empty {
					instructions[y-1][x] = rock
					instructions[y][x] = empty
					moves++
				}
			}
		}

		if moves == 0 {
			break
		}
	}

	for y := 0; y <= height-1; y++ {
		for x := 0; x <= width-1; x++ {
			if instructions[y][x] == rock {
				result += height - y
			}
		}
	}

	if cfg.debug {
		displayGrid(instructions)
	}

	return
}

// 2023-14-2: 104533
// Laura solved it on a calculator...
func (cfg *config) day14part2(instructions [][]string) (result int) {
	return
}
