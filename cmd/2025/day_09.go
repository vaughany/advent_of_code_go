package main

import (
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day09() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day09part1,
		cfg.day09part2,
	)
}

type coord struct {
	x, y int
}

type segment struct {
	a, b coord
}

// 2025-09-1: 4777409595
func (cfg *config) day09part1(instructions []string) int {
	maxX := 0
	maxY := 0
	coordinates := make([]coord, 0, len(instructions))

	for _, ins := range instructions {
		parts := strings.Split(ins, ",")
		x, _ := strconv.Atoi(parts[0])
		if x > maxX {
			maxX = x
		}

		y, _ := strconv.Atoi(parts[1])
		if y > maxY {
			maxY = y
		}

		coordinates = append(coordinates, coord{x: x, y: y})
	}

	lenCoordinates := len(coordinates)

	// grid := make([][]colour, maxY+1)
	// for r := 0; r <= maxY; r++ {
	// 	grid[r] = make([]colour, maxX+1)
	// }

	// for _, c := range coordinates {
	// 	// fmt.Println(c.x, c.y)
	// 	grid[c.y][c.x].red = true
	// }

	// drawGridDay9(grid, false)

	maxArea := 0
	for i := range lenCoordinates {
		for j := i + 1; j < lenCoordinates; j++ {
			a := rectAreaTiles(coordinates[i], coordinates[j])
			if a > maxArea {
				maxArea = a
			}
		}
	}

	return maxArea
}

// 2025-09-2: 1473551379
func (cfg *config) day09part2(instructions []string) int {
	// Parse coordinates (same as part 1, but we don't need maxX/maxY here).
	coordinates := make([]coord, 0, len(instructions))

	for _, ins := range instructions {
		parts := strings.Split(ins, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		coordinates = append(coordinates, coord{x: x, y: y})
	}

	n := len(coordinates)

	// Build the polygon segments from consecutive red tiles, wrapping around.
	segs := make([]segment, 0, n)
	for i := range n {
		a := coordinates[i]
		b := coordinates[(i+1)%n] // wrap last->first
		segs = append(segs, segment{a: a, b: b})
	}

	maxArea := 0

	// Try all pairs of red tiles as opposite corners.
	for i := range n {
		for j := i + 1; j < n; j++ {
			a := coordinates[i]
			b := coordinates[j]

			// Skip degenerate rectangles (zero width or height).
			if a.x == b.x || a.y == b.y {
				continue
			}

			// If the polygon boundary crosses the interior, this rectangle uses some non-green tiles and is invalid.
			if intersectsRect(a, b, segs) {
				continue
			}

			area := rectAreaTiles(a, b)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func rectAreaTiles(a, b coord) int {
	w := abs(a.x-b.x) + 1
	h := abs(a.y-b.y) + 1

	return w * h
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// intersectsRect returns true if any segment crosses the *interior* of the axis-aligned rectangle with opposite corners p and q.
func intersectsRect(p, q coord, segments []segment) bool {
	rx1, ry1 := p.x, p.y
	rx2, ry2 := q.x, q.y

	minX, maxX := rx1, rx2
	if minX > maxX {
		minX, maxX = maxX, minX
	}

	minY, maxY := ry1, ry2
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	// If rectangle is degenerate (zero width or height), it can't have area and we don't care about it anyway; but we still handle it safely.
	if minX == maxX || minY == maxY {
		return false
	}

	for _, s := range segments {
		sx1, sy1 := s.a.x, s.a.y
		sx2, sy2 := s.b.x, s.b.y

		if sx1 == sx2 {
			// Vertical segment at x = sx1 from sy1 to sy2.
			x := sx1
			yMin, yMax := sy1, sy2
			if yMin > yMax {
				yMin, yMax = yMax, yMin
			}

			// Check: x is strictly inside the rectangle's x-range, and the y-ranges overlap in the interior.
			if minX < x && x < maxX && max(yMin, minY) < min(yMax, maxY) {
				return true
			}

		} else if sy1 == sy2 {
			// Horizontal segment at y = sy1 from sx1 to sx2.
			y := sy1
			xMin, xMax := sx1, sx2
			if xMin > xMax {
				xMin, xMax = xMax, xMin
			}

			// Check: y is strictly inside the rectangle's y-range, and the x-ranges overlap in the interior.
			if minY < y && y < maxY && max(xMin, minX) < min(xMax, maxX) {
				return true
			}
		}
	}

	return false
}

// type colour struct {
// 	red   bool
// 	green bool
// 	rect  bool // Part of the chosen rectangle.
// }

// func drawGridDay9(grid [][]colour, pause bool) {
// 	var b strings.Builder

// 	for _, row := range grid {
// 		for _, col := range row {
// 			switch {
// 			case col.rect:
// 				// The chosen rectangle area
// 				b.WriteByte('O')
// 			case col.red:
// 				// Red tiles (given coordinates)
// 				b.WriteByte('#')
// 			case col.green:
// 				// Green tiles (border + interior)
// 				b.WriteByte('X')
// 			default:
// 				// Outside / unused
// 				b.WriteByte('.')
// 			}
// 		}
// 		b.WriteByte('\n')
// 	}

// 	b.WriteByte('\n')

// 	fmt.Print(b.String())

// 	if pause {
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
