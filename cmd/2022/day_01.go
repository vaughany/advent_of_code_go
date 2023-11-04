package main

import (
	"sort"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day01(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetInts(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day01part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day01part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-01-1: 70369
func (cfg *config) day01part1(instructions []int) int {
	var (
		elves        = make(map[int]int, len(instructions))
		elf          int
		calorificElf int
	)

	for _, ins := range instructions {
		elves[elf] += ins

		if ins == 0 {
			if elves[elf] > calorificElf {
				calorificElf = elves[elf]
			}
			elf++
		}
	}

	return calorificElf
}

// 2022-01-2: 203002
func (cfg *config) day01part2(instructions []int) int {
	var (
		elves = make(map[int]int, len(instructions))
		elf   int
	)

	for _, ins := range instructions {
		elves[elf] += ins

		if ins == 0 {
			elf++
		}
	}

	keys := make([]int, 0, len(elves))
	for key := range elves {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return elves[keys[i]] > elves[keys[j]]
	})

	return elves[keys[0]] + elves[keys[1]] + elves[keys[2]]
}
