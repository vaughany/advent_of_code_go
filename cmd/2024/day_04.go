package main

import (
	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day04() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day04part1,
		cfg.day04part2,
	)
}

// 2024-04-1: 2454
func (cfg *config) day04part1(grid []string) (total int) {
	maxX := len(grid[0])
	maxY := len(grid)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if string(grid[y][x]) == `X` || string(grid[y][x]) == `S` {
				total += day3Vertical(grid, x, y, maxY)
				total += day3Horizontal(grid, x, y, maxX)
				total += day3DiagonalUp(grid, x, y, maxX)
				total += day3DiagonalDown(grid, x, y, maxX, maxY)
			}
		}
	}

	return
}

func day3Vertical(grid []string, x, y, maxY int) int {
	if y+4 > maxY {
		return 0
	}

	guess := string(grid[y][x]) + string(grid[y+1][x]) + string(grid[y+2][x]) + string(grid[y+3][x])
	if guess == `XMAS` || guess == `SAMX` {
		return 1
	}

	return 0
}

func day3Horizontal(grid []string, x, y, maxX int) int {
	if x+4 > maxX {
		return 0
	}

	guess := string(grid[y][x]) + string(grid[y][x+1]) + string(grid[y][x+2]) + string(grid[y][x+3])
	if guess == `XMAS` || guess == `SAMX` {
		return 1
	}

	return 0
}

func day3DiagonalUp(grid []string, x, y, maxX int) int {
	if x+4 > maxX || y-3 < 0 {
		return 0
	}

	guess := string(grid[y][x]) + string(grid[y-1][x+1]) + string(grid[y-2][x+2]) + string(grid[y-3][x+3])
	if guess == `XMAS` || guess == `SAMX` {
		return 1
	}

	return 0
}

func day3DiagonalDown(grid []string, x, y, maxX, maxY int) int {
	if x+4 > maxX || y+4 > maxY {
		return 0
	}

	guess := string(grid[y][x]) + string(grid[y+1][x+1]) + string(grid[y+2][x+2]) + string(grid[y+3][x+3])
	if guess == `XMAS` || guess == `SAMX` {
		return 1
	}

	return 0
}

// 2024-04-2: 1858
func (cfg *config) day04part2(grid []string) (total int) {
	maxX := len(grid[0])
	maxY := len(grid)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			total += day3X(grid, x, y, maxX, maxY)
		}
	}

	return
}

func day3X(grid []string, x, y, maxX, maxY int) int {
	if x+3 > maxX || y+3 > maxY {
		return 0
	}

	guess := string(grid[y][x]) + string(grid[y][x+2]) + string(grid[y+1][x+1]) + string(grid[y+2][x]) + string(grid[y+2][x+2])
	if guess == `MMASS` || guess == `MSAMS` || guess == `SSAMM` || guess == `SMASM` {
		return 1
	}

	return 0
}
