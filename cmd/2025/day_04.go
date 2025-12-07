package main

import (
	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day04() error {
	return runDayWithInput(
		cfg,
		func(l loaders.Loader) ([]string, error) { return l.GetStrings() },
		cfg.day04part1,
		cfg.day04part2,
	)
}

// 2025-04-1: 1445
func (cfg *config) day04part1(instructions []string) int {
	rows := len(instructions)
	cols := len(instructions[0])

	// Build grid of rolls.
	grid := make([][]bool, rows)
	for y := range rows {
		grid[y] = make([]bool, cols)
		for x, char := range instructions[y] {
			if char == '@' {
				grid[y][x] = true
			}
		}
	}

	moveableRolls := 0

	for y := range rows {
		for x := range cols {
			if !grid[y][x] {
				continue
			}

			if canMoveRoll(grid, x, y) {
				moveableRolls++
			}
		}
	}

	return moveableRolls
}

func canMoveRoll(grid [][]bool, x, y int) bool {
	rows := len(grid)
	cols := len(grid[0])

	adjRolls := 0

	for yy := y - 1; yy <= y+1; yy++ {
		for xx := x - 1; xx <= x+1; xx++ {
			// Skip the roll we're on.
			if xx == x && yy == y {
				continue
			}

			// Bounds check, or Go complains *violently*.
			if xx < 0 || xx >= cols || yy < 0 || yy >= rows {
				continue
			}

			if grid[yy][xx] {
				adjRolls++

				// Exit early if we exceed 4.
				if adjRolls >= 4 {
					return false
				}
			}
		}
	}

	return true
}

// 2025-04-2: 8317
func (cfg *config) day04part2(instructions []string) int {
	rows := len(instructions)
	cols := len(instructions[0])

	// Build grid of rolls.
	grid := make([][]bool, rows)
	for y := range rows {
		grid[y] = make([]bool, cols)
		for x, char := range instructions[y] {
			if char == '@' {
				grid[y][x] = true
			}
		}
	}

	totalRemoved := 0

	for {
		moveableRolls := 0
		toClear := make([][2]int, 0) // A slice of 2 integers.

		for y := range rows {
			for x := range cols {
				if !grid[y][x] {
					continue
				}

				if canMoveRoll(grid, x, y) {
					moveableRolls++
					toClear = append(toClear, [2]int{y, x})
				}
			}
		}

		// Remove rolls.
		for _, pos := range toClear {
			y, x := pos[0], pos[1]
			grid[y][x] = false
		}

		totalRemoved += moveableRolls

		if moveableRolls == 0 {
			break
		}
	}

	return totalRemoved
}
