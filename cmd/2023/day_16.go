package main

import (
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day16(loader loaders.Loader) error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStringsGrid(loader)
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day16part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day16part2(instructions))

	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

// 2023-16-1: 7939
func (cfg *config) day16part1(instructions [][]string) (result int) {
	width := len(instructions[0])
	height := len(instructions)

	type position struct {
		x    int
		y    int
		dir  direction
		dead bool
	}

	// ~50 is fine for sample data, ~800 needed for main data.
	numLoops := 800
	if cfg.sample {
		numLoops = 50
	}

	positions := []position{
		{x: 0, y: 0, dir: right},
	}

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	for l := 0; l <= numLoops; l++ {
		for p, pos := range positions {
			if pos.dead == true {
				continue
			}

			visited[pos.y][pos.x] = true

			// fmt.Println(pos)

			switch instructions[pos.y][pos.x] {
			case "|":
				// If direction is left or right, split up and down.
				if pos.dir == left || pos.dir == right {
					if pos.y-1 >= 0 {
						positions = append(positions, position{x: pos.x, y: pos.y - 1, dir: up})
					}
					if pos.y+1 <= height-1 {
						positions = append(positions, position{x: pos.x, y: pos.y + 1, dir: down})
					}

					positions[p].dead = true

				} else if pos.dir == up {
					if pos.y-1 >= 0 {
						positions[p].y--
					} else {
						positions[p].dead = true
					}

				} else if pos.dir == down {
					if pos.y+1 <= height-1 {
						positions[p].y++
					} else {
						positions[p].dead = true
					}
				}

			case "-":
				// If direction is up or down, split left and right.
				if pos.dir == up || pos.dir == down {
					if pos.x-1 >= 0 {
						positions = append(positions, position{x: pos.x - 1, y: pos.y, dir: left})
					}
					if pos.x+1 <= width-1 {
						positions = append(positions, position{x: pos.x + 1, y: pos.y, dir: right})
					}

					positions[p].dead = true

				} else if pos.dir == left {
					if pos.x-1 >= 0 {
						positions[p].x--
					} else {
						positions[p].dead = true
					}

				} else if pos.dir == right {
					if pos.x+1 <= width-1 {
						positions[p].x++
					} else {
						positions[p].dead = true
					}
				}

			case "/":
				switch pos.dir {
				case up:
					if pos.x+1 <= width-1 {
						positions[p].x++
						positions[p].dir = right
					} else {
						positions[p].dead = true
					}

				case down:
					if pos.x-1 >= 0 {
						positions[p].x--
						positions[p].dir = left
					} else {
						positions[p].dead = true
					}

				case left:
					if pos.y+1 <= height-1 {
						positions[p].y++
						positions[p].dir = down
					} else {
						positions[p].dead = true
					}

				case right:
					if pos.y-1 >= 0 {
						positions[p].y--
						positions[p].dir = up
					} else {
						positions[p].dead = true
					}
				}

			case `\`:
				switch pos.dir {
				case up:
					if pos.x-1 >= 0 {
						positions[p].x--
						positions[p].dir = left
					} else {
						positions[p].dead = true
					}

				case down:
					if pos.x+1 <= width-1 {
						positions[p].x++
						positions[p].dir = right
					} else {
						positions[p].dead = true
					}

				case left:
					if pos.y-1 >= 0 {
						positions[p].y--
						positions[p].dir = up
					} else {
						positions[p].dead = true
					}

				case right:
					if pos.y+1 <= height-1 {
						positions[p].y++
						positions[p].dir = down
					} else {
						positions[p].dead = true
					}
				}

			default:
				// Overwrite this node as 'visited'.
				instructions[pos.y][pos.x] = "*"

				switch pos.dir {
				case up:
					if pos.y-1 >= 0 {
						positions[p].y--
					} else {
						positions[p].dead = true
					}
				case down:
					if pos.y+1 <= height-1 {
						positions[p].y++
					} else {
						positions[p].dead = true
					}
				case left:
					if pos.x-1 >= 0 {
						positions[p].x--
					} else {
						positions[p].dead = true
					}
				case right:
					if pos.x+1 <= width-1 {
						positions[p].x++
					} else {
						positions[p].dead = true
					}
				}

			}

		}

		// displayGrid(instructions)
		// fmt.Println()

		var visitCount int
		for yi, y := range visited {
			for xi := range y {
				if visited[yi][xi] {
					visitCount++
				}
			}
		}

		result = visitCount

		// fmt.Println("loops:", l, "visit count:", visitCount)
		// fmt.Println()

		// time.Sleep(time.Second)
	}

	return
}

// 2023-16-2:
func (cfg *config) day16part2(instructions [][]string) (result int) {
	return
}
