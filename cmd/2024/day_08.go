package main

import (
	"fmt"
	"math"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day08(loader loaders.Loader) error {
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
	output.AnswerPart1(cfg.day08part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day08part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

type day08node struct {
	x, y     int
	freq     string
	antinode bool
}

type day08xy struct {
	x, y int
}

// 2024-08-1: 259
func (cfg *config) day08part1(instructions []string) (total int) {
	var (
		maxX     = len(instructions[0])
		maxY     = len(instructions)
		grid     = make([][]day08node, maxY)
		nodeList = make(map[string][]day08xy)
	)

	// Initialise grid.
	for y := 0; y < maxY; y++ {
		grid[y] = make([]day08node, maxX)
		for x := 0; x < maxX; x++ {
			grid[y][x] = day08node{x: x, y: y, freq: string(instructions[y][x])}

			if grid[y][x].freq != `.` {
				nodeList[grid[y][x].freq] = append(nodeList[grid[y][x].freq], day08xy{x: x, y: y})
			}
		}
	}

	for _, nodelist := range nodeList {
		for i := 0; i < len(nodelist); i++ {
			for j := i + 1; j < len(nodelist); j++ {
				diffX := int(math.Abs(float64(nodelist[i].x) - float64(nodelist[j].x)))
				diffY := int(math.Abs(float64(nodelist[i].y) - float64(nodelist[j].y)))

				day08ProcessNodes(grid, nodelist, i, j, diffX, diffY, maxX, maxY)
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x].antinode {
				total++
			}
		}
	}

	if cfg.debug {
		day08DrawGrid(grid)
	}

	return
}

// 2024-08-2: 927
func (cfg *config) day08part2(instructions []string) (total int) {
	var (
		maxX     = len(instructions[0])
		maxY     = len(instructions)
		grid     = make([][]day08node, maxY)
		nodeList = make(map[string][]day08xy)
	)

	// Initialise grid.
	for y := 0; y < maxY; y++ {
		grid[y] = make([]day08node, maxX)
		for x := 0; x < maxX; x++ {
			grid[y][x] = day08node{x: x, y: y, freq: string(instructions[y][x])}

			if grid[y][x].freq != `.` {
				nodeList[grid[y][x].freq] = append(nodeList[grid[y][x].freq], day08xy{x: x, y: y})
			}
		}
	}

	for _, nodelist := range nodeList {
		for i := 0; i < len(nodelist); i++ {
			for j := i + 1; j < len(nodelist); j++ {
				diffX := int(math.Abs(float64(nodelist[i].x) - float64(nodelist[j].x)))
				diffY := int(math.Abs(float64(nodelist[i].y) - float64(nodelist[j].y)))

				for z := 1; z < maxX-1; z++ {
					dX := diffX * z
					dY := diffY * z

					day08ProcessNodes(grid, nodelist, i, j, dX, dY, maxX, maxY)
				}
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x].antinode || grid[y][x].freq != `.` {
				total++
			}
		}
	}

	if cfg.debug {
		day08DrawGrid(grid)
	}

	return
}

func day08ProcessNodes(grid [][]day08node, nodelist []day08xy, i, j, diffX, diffY, maxX, maxY int) {
	var newX1, newX2, newY1, newY2 int

	if nodelist[i].x < nodelist[j].x {
		newX1 = nodelist[i].x - diffX
		newX2 = nodelist[j].x + diffX
	} else {
		newX1 = nodelist[i].x + diffX
		newX2 = nodelist[j].x - diffX
	}

	if nodelist[i].y < nodelist[j].y {
		newY1 = nodelist[i].y - diffY
		newY2 = nodelist[j].y + diffY
	} else {
		newY1 = nodelist[i].y + diffY
		newY2 = nodelist[j].y - diffY
	}

	// CBA to invert this lot.
	if newX1 > maxX-1 || newX1 < 0 || newY1 > maxY-1 || newY1 < 0 {
		//
	} else {
		grid[newY1][newX1].antinode = true
	}

	if newX2 > maxX-1 || newX2 < 0 || newY2 > maxY-1 || newY2 < 0 {
		//
	} else {
		grid[newY2][newX2].antinode = true
	}
}

func day08DrawGrid(in [][]day08node) {
	for _, row := range in {
		for _, n := range row {
			fmt.Print(n.freq)
		}
		fmt.Println()
	}

	fmt.Println()

	for _, row := range in {
		for _, n := range row {
			if n.antinode {
				fmt.Print(`#`)
			} else {
				fmt.Print(n.freq)
			}
		}
		fmt.Println()
	}

}
