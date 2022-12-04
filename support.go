package main

import "fmt"

type area struct {
	x, y, z, xy, yz, zx, area, volume, smallestSide, smallestPerimeter int
}

func (s *area) getAreaOfBox() {
	s.xy = s.x * s.y
	s.yz = s.y * s.z
	s.zx = s.z * s.x

	s.area = (2 * s.xy) + (2 * s.yz) + (2 * s.zx)
}

func (s *area) getVolumeOfBox() {
	s.volume = s.x * s.y * s.z
}

func (s *area) getSmallestSide() {
	var n [3]int
	n[0] = s.xy
	n[1] = s.yz
	n[2] = s.zx

	smallest := n[0]
	for _, num := range n[1:] {
		if num < smallest {
			smallest = num
		}
	}

	s.smallestSide = smallest
}

func (s *area) getSmallestPerimeter() {
	var n [3]int
	n[0] = (s.x * 2) + (s.y * 2)
	n[1] = (s.y * 2) + (s.z * 2)
	n[2] = (s.z * 2) + (s.x * 2)

	smallest := n[0]
	for _, num := range n[1:] {
		if num < smallest {
			smallest = num
		}
	}

	s.smallestPerimeter = smallest
}

//

func (cfg *config) printGrid(grid [][]bool) {
	for i := range grid {
		var line string

		for j := range grid {
			if grid[i][j] {
				line += "@"
			} else {
				line += "."
			}
		}

		cfg.info(fmt.Sprint(line))
	}
}

func (cfg *config) setGrid(grid [][]bool, x1, y1, x2, y2 int, on bool) [][]bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = on
		}
	}

	return grid
}

func (cfg *config) toggleGrid(grid [][]bool, x1, y1, x2, y2 int) [][]bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = !grid[i][j]
		}
	}

	return grid
}

func (cfg *config) countGridTrue(grid [][]bool) int {
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

func (cfg *config) changeBrightness(grid [][]int, x1, y1, x2, y2 int, change bool) [][]int {
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

func (cfg *config) doubleBrightness(grid [][]int, x1, y1, x2, y2 int) [][]int {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] += 2
		}
	}

	return grid
}

func (cfg *config) countBrightness(grid [][]int) int {
	var out int

	for _, i := range grid {
		for _, j := range i {
			out += j
		}
	}

	return out
}
