package aoc_config

import (
	"fmt"
)

const (
	Name        = "Advent of Code"
	Version     = "0.0.2"
	BuildDate   = "2025-11-30"
	Description = `Every year, for the period of Advent, there are 25 two-part Christmassy puzzles to solve (12 from 2025 onwards): https://adventofcode.com/. This is my attempt, in Go.`
	Author      = "Paul Vaughan"
	RepoURL     = "https://github.com/vaughany/advent_of_code_go"
)

func GetNameWithYear(year int) string {
	return fmt.Sprintf("%s %d", Name, year)
}

func GetVersionInfo(year int) string {
	name := Name
	if year > 0 {
		name = GetNameWithYear(year)
	}
	return fmt.Sprintf(
		`%s by %s. %s
Version %s, built %s.

%s`,
		name, Author, RepoURL, Version, BuildDate, Description,
	)
}
