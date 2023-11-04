package support

import (
	"vaughany.com/advent_of_code_go/internal/output"
)

const (
	filled = `@`
	empty  = `.`
)

func PrintGrid(grid [][]bool) {
	for i := range grid {
		var line string

		for j := range grid {
			if grid[i][j] {
				line += filled
			} else {
				line += empty
			}
		}

		output.Info(line)
	}
}

func SetGrid(grid [][]bool, x1, y1, x2, y2 int, on bool) [][]bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = on
		}
	}

	return grid
}

func ToggleGrid(grid [][]bool, x1, y1, x2, y2 int) [][]bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = !grid[i][j]
		}
	}

	return grid
}

func CountGridTrue(grid [][]bool) int {
	var out int

	for _, i := range grid {
		for _, j := range i {
			if j {
				out++
			}
		}
	}

	return out
}

func ChangeBrightness(grid [][]int, x1, y1, x2, y2 int, change bool) [][]int {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if change {
				grid[i][j]++
			} else {
				if grid[i][j] > 0 {
					grid[i][j]--
				}
			}
		}
	}

	return grid
}

func DoubleBrightness(grid [][]int, x1, y1, x2, y2 int) [][]int {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] += 2
		}
	}

	return grid
}

func CountBrightness(grid [][]int) int {
	var out int

	for _, i := range grid {
		for _, j := range i {
			out += j
		}
	}

	return out
}
