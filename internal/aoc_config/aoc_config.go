package aoc_config

import (
	"fmt"
)

const (
	Name        = "Advent of Code"
	Version     = "0.0.1"
	BuildDate   = "2023-11-03"
	Description = "Every year, for the period of Advent, there are 25 two-part Christmassy puzzles to solve: https://adventofcode.com/. This is my attempt, in Go."
)

func GetNameWithYear(year int) string {
	return fmt.Sprintf("%s %d", Name, year)
}

func GetVersionInfo(year int) string {
	return fmt.Sprintf("%s by Paul Vaughan. https://github.com/vaughany/advent_of_code_go\nVersion %s, built %s.\n\n%s", GetNameWithYear(year), Version, BuildDate, Description)
}
