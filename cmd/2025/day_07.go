package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day07() error {
	return runDayWithInput(
		cfg,
		func(l loaders.Loader) ([]string, error) { return l.GetStrings() },
		cfg.day07part1,
		cfg.day07part2,
	)
}

type cell struct {
	splitter bool
	beam     bool
	ways     int
}

// 2025-07-1: 1640
func (cfg *config) day07part1(instructions []string) int {
	height := len(instructions)
	width := len(instructions[0])

	grid := make([][]cell, height)
	for r := range height {
		grid[r] = make([]cell, width)
	}

	for r, line := range instructions {
		for c := range width {
			switch line[c] {
			case '^':
				grid[r][c].splitter = true
			case 'S':
				grid[r][c].beam = true
			}
		}
	}

	if cfg.debug {
		drawGrid(grid, true, false)
	}

	splittersHit := 0

	for r := 1; r < height; r++ {
		for c := range width {
			// See if there's a beam above.
			aboveBeam := grid[r-1][c].beam

			// If this is a splitter *with a beam above* - important qualification.
			if grid[r][c].splitter && aboveBeam {
				splittersHit++

				// Make new beams either side of the splitter.
				grid[r][c-1].beam = true
				grid[r][c+1].beam = true

				continue
			}

			// Propagate beam.
			if aboveBeam {
				grid[r][c].beam = true
			}
		}

		if cfg.debug {
			drawGrid(grid, true, false)
		}
	}

	return splittersHit
}

// 2025-07-2: 40999072541589
func (cfg *config) day07part2(instructions []string) int {
	height := len(instructions)
	width := len(instructions[0])
	timelines := 1 // Afer all, we're already in a timeline.

	grid := make([][]cell, height)
	for r := range height {
		grid[r] = make([]cell, width)
	}

	for r, ins := range instructions {
		for c := range width {
			switch ins[c] {
			case '^':
				grid[r][c].splitter = true
			case 'S':
				grid[r][c].beam = true
				grid[r][c].ways = 1
			}
		}
	}

	if cfg.debug {
		drawGrid(grid, true, false)
	}

	for r := 1; r < height; r++ {
		for c := range width {
			// How many timelines are coming from above?
			aboveWays := grid[r-1][c].ways
			if aboveWays == 0 {
				continue
			}

			// If this is a splitter hit by some timelines...
			if grid[r][c].splitter {
				// Each incoming timeline splits into one extra timeline.
				timelines += aboveWays

				// Send those timelines to the cells either side.
				grid[r][c-1].ways += aboveWays
				grid[r][c-1].beam = true
				grid[r][c+1].ways += aboveWays
				grid[r][c+1].beam = true

				continue
			}

			// Otherwise, they just continue straight down.
			grid[r][c].ways += aboveWays
			grid[r][c].beam = true
		}

		if cfg.debug {
			drawGrid(grid, true, false)
		}
	}

	return timelines

}

func drawGrid(grid [][]cell, pause bool, nums bool) {
	var b strings.Builder

	for _, row := range grid {
		for _, col := range row {
			if nums && col.ways > 0 {
				b.WriteString(strconv.Itoa(col.ways))
			} else if col.beam {
				b.WriteByte('|')
			} else if col.splitter {
				b.WriteByte('^')
			} else {
				b.WriteByte('.')
			}
		}

		b.WriteByte('\n')
	}

	b.WriteByte('\n')

	fmt.Print(b.String())

	if pause {
		time.Sleep(500 * time.Millisecond)
	}
}
