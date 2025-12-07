package main

import (
	"fmt"
	"strconv"

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

// 2025-01-1: 1158
func (cfg *config) day01part1(instructions []string) int {
	var (
		position  = 50
		totalZero = 0
	)

	for _, instruction := range instructions {
		click, err := strconv.Atoi(instruction[1:])
		if err != nil {
			fmt.Println("ERROR: string to int conversion failed for instruction:", instruction)
			return 0
		}

		// Update position based on direction
		if instruction[0] == 'L' {
			position -= click
		} else {
			position += click
		}

		// Use existing helper function for consistency and potential optimization
		position = cfg.day01mod100(position)

		if position == 0 {
			totalZero++
		}

		if cfg.debug {
			fmt.Println("position:", position)
		}
	}

	return totalZero
}

// 2025-01-2: 6860
func (cfg *config) day01part2(instructions []string) int {
	var (
		position  = 50
		totalZero = 0
	)

	for _, instruction := range instructions {
		click, err := strconv.Atoi(instruction[1:])
		if err != nil {
			fmt.Println("ERROR: string to int conversion failed for instruction:", instruction)
			return 0
		}

		// Log the position before the change.
		low := position

		// Update position based on direction
		if instruction[0] == 'L' {
			position -= click
		} else {
			position += click
		}

		// Log the position after the change.
		high := position

		// The main logic.
		passes := cfg.day01passesBetween(low, high)

		// If we land right on 0, log that.
		land := 0
		if cfg.day01mod100(position) == 0 {
			land = 1
		}

		// Add all the things.
		totalZero += passes + land

		if cfg.debug {
			fmt.Println("position:", cfg.day01mod100(position), "totalZero:", totalZero)
		}
	}

	return totalZero
}

// Absolute modulus function.
func (cfg *config) day01mod100(in int) int {
	return ((in % 100) + 100) % 100
}

// Figure out the passes, but drop 1 off the higher number to ensure we don't over-count them.
func (cfg *config) day01passesBetween(low, high int) int {
	// Ensure high is the higher of the two, and low the lowest.
	if low > high {
		low, high = high, low
	}
	hiInner := high - 1

	return max(cfg.day01floorDiv(hiInner, 100)-cfg.day01floorDiv(low, 100), 0)
}

// day01floorDiv does mathematical floor division (towards -∞).
func (cfg *config) day01floorDiv(in, d int) int {
	q := in / d // Truncates toward zero.
	r := in % d // Remainder.

	// If in is negative and there's a remainder, truncating division moved us UP toward zero.
	// Floor division should move us DOWN by 1.
	if r != 0 && ((in < 0) != (d < 0)) {
		q--
	}

	return q
}
