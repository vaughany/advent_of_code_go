package main

import "testing"

func TestYear2022day01part1(t *testing.T) {
	var (
		cfg          config
		instructions = cfg.getInputAsInts(2022, 1)
		expected     = 70369
		actual       = cfg.year2022day01part1(instructions)
	)

	if expected != actual {
		t.Errorf("TestYear2022day01part1: expected %d, actual %d", expected, actual)
	}
}

func TestYear2022day01part2(t *testing.T) {
	var (
		cfg          config
		instructions = cfg.getInputAsInts(2022, 1)
		expected     = 203002
		actual       = cfg.year2022day01part2(instructions)
	)

	if expected != actual {
		t.Errorf("TestYear2022day01part1: expected %d, actual %d", expected, actual)
	}
}

func TestYear2022day02part1(t *testing.T) {
	var (
		cfg          config
		instructions = cfg.getInputAsStrings(2022, 2)
		expected     = 11841
		actual       = cfg.year2022day02part1(instructions)
	)

	if expected != actual {
		t.Errorf("TestYear2022day02part1: expected %d, actual %d", expected, actual)
	}
}

func TestYear2022day02part2(t *testing.T) {
	var (
		cfg          config
		instructions = cfg.getInputAsStrings(2022, 2)
		expected     = 13022
		actual       = cfg.year2022day02part2(instructions)
	)

	if expected != actual {
		t.Errorf("TestYear2022day02part1: expected %d, actual %d", expected, actual)
	}
}

func TestYear2022day03part1(t *testing.T) {
	var (
		cfg          config
		instructions = cfg.getInputAsStrings(2022, 3)
		expected     = 8018
		actual       = cfg.year2022day03part1(instructions)
	)

	if expected != actual {
		t.Errorf("TestYear2022day03part1: expected %d, actual %d", expected, actual)
	}
}

func TestYear2022day03part2(t *testing.T) {
	var (
		cfg          config
		instructions = cfg.getInputAsStrings(2022, 3)
		expected     = 2518
		actual       = cfg.year2022day03part2(instructions)
	)

	if expected != actual {
		t.Errorf("TestYear2022day03part2: expected %d, actual %d", expected, actual)
	}
}

//

//

//

func BenchmarkYear2022day01part1(b *testing.B) {
	var cfg config

	cfg.year2022day01part1(cfg.getInputAsInts(2022, 1))
}

func BenchmarkYear2022day01part2(b *testing.B) {
	var cfg config

	cfg.year2022day01part2(cfg.getInputAsInts(2022, 1))
}

func BenchmarkYear2022day02part1(b *testing.B) {
	var cfg config

	cfg.year2022day02part1(cfg.getInputAsStrings(2022, 2))
}

func BenchmarkYear2022day02part2(b *testing.B) {
	var cfg config

	cfg.year2022day02part2(cfg.getInputAsStrings(2022, 2))
}

func BenchmarkYear2022day03part1(b *testing.B) {
	var cfg config

	cfg.year2022day03part1(cfg.getInputAsStrings(2022, 3))
}

func BenchmarkYear2022day03part2(b *testing.B) {
	var cfg config

	cfg.year2022day03part2(cfg.getInputAsStrings(2022, 3))
}
