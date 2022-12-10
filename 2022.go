package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (cfg *config) run2022() {
	if cfg.day == 1 || cfg.day == 0 {
		cfg.title(2022, 1)
		timingStart := time.Now()
		instructions := cfg.getInputAsInts(2022, 1)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2022day01part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2022day01part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 2 || cfg.day == 0 {
		cfg.title(2022, 2)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2022, 2)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2022day02part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2022day02part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 3 || cfg.day == 0 {
		cfg.title(2022, 3)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2022, 3)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2022day03part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2022day03part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 4 || cfg.day == 0 {
		cfg.title(2022, 4)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2022, 4)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		pt1, pt2 := cfg.year2022day04(instructions)
		cfg.answerPart1(pt1)
		cfg.answerPart2(pt2)

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 5 || cfg.day == 0 {
		cfg.title(2022, 5)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2022, 5)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2022day05part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2022day05part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 6 || cfg.day == 0 {
		cfg.title(2022, 6)
		timingStart := time.Now()
		instructions := cfg.getInputAsString(2022, 6)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2022day06part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2022day06part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 7 || cfg.day == 0 {
		cfg.title(2022, 7)
		cfg.notYetImplemented()
		// timingStart := time.Now()
		// instructions := cfg.getInputAsStrings(2022, 7)
		// cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		// timingPartOne := time.Now()
		// cfg.answerPart1(cfg.year2022day07part1(instructions))
		// cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		// timingPartTwo := time.Now()
		// cfg.answerPart2(cfg.year2022day07part2(instructions))
		// cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		// cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		// cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 8 || cfg.day == 0 {
		cfg.title(2022, 8)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2022, 8)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2022day08part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2022day08part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 9 || cfg.day == 0 {
		cfg.title(2015, 9)
		cfg.notYetImplemented()
	}

	if cfg.day == 10 || cfg.day == 0 {
		cfg.title(2015, 10)
		cfg.notYetImplemented()
	}

	if cfg.day == 11 || cfg.day == 0 {
		cfg.title(2015, 11)
		cfg.notYetImplemented()
	}

	if cfg.day == 12 || cfg.day == 0 {
		cfg.title(2015, 12)
		cfg.notYetImplemented()
	}

	if cfg.day == 13 || cfg.day == 0 {
		cfg.title(2015, 13)
		cfg.notYetImplemented()
	}

	if cfg.day == 14 || cfg.day == 0 {
		cfg.title(2015, 14)
		cfg.notYetImplemented()
	}

	if cfg.day == 15 || cfg.day == 0 {
		cfg.title(2015, 15)
		cfg.notYetImplemented()
	}

	if cfg.day == 16 || cfg.day == 0 {
		cfg.title(2015, 16)
		cfg.notYetImplemented()
	}

	if cfg.day == 17 || cfg.day == 0 {
		cfg.title(2015, 17)
		cfg.notYetImplemented()
	}

	if cfg.day == 18 || cfg.day == 0 {
		cfg.title(2015, 18)
		cfg.notYetImplemented()
	}

	if cfg.day == 19 || cfg.day == 0 {
		cfg.title(2015, 19)
		cfg.notYetImplemented()
	}

	if cfg.day == 20 || cfg.day == 0 {
		cfg.title(2015, 20)
		cfg.notYetImplemented()
	}

	if cfg.day == 21 || cfg.day == 0 {
		cfg.title(2015, 21)
		cfg.notYetImplemented()
	}

	if cfg.day == 22 || cfg.day == 0 {
		cfg.title(2015, 22)
		cfg.notYetImplemented()
	}

	if cfg.day == 23 || cfg.day == 0 {
		cfg.title(2015, 23)
		cfg.notYetImplemented()
	}

	if cfg.day == 24 || cfg.day == 0 {
		cfg.title(2015, 24)
		cfg.notYetImplemented()
	}

	if cfg.day == 25 || cfg.day == 0 {
		cfg.title(2015, 25)
		cfg.notYetImplemented()
	}

}

// 2022-01-1: 70369
func (cfg *config) year2022day01part1(instructions []int) int {
	var (
		elves        = make(map[int]int, len(instructions))
		elf          int
		calorificElf int
	)

	for _, ins := range instructions {
		elves[elf] += ins

		if ins == 0 {
			if elves[elf] > calorificElf {
				calorificElf = elves[elf]
			}
			elf++
		}
	}

	return calorificElf
}

// 2022-01-2: 203002
func (cfg *config) year2022day01part2(instructions []int) int {
	var (
		elves = make(map[int]int, len(instructions))
		elf   int
	)

	for _, ins := range instructions {
		elves[elf] += ins

		if ins == 0 {
			elf++
		}
	}

	keys := make([]int, 0, len(elves))
	for key := range elves {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return elves[keys[i]] > elves[keys[j]]
	})

	return elves[keys[0]] + elves[keys[1]] + elves[keys[2]]
}

// 2022-02-1: 11841
func (cfg *config) year2022day02part1(instructions []string) int {
	var totalScore int

	const (
		rock, paper, scissors          = "X", "Y", "Z"
		oppRock, oppPaper, oppScissors = "A", "B", "C"
	)

	scores := struct {
		rock, paper, scissors int
		win, lose, draw       int
	}{
		rock:     1,
		paper:    2,
		scissors: 3,
		win:      6,
		lose:     0,
		draw:     3,
	}

	for _, ins := range instructions {
		split := strings.Split(ins, " ")
		opponent, me := split[0], split[1]

		switch me {
		case rock:
			totalScore += scores.rock
		case paper:
			totalScore += scores.paper
		case scissors:
			totalScore += scores.scissors
		}

		switch opponent {
		case oppRock:
			switch me {
			case rock:
				totalScore += scores.draw
			case paper:
				totalScore += scores.win
			case scissors:
				totalScore += scores.lose
			}
		case oppPaper:
			switch me {
			case rock:
				totalScore += scores.lose
			case paper:
				totalScore += scores.draw
			case scissors:
				totalScore += scores.win
			}
		case oppScissors:
			switch me {
			case rock:
				totalScore += scores.win
			case paper:
				totalScore += scores.lose
			case scissors:
				totalScore += scores.draw
			}
		}
	}

	return totalScore
}

// 2022-02-2: 13022
func (cfg *config) year2022day02part2(instructions []string) int {
	var totalScore int

	const (
		lose, draw, win       = "X", "Y", "Z"
		rock, paper, scissors = "A", "B", "C"
	)

	scores := struct {
		rock, paper, scissors int
		win, lose, draw       int
	}{
		rock:     1,
		paper:    2,
		scissors: 3,
		win:      6,
		lose:     0,
		draw:     3,
	}

	for _, ins := range instructions {
		split := strings.Split(ins, " ")
		opponent, outcome := split[0], split[1]

		switch outcome {
		case lose:
			totalScore += scores.lose
		case draw:
			totalScore += scores.draw
		case win:
			totalScore += scores.win
		}

		switch opponent {
		case rock:
			switch outcome {
			case lose:
				totalScore += scores.scissors
			case draw:
				totalScore += scores.rock
			case win:
				totalScore += scores.paper
			}
		case paper:
			switch outcome {
			case lose:
				totalScore += scores.rock
			case draw:
				totalScore += scores.paper
			case win:
				totalScore += scores.scissors
			}
		case scissors:
			switch outcome {
			case lose:
				totalScore += scores.paper
			case draw:
				totalScore += scores.scissors
			case win:
				totalScore += scores.rock
			}
		}
	}

	return totalScore
}

// 2022-03-1: 8018
func (cfg *config) year2022day03part1(instructions []string) int {
	var score int

	for _, ins := range instructions {
		half1, half2 := ins[:len(ins)/2], ins[len(ins)/2:]

		for _, letter := range half1 {
			if strings.Contains(half2, string(letter)) {
				asciiOfRune := int(letter)
				if asciiOfRune >= 97 && asciiOfRune <= 122 {
					score += asciiOfRune - 96
				} else if asciiOfRune >= 65 && asciiOfRune <= 90 {
					score += asciiOfRune - 38
				}
				break
			}
		}
	}

	return score
}

// 2022-03-2: 2518
func (cfg *config) year2022day03part2(instructions []string) int {
	var (
		elfGroup   = make([]string, 3)
		elfInGroup int
		score      int
	)

	for _, ins := range instructions {
		elfGroup[elfInGroup] = ins

		elfInGroup++
		if elfInGroup == 3 {
			elfInGroup = 0

			for _, letter := range elfGroup[0] {
				if strings.Contains(elfGroup[1], string(letter)) {
					if strings.Contains(elfGroup[2], string(letter)) {
						asciiOfRune := int(letter)
						if asciiOfRune >= 97 && asciiOfRune <= 122 {
							score += asciiOfRune - 96
						} else if asciiOfRune >= 65 && asciiOfRune <= 90 {
							score += asciiOfRune - 38
						}

						break
					}
				}
			}
		}
	}

	return score
}

// 2022-04-1: 573
// 2022-04-2: 867
func (cfg *config) year2022day04(instructions []string) (int, int) {
	var (
		contained, overlapping int
		s1s, s1e, s2s, s2e     int
	)

	for _, ins := range instructions {
		sections := strings.Split(ins, ",")

		s1 := strings.Split(sections[0], "-")
		s2 := strings.Split(sections[1], "-")

		s1s, _ = strconv.Atoi(s1[0])
		s1e, _ = strconv.Atoi(s1[1])
		s2s, _ = strconv.Atoi(s2[0])
		s2e, _ = strconv.Atoi(s2[1])

		if cfg.debug {
			fmt.Println(s1s, s1e, s2s, s2e)
		}

		if s1s <= s2s && s1e >= s2e || s2s <= s1s && s2e >= s1e {
			contained++
		}

		if (s1s >= s2s && s1s <= s2e || s1e >= s2s && s1e <= s2e) || (s1s <= s2s && s1e >= s2e || s2s <= s1s && s2e >= s1e) {
			overlapping++
		}
	}

	return contained, overlapping
}

// 2022-05-1: CFFHVVHNC
func (cfg *config) year2022day05part1(instructions []string) string {
	const cutset = `[] 123456789`

	var (
		tmpStacks = map[int][]string{}
		stacks    = map[int][]string{}
		setup     = true
	)

	for _, ins := range instructions {
		// Parse the first bit of the input file.
		if setup {
			if ins == "" {
				setup = false

				for j := 1; j <= len(tmpStacks); j++ {
					for k := len(tmpStacks[j]) - 1; k >= 0; k-- {
						if tmpStacks[j][k] != "" {
							stacks[j] = append(stacks[j], tmpStacks[j][k])
						}
					}
				}

				continue
			}

			for j := 0; j < (len(ins)+1)/4; j++ {
				tmpStacks[j+1] = append(tmpStacks[j+1], strings.Trim(ins[j*4:j*4+3], cutset))
			}
		}

		var count, from, to int
		fmt.Sscanf(ins, "move %2d from %2d to %2d", &count, &from, &to)

		if count == 0 {
			continue
		}

		for j := 1; j <= count; j++ {
			// Get the package
			pack := stacks[from][len(stacks[from])-1:][0]

			// Remove it from the 'old' stack
			stacks[from] = stacks[from][:len(stacks[from])-1]

			// Add it to the 'new' stack
			stacks[to] = append(stacks[to], pack)
		}
	}

	// Get top-most package from each stack
	var out string
	for j := 1; j <= len(stacks); j++ {
		out += stacks[j][len(stacks[j])-1:][0]
	}

	return out
}

// 2022-05-2: FSZWBPTBG
func (cfg *config) year2022day05part2(instructions []string) string {
	const cutset = `[] 123456789`

	var (
		tmpStacks = map[int][]string{}
		stacks    = map[int][]string{}
		setup     = true
	)

	for _, ins := range instructions {
		// Parse the first bit of the input file.
		if setup {
			if ins == "" {
				setup = false

				for j := 1; j <= len(tmpStacks); j++ {
					for k := len(tmpStacks[j]) - 1; k >= 0; k-- {
						if tmpStacks[j][k] != "" {
							stacks[j] = append(stacks[j], tmpStacks[j][k])
						}
					}
				}

				continue
			}

			for j := 0; j < (len(ins)+1)/4; j++ {
				tmpStacks[j+1] = append(tmpStacks[j+1], strings.Trim(ins[j*4:j*4+3], cutset))
			}
		}

		var count, from, to int
		fmt.Sscanf(ins, "move %2d from %2d to %2d", &count, &from, &to)

		if count == 0 {
			continue
		}

		// Get the package/s
		pack := stacks[from][len(stacks[from])-count:]

		// Remove it/them from the 'old' stack
		stacks[from] = stacks[from][:len(stacks[from])-count]

		// Add it/them to the 'new' stack
		stacks[to] = append(stacks[to], pack...)
	}

	// Get top-most package from each stack
	var out string
	for j := 1; j <= len(stacks); j++ {
		out += stacks[j][len(stacks[j])-1:][0]
	}

	return out
}

// 2022-06-1: 1175
func (cfg *config) year2022day06part1(instructions string) int {
	var (
		window = 4
		out    int
	)

	for j := 0; j <= len(instructions)-window; j++ {
		slice := instructions[j : j+window]

		chars := make(map[byte]int, window)
		for i := 0; i < window; i++ {
			chars[slice[i]]++
		}

		if len(chars) == window {
			out = j + window

			break
		}
	}

	return out
}

// 2022-06-2: 3217
func (cfg *config) year2022day06part2(instructions string) int {
	var (
		window = 14
		out    int
	)

	for j := 0; j <= len(instructions)-window; j++ {
		slice := instructions[j : j+window]

		chars := make(map[byte]int, window)
		for i := 0; i < window; i++ {
			chars[slice[i]]++
		}

		if len(chars) == window {
			out = j + window

			break
		}
	}

	return out
}

// 2022-07-1: NOT 1226869, too log
// func (cfg *config) year2022day07part1(instructions []string) int {
// path := ""
// paths := make(map[string]int)
// currentFolder := ""

// for _, ins := range instructions {
// 	// fmt.Println(ins)
// 	if strings.HasPrefix(ins, "$") {
// 		command := strings.Split(ins, " ")[1:]
// 		if command[0] == "ls" {
// 			continue
// 		}

// 		if command[0] == "cd" {
// 			switch command[1] {
// 			case "/":
// 				fmt.Println("GO TO ROOT")
// 				path = "/"
// 			case "..":
// 				fmt.Println("BACK")
// 				path = strings.Replace(path, currentFolder+"/", "", 1)
// 			default:
// 				fmt.Println("CHANGE TO FOLDER")
// 				currentFolder = command[1]
// 				path += currentFolder + "/"
// 			}
// 		}

// 	} else {
// 		file := strings.Split(ins, " ")
// 		if file[0] == "dir" {
// 			fmt.Println("CREATE FOLDER")
// 			paths[path+file[1]+"/"] += 0
// 		} else {
// 			fmt.Println("FILE SIZE")
// 			size, _ := strconv.Atoi(file[0])
// 			paths[path+file[1]] = size
// 		}

// 	}

// 	fmt.Printf("%#v\n", path)
// 	fmt.Printf("%#v\n\n", paths)
// }

// return 0
// }

// 2022-07-2:
// func (cfg *config) year2022day07part2(instructions []string) int {
// 	return 0
// }

// 2022-08-1: 1776
func (cfg *config) year2022day08part1(instructions []string) int {
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
		cfg.drawGrid(grid)
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

func (cfg *config) drawGrid(grid [][]int) {
	for y := 0; y <= len(grid)-1; y++ {
		for x := 0; x <= len(grid[0])-1; x++ {
			fmt.Print(grid[x][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

// 2022-08-2: 234416
func (cfg *config) year2022day08part2(instructions []string) int {
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
		cfg.drawGrid(grid)
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
