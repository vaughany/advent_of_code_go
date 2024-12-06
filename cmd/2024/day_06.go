package main

import (
	"fmt"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day06(loader loaders.Loader) error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day06part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day06part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2024-06-1: 4665
func (cfg *config) day06part1(instructions []string) (total int) {
	const (
		up, down, left, right = `u`, `d`, `l`, `r`
		obstacle              = `#`
		traversed             = `x`
		guard                 = `g`
		space                 = `.`
		startGuard            = `^`
	)

	var (
		guardX, guardY = 0, 0
		maxX           = len(instructions[0])
		maxY           = len(instructions)
		grid           = make([][]string, maxY)

		// Initial direction.
		dir = `u`
	)

	// Initialise grid.
	grid = make([][]string, maxY)
	for i := 0; i < maxY; i++ {
		grid[i] = make([]string, maxX)
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			grid[y][x] = string(instructions[y][x])

			// Guard's start position: we don't care about direction.
			if grid[y][x] == `^` {
				guardX = x
				guardY = y

				grid[y][x] = guard
			}
		}
	}

	outOfBounds := false
	for {
		switch dir {
		case up:
			if guardY-1 < 0 {
				grid[guardY][guardX] = traversed
				outOfBounds = true
			} else if grid[guardY-1][guardX] == obstacle {
				dir = right
			} else {
				grid[guardY][guardX] = traversed
				guardY--
				grid[guardY][guardX] = guard
			}

		case right:
			if guardX+1 >= maxX {
				grid[guardY][guardX] = traversed
				outOfBounds = true
			} else if grid[guardY][guardX+1] == obstacle {
				dir = down
			} else {
				grid[guardY][guardX] = traversed
				guardX++
				grid[guardY][guardX] = guard
			}

		case down:
			if guardY+1 >= maxY {
				grid[guardY][guardX] = traversed
				outOfBounds = true
			} else if grid[guardY+1][guardX] == obstacle {
				dir = left
			} else {
				grid[guardY][guardX] = traversed
				guardY++
				grid[guardY][guardX] = guard
			}

		case left:
			if guardX-1 < 0 {
				grid[guardY][guardX] = traversed
				outOfBounds = true
			} else if grid[guardY][guardX-1] == obstacle {
				dir = up
			} else {
				grid[guardY][guardX] = traversed
				guardX--
				grid[guardY][guardX] = guard
			}
		}

		if outOfBounds {
			break
		}

	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[y][x] == traversed {
				total++
			}
		}
	}

	if cfg.debug {
		day06DrawGrid(grid)
	}

	return
}

func day06DrawGrid(grid [][]string) {
	for _, gr := range grid {
		for _, gc := range gr {
			fmt.Print(gc)
		}
		fmt.Println()
	}
	fmt.Println()
}

type day06cell struct {
	path  string
	count int
}

// 2024-06-2: 1688
func (cfg *config) day06part2(instructions []string) (total int) {
	const (
		up, down, left, right = `u`, `d`, `l`, `r`
		obstacle              = `#`
		traversed             = `x`
		guard                 = `g`
		space                 = `.`
		startGuard            = `^`
	)

	var (
		guardX, guardY = 0, 0
		maxX           = len(instructions[0])
		maxY           = len(instructions)
		grid           = make([][]day06cell, maxY)
	)

	// 'Obstacle' X and Y.
	for oy := 0; oy < maxY; oy++ {
		for ox := 0; ox < maxX; ox++ {
			// Reset these each run.
			outOfBounds := false
			dir := `u`

			// Reset the grid.
			for i := 0; i < maxY; i++ {
				grid[i] = make([]day06cell, maxX)
			}

			for y := 0; y < maxY; y++ {
				for x := 0; x < maxX; x++ {
					grid[y][x].path = string(instructions[y][x])

					// Guard's start position.
					if grid[y][x].path == startGuard {
						guardX = x
						guardY = y

						grid[y][x].path = guard
					}
				}
			}

			// Insert the new obstacle or skip if the cell is not space.
			if grid[oy][ox].path == space {
				grid[oy][ox].path = obstacle
			} else {
				continue
			}

			for {
				switch dir {
				case up:
					if guardY-1 < 0 {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						outOfBounds = true
					} else if grid[guardY-1][guardX].path == obstacle {
						dir = right
					} else {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						guardY--
						grid[guardY][guardX].path = guard
					}

				case right:
					if guardX+1 >= maxX {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						outOfBounds = true
					} else if grid[guardY][guardX+1].path == obstacle {
						dir = down
					} else {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						guardX++
						grid[guardY][guardX].path = guard
					}

				case down:
					if guardY+1 >= maxY {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						outOfBounds = true
					} else if grid[guardY+1][guardX].path == obstacle {
						dir = left
					} else {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						guardY++
						grid[guardY][guardX].path = guard
					}

				case left:
					if guardX-1 < 0 {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						outOfBounds = true
					} else if grid[guardY][guardX-1].path == obstacle {
						dir = up
					} else {
						grid[guardY][guardX].path = traversed
						grid[guardY][guardX].count++
						guardX--
						grid[guardY][guardX].path = guard
					}
				}

				if outOfBounds {
					goto done
				}

				if grid[guardY][guardX].count > 3 {
					total++

					if cfg.debug {
						day06DrawGrid2(grid)
					}

					break
				}
			}

		done:
		}
	}

	return
}

func day06DrawGrid2(grid [][]day06cell) {
	summary := make(map[int]int)

	for _, gr := range grid {
		for _, gc := range gr {
			if gc.count > 2 {
				fmt.Printf("\033[0;31m%s\033[0m", gc.path)
			} else {
				fmt.Print(gc.path)
			}

			summary[gc.count]++
		}
		fmt.Println()
	}
	fmt.Println()
}
