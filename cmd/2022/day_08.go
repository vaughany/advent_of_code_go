package main

import (
	"fmt"
	"strconv"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day08(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day08part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day08part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-08-1: 1776
func (cfg *config) day08part1(instructions []string) int {
	var (
		out      int
		gridSize = len(instructions[0])
		grid     = make([][]int, gridSize)
	)

	// Create grid.
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	// Populate grid from instructions.
	for li, line := range instructions {
		for di, digit := range line {
			d, _ := strconv.Atoi(string(digit))
			grid[di][li] = d
		}
	}

	if cfg.debug {
		day08drawGrid(grid)
	}

	// Print the grid, but ignore the perimeter trees.
	if cfg.debug {
		for y := 1; y <= len(grid)-2; y++ {
			for x := 1; x <= len(grid[0])-2; x++ {
				fmt.Print(grid[x][y])
			}
			fmt.Println()
		}
	}

	for y := 1; y <= len(grid)-2; y++ {
		for x := 1; x <= len(grid[0])-2; x++ {
		restart:
			if x > len(grid[0])-2 || y > len(grid[0])-2 {
				break
			}

			if cfg.debug {
				fmt.Printf("x: %d, y: %d, id: %d\n", x, y, grid[x][y])
			}

			var visible bool

			// UP:
			visible = true
			for j := 1; j <= y; j++ {
				if cfg.debug {
					fmt.Println("UP:", grid[x][y-j])
				}
				if grid[x][y-j] >= grid[x][y] {
					visible = false
					break
				}
			}
			if visible {
				out++
				x++
				goto restart
			}

			// DOWN:
			visible = true
			for j := y; j <= len(grid)-2; j++ {
				if cfg.debug {
					fmt.Println("DOWN:", grid[x][j+1])
				}
				if grid[x][j+1] >= grid[x][y] {
					visible = false
					break
				}
			}
			if visible {
				out++
				x++
				goto restart
			}

			// LEFT:
			visible = true
			for j := 1; j <= x; j++ {
				if cfg.debug {
					fmt.Println("LEFT:", grid[x-j][y])
				}
				if grid[x-j][y] >= grid[x][y] {
					visible = false
					break
				}
			}
			if visible {
				out++
				x++
				goto restart
			}

			// RIGHT:
			visible = true
			for j := x; j <= len(grid)-2; j++ {
				if cfg.debug {
					fmt.Println("RIGHT:", grid[j+1][y])
				}
				if grid[j+1][y] >= grid[x][y] {
					visible = false
					break
				}
			}
			if visible {
				out++
				x++
				goto restart
			}

		}
	}

	return out + (gridSize * 4) - 4
	// return out
}

// 2022-08-2: 234416
func (cfg *config) day08part2(instructions []string) int {
	var (
		out      int
		gridSize = len(instructions[0])
		grid     = make([][]int, gridSize)
	)

	// Create grid.
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	// Populate grid from instructions.
	for li, line := range instructions {
		for di, digit := range line {
			d, _ := strconv.Atoi(string(digit))
			grid[di][li] = d
		}
	}

	if cfg.debug {
		day08drawGrid(grid)
	}

	for y := 1; y <= len(grid)-2; y++ {
		for x := 1; x <= len(grid[0])-2; x++ {
			if x > len(grid[0])-2 || y > len(grid[0])-2 {
				break
			}

			if cfg.debug {
				fmt.Printf("x: %d, y: %d, id: %d\n", x, y, grid[x][y])
			}

			// UP:
			var uc int
			for j := 1; j <= y; j++ {
				if grid[x][y-j] < grid[x][y] {
					uc++
				} else if grid[x][y-j] >= grid[x][y] {
					uc++
					break
				}
			}
			if uc == 0 {
				uc = 1
			}
			if cfg.debug {
				fmt.Println("UP:", uc)
			}

			// DOWN:
			var dc int
			for j := y; j <= len(grid)-2; j++ {
				if grid[x][j+1] < grid[x][y] {
					dc++
				} else if grid[x][j+1] >= grid[x][y] {
					dc++
					break
				}
			}
			if dc == 0 {
				dc = 1
			}
			if cfg.debug {
				fmt.Println("DOWN:", dc)
			}

			// LEFT:
			var lc int
			for j := 1; j <= x; j++ {
				if grid[x-j][y] < grid[x][y] {
					lc++
				} else if grid[x-j][y] >= grid[x][y] {
					lc++
					break
				}
			}
			if lc == 0 {
				lc = 1
			}
			if cfg.debug {
				fmt.Println("LEFT:", lc)
			}

			// RIGHT:
			var rc int
			for j := x; j <= len(grid)-2; j++ {
				if grid[j+1][y] < grid[x][y] {
					rc++
				} else if grid[j+1][y] >= grid[x][y] {
					rc++
					break
				}
			}
			if rc == 0 {
				rc = 1
			}
			if cfg.debug {
				fmt.Println("RIGHT:", rc)
			}

			tmpOut := uc * lc * rc * dc
			if tmpOut > out {
				out = tmpOut
			}
		}
	}

	return out
}

func day08drawGrid(grid [][]int) {
	for y := 0; y <= len(grid)-1; y++ {
		for x := 0; x <= len(grid[0])-1; x++ {
			fmt.Print(grid[x][y])
		}

		fmt.Println()
	}

	fmt.Println()
}
