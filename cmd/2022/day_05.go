package main

import (
	"fmt"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day05(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day05part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day05part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-05-1: CFFHVVHNC
func (cfg *config) day05part1(instructions []string) string {
	const cutset = `[] 123456789`

	var (
		tmpStacks = map[int][]string{}
		stacks    = map[int][]string{}
		setup     = true
	)

	for _, ins := range instructions {
		// Parse the first bit of the input file.
		if setup {
			if ins == "" {
				setup = false

				for j := 1; j <= len(tmpStacks); j++ {
					for k := len(tmpStacks[j]) - 1; k >= 0; k-- {
						if tmpStacks[j][k] != "" {
							stacks[j] = append(stacks[j], tmpStacks[j][k])
						}
					}
				}

				continue
			}

			for j := 0; j < (len(ins)+1)/4; j++ {
				tmpStacks[j+1] = append(tmpStacks[j+1], strings.Trim(ins[j*4:j*4+3], cutset))
			}
		}

		var count, from, to int
		fmt.Sscanf(ins, "move %2d from %2d to %2d", &count, &from, &to)

		if count == 0 {
			continue
		}

		for j := 1; j <= count; j++ {
			// Get the package
			pack := stacks[from][len(stacks[from])-1:][0]

			// Remove it from the 'old' stack
			stacks[from] = stacks[from][:len(stacks[from])-1]

			// Add it to the 'new' stack
			stacks[to] = append(stacks[to], pack)
		}
	}

	// Get top-most package from each stack
	var out string
	for j := 1; j <= len(stacks); j++ {
		out += stacks[j][len(stacks[j])-1:][0]
	}

	return out
}

// 2022-05-2: FSZWBPTBG
func (cfg *config) day05part2(instructions []string) string {
	const cutset = `[] 123456789`

	var (
		tmpStacks = map[int][]string{}
		stacks    = map[int][]string{}
		setup     = true
	)

	for _, ins := range instructions {
		// Parse the first bit of the input file.
		if setup {
			if ins == "" {
				setup = false

				for j := 1; j <= len(tmpStacks); j++ {
					for k := len(tmpStacks[j]) - 1; k >= 0; k-- {
						if tmpStacks[j][k] != "" {
							stacks[j] = append(stacks[j], tmpStacks[j][k])
						}
					}
				}

				continue
			}

			for j := 0; j < (len(ins)+1)/4; j++ {
				tmpStacks[j+1] = append(tmpStacks[j+1], strings.Trim(ins[j*4:j*4+3], cutset))
			}
		}

		var count, from, to int
		fmt.Sscanf(ins, "move %2d from %2d to %2d", &count, &from, &to)

		if count == 0 {
			continue
		}

		// Get the package/s
		pack := stacks[from][len(stacks[from])-count:]

		// Remove it/them from the 'old' stack
		stacks[from] = stacks[from][:len(stacks[from])-count]

		// Add it/them to the 'new' stack
		stacks[to] = append(stacks[to], pack...)
	}

	// Get top-most package from each stack
	var out string
	for j := 1; j <= len(stacks); j++ {
		out += stacks[j][len(stacks[j])-1:][0]
	}

	return out
}
