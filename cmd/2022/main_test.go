package main

import (
	"testing"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func TestDay01(t *testing.T) {
	var (
		cfg           config
		part1expected = 70369
		part2expected = 203002
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 1, false)
	instructions, err := loaders.GetInts(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day01part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay01part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day01part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay01part2: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay02(t *testing.T) {
	var (
		cfg           config
		part1expected = 11841
		part2expected = 13022
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 2, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day02part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay02part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day02part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay02part2: expected %d, actual %d", part2expected, part2actual)
	}
}
func TestDay03(t *testing.T) {
	var (
		cfg           config
		part1expected = 8018
		part2expected = 2518
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 3, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day03part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay03part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day03part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay03part2: expected %d, actual %d", part2expected, part2actual)
	}
}
