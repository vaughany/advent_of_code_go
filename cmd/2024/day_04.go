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

var word = []byte("XMAS")

// 2024-04-1: 2454
func (cfg *config) day04part1(grid []string) (total int) {
	maxY := len(grid)
	maxX := len(grid[0])

	// Directions.
	dirs := [][2]int{
		{1, 0}, {-1, 0},
		{0, 1}, {0, -1},
		{1, 1}, {-1, -1},
		{1, -1}, {-1, 1},
	}

	for y := range maxY {
		row := []byte(grid[y])
		for x := range maxX {
			// Quick filter: only start from X.
			if row[x] != 'X' {
				continue
			}

			for _, d := range dirs {
				if matchWord(grid, x, y, d[0], d[1]) {
					total++
				}
			}
		}
	}

	return
}

func matchWord(grid []string, x, y, dx, dy int) bool {
	maxY := len(grid)
	maxX := len(grid[0])

	for i := range word {
		xx := x + dx*i
		yy := y + dy*i

		if xx < 0 || xx >= maxX || yy < 0 || yy >= maxY {
			return false
		}

		if grid[yy][xx] != word[i] {
			return false
		}
	}

	return true
}

// 2024-04-2: 1858
func (cfg *config) day04part2(grid []string) (total int) {
	maxY := len(grid)
	maxX := len(grid[0])

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			total += day4XMAS(grid, x, y, maxX, maxY)
		}
	}

	return
}

func day4XMAS(grid []string, x, y, maxX, maxY int) int {
	// We’re treating x,y as the top-left of a 3x3 block.
	if x+2 >= maxX || y+2 >= maxY {
		return 0
	}

	// Center must be 'A'.
	if grid[y+1][x+1] != 'A' {
		return 0
	}

	// Diagonal 1: top-left to center to bottom-right.
	d1a := grid[y][x]
	d1b := grid[y+1][x+1]
	d1c := grid[y+2][x+2]

	// Diagonal 2: top-right to center to bottom-left.
	d2a := grid[y][x+2]
	d2b := grid[y+1][x+1]
	d2c := grid[y+2][x]

	diagIsMAS := func(a, b, c byte) bool {
		return (a == 'M' && b == 'A' && c == 'S') || (a == 'S' && b == 'A' && c == 'M')
	}

	if diagIsMAS(d1a, d1b, d1c) && diagIsMAS(d2a, d2b, d2c) {
		return 1
	}

	return 0
}
