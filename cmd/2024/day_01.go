package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day01() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day01part1,
		cfg.day01part2,
	)
}

// 2024-01-1: 2192892
func (cfg *config) day01part1(instructions []string) (total int) {
	lefts, rights := cfg.day01ProcessInput(instructions)

	slices.Sort(lefts)
	slices.Sort(rights)

	for i := range lefts {
		diff := lefts[i] - rights[i]
		if diff < 0 {
			diff = -diff
		}

		total += diff
	}

	return
}

// 2024-01-2: 22962826
func (cfg *config) day01part2(instructions []string) (total int) {
	lefts, rights := cfg.day01ProcessInput(instructions)

	// Count occurrences of each right-hand value
	freq := make(map[int]int)
	for _, r := range rights {
		freq[r]++
	}

	// For each left, multiply by how often it appears on the right
	for _, l := range lefts {
		total += l * freq[l]
	}

	return total
}

func (cfg *config) day01ProcessInput(instructions []string) (lefts, rights []int) {
	lefts = make([]int, 0, len(instructions))
	rights = make([]int, 0, len(instructions))

	for _, line := range instructions {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid input line: %q", line))
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	return
}
