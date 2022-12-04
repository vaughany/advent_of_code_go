package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (cfg *config) run2015() {
	if cfg.day == 1 || cfg.day == 0 {
		cfg.title(2015, 1)
		timingStart := time.Now()
		instructions := cfg.getInputAsString(2015, 1)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		pt1, pt2 := cfg.year2015day01(instructions)
		cfg.answerPart1(pt1)
		cfg.answerPart2(pt2)

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 2 || cfg.day == 0 {
		cfg.title(2015, 2)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2015, 2)

		newInstructions := []area{}
		for _, line := range instructions {
			s := strings.Split(line, "x")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])
			z, _ := strconv.Atoi(s[2])

			newInstructions = append(newInstructions, area{x: x, y: y, z: z})
		}
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2015day02part1(newInstructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2015day02part2(newInstructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 3 || cfg.day == 0 {
		cfg.title(2015, 3)
		timingStart := time.Now()
		instructions := cfg.getInputAsBytes(2015, 3)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2015day03part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2015day03part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 4 || cfg.day == 0 {
		cfg.title(2015, 4)
		timingStart := time.Now()
		instructions := cfg.getInputAsString(2015, 4)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2015day04part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2015day04part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 5 || cfg.day == 0 {
		cfg.title(2015, 5)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2015, 5)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2015day05part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2015day05part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 6 || cfg.day == 0 {
		cfg.title(2015, 6)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2015, 6)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2015day06part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2015day06part2(instructions))
		cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 7 || cfg.day == 0 {
		cfg.title(2015, 7)
		cfg.notYetImplemented()
		// timingStart := time.Now()
		// instructions := cfg.getInputAsStrings(2015, 7)
		// cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		// timingPartOne := time.Now()
		// cfg.answerPart1(cfg.year2015day07part1(instructions))
		// cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		// timingPartTwo := time.Now()
		// cfg.answerPart2(cfg.year2015day07part2(instructions))
		// cfg.timeInfo(InfoTypeTwo, time.Since(timingPartTwo))

		// cfg.timeInfo(InfoTypeBoth, time.Since(timingPartOne))
		// cfg.timeInfo(InfoTypeEverything, time.Since(timingStart))
	}

	if cfg.day == 8 || cfg.day == 0 {
		cfg.title(2015, 8)
		timingStart := time.Now()
		instructions := cfg.getInputAsStrings(2015, 8)
		cfg.timeInfo(InfoTypeSetup, time.Since(timingStart))

		timingPartOne := time.Now()
		cfg.answerPart1(cfg.year2015day08part1(instructions))
		cfg.timeInfo(InfoTypeOne, time.Since(timingPartOne))

		timingPartTwo := time.Now()
		cfg.answerPart2(cfg.year2015day08part2(instructions))
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

//

// 2015-01-1:
// 2015-01-2:
func (cfg *config) year2015day01(instructions string) (int, int) {
	var floor, firstBasement int

	for index, instruction := range instructions {
		switch instruction {
		case 40:
			floor++
		case 41:
			floor--
		}

		if firstBasement == 0 && floor < 0 {
			firstBasement = index + 1
		}
	}

	return floor, firstBasement
}

//

// 2015-02-1:
func (cfg *config) year2015day02part1(instructions []area) int {
	var out int

	for _, i := range instructions {
		i.getAreaOfBox()
		i.getSmallestSide()

		out += i.area + i.smallestSide
	}

	return out
}

// 2015-02-2:
func (cfg *config) year2015day02part2(instructions []area) int {
	var out int

	for _, i := range instructions {
		i.getVolumeOfBox()
		i.getSmallestPerimeter()
		out += i.volume + i.smallestPerimeter
	}

	return out
}

//

// 2015-03-1:
func (cfg *config) year2015day03part1(instructions []byte) int {
	var (
		x, y int
		grid = map[string]int{}
	)

	for _, j := range instructions {
		switch j {
		case 94: // Up.
			y++
		case 118: // Down.
			y--
		case 60: // Left.
			x--
		case 62: // Right.
			x++
		}

		grid[fmt.Sprintf("%d,%d", x, y)]++

		if cfg.debug {
			cfg.info(fmt.Sprintf("X: %d, Y: %d.", x, y))
		}
	}

	if cfg.debug {
		cfg.info(fmt.Sprintf("%#v", grid))
	}

	return len(grid)
}

// 2015-03-2:
func (cfg *config) year2015day03part2(instructions []byte) int {
	var (
		x1, y1, x2, y2 int
		grid           = map[string]int{}
	)

	for i, j := range instructions {
		if i%2 == 0 {
			switch j {
			case 94: // Up.
				y1++
			case 118: // Down.
				y1--
			case 60: // Left.
				x1--
			case 62: // Right.
				x1++
			}

			grid[fmt.Sprintf("%d,%d", x1, y1)]++

		} else {
			switch j {
			case 94: // Up.
				y2++
			case 118: // Down.
				y2--
			case 60: // Left.
				x2--
			case 62: // Right.
				x2++
			}

			grid[fmt.Sprintf("%d,%d", x2, y2)]++
		}

		if cfg.debug {
			cfg.info(fmt.Sprintf("X1: %d, Y1: %d.", x1, y1))
		}
	}

	if cfg.debug {
		cfg.info(fmt.Sprintf("%#v", grid))
	}

	return len(grid)
}

//

// 2015-04-1:
func (cfg *config) year2015day04part1(instructions string) int {
	var loop int

	for {
		loopBytes := []byte(fmt.Sprintf("%s%d", instructions, loop))
		md5 := fmt.Sprintf("%x", md5.Sum(loopBytes))

		if cfg.debug {
			cfg.info(fmt.Sprint(md5))
		}

		if strings.HasPrefix(md5, "00000") {
			return loop
		}

		loop++
	}
}

// 2015-04-2:
func (cfg *config) year2015day04part2(instructions string) int {
	var loop int

	for {
		loopBytes := []byte(fmt.Sprintf("%s%d", instructions, loop))
		md5 := fmt.Sprintf("%x", md5.Sum(loopBytes))

		if cfg.debug {
			cfg.info(fmt.Sprint(md5))
		}

		if strings.HasPrefix(md5, "000000") {
			return loop
		}

		loop++
	}
}

//

// 2015-05-1:
func (cfg *config) year2015day05part1(instructions []string) int {
	var (
		out      int
		banned   = []string{"ab", "cd", "pq", "xy"}
		vowels   = []string{"a", "e", "i", "o", "u"}
		alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		doubles  = []string{}
	)

	// Create the 'doubles' slice.
	for _, a1 := range alphabet {
		doubles = append(doubles, fmt.Sprint(a1, a1))
	}

instructions:
	for index, i := range instructions {

		// Filter out the excluded strings first.
		for _, sub := range banned {
			if strings.Contains(i, sub) {
				if cfg.debug {
					cfg.info(fmt.Sprintf("%d / %s contains '%s'", index, i, sub))
				}

				continue instructions
			}
		}

		// At least three vowels.
		var vowelCount int
		for _, vowel := range vowels {
			vowelCount += strings.Count(i, vowel)
		}

		if vowelCount < 3 {
			if cfg.debug {
				cfg.info(fmt.Sprintf("%d / %s contains only %d vowel/s", index, i, vowelCount))
			}

			continue instructions
		}

		// Double letters.
		var found bool
		for _, sub := range doubles {
			if strings.Contains(i, sub) {
				found = true
				break
			}
		}

		if !found {
			if cfg.debug {
				cfg.info(fmt.Sprintf("%d / %s does not contain any double letters", index, i))
			}

			continue instructions
		}

		out++
	}

	return out
}

// 2015-05-2:
func (cfg *config) year2015day05part2(instructions []string) int {
	var (
		out      int
		alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		repeats  = []string{}
		pairs    = []string{}
	)

	// Create 'repeats'.
	for _, a1 := range alphabet {
		for _, a2 := range alphabet {
			repeats = append(repeats, fmt.Sprint(a1, a2, a1))
			pairs = append(pairs, fmt.Sprint(a1, a2))
		}
	}

instructions:
	for index, i := range instructions {

		// "xyx" / "xxx" repeating pattern.
		var found bool
		for _, sub := range repeats {
			if strings.Contains(i, sub) {
				found = true
				break
			}
		}

		if !found {
			if cfg.debug {
				cfg.info(fmt.Sprintf("%d / %s does not contain any three-letter repeats", index, i))
			}

			continue instructions
		}

		// Pairs of letters, appearing at least twice.
		var found2 bool
		for _, pair := range pairs {
			if strings.Count(i, pair) >= 2 {
				found2 = true
				break
			}
		}

		if !found2 {
			if cfg.debug {
				cfg.info(fmt.Sprintf("%d / %s does not contain any repeated pairs", index, i))
			}

			continue instructions
		}

		if cfg.debug {
			cfg.info(fmt.Sprintf("%d / %s is good", index, i))
		}

		out++
	}

	return out
}

//

// 2015-06-1:
func (cfg *config) year2015day06part1(instructions []string) int {
	var (
		gridSize = 1000
		grid     = make([][]bool, gridSize)
	)

	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}

	if cfg.debug {
		cfg.printGrid(grid)
	}

	for _, ins := range instructions {
		var x1, y1, x2, y2 int

		switch {
		case strings.HasPrefix(ins, "turn on"):
			fmt.Sscanf(ins, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			cfg.setGrid(grid, x1, y1, x2, y2, true)

		case strings.HasPrefix(ins, "turn off"):
			fmt.Sscanf(ins, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			cfg.setGrid(grid, x1, y1, x2, y2, false)

		case strings.HasPrefix(ins, "toggle"):
			fmt.Sscanf(ins, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			cfg.toggleGrid(grid, x1, y1, x2, y2)
		}
	}

	return cfg.countGridTrue(grid)
}

// 2015-06-2:
func (cfg *config) year2015day06part2(instructions []string) int {
	var (
		gridSize = 1000
		grid     = make([][]int, gridSize)
	)

	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	for _, ins := range instructions {
		var x1, y1, x2, y2 int

		switch {
		case strings.HasPrefix(ins, "turn on"):
			fmt.Sscanf(ins, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			cfg.changeBrightness(grid, x1, y1, x2, y2, true)

		case strings.HasPrefix(ins, "turn off"):
			fmt.Sscanf(ins, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			cfg.changeBrightness(grid, x1, y1, x2, y2, false)

		case strings.HasPrefix(ins, "toggle"):
			fmt.Sscanf(ins, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			cfg.doubleBrightness(grid, x1, y1, x2, y2)
		}
	}

	return cfg.countBrightness(grid)
}

//

// // 2015-07-1:
// func (cfg *config) year2015day07part1(instructions []string) int {
// 	wires := map[string]int{}

// 	for {
// 		if wires["a"] != 0 {
// 			break
// 		}

// 		for _, ins := range instructions {
// 			// if len(ins) == 0 {
// 			// 	continue
// 			// }

// 			var (
// 				first, second string
// 				target        string
// 				amount        int
// 			)

// 			switch {
// 			case strings.Contains(ins, "AND"):
// 				fmt.Sscanf(ins, "%s AND %s -> %s", &first, &second, &target)
// 				wires[target] = wires[first] & wires[second]
// 				// fmt.Println("AND", first, wires[first], second, wires[second], target, wires[target])

// 			case strings.Contains(ins, "OR"):
// 				fmt.Sscanf(ins, "%s OR %s -> %s", &first, &second, &target)
// 				wires[target] = wires[first] | wires[second]
// 				// fmt.Println("OR", first, wires[first], second, wires[second], target, wires[target])

// 			case strings.Contains(ins, "LSHIFT"):
// 				fmt.Sscanf(ins, "%s LSHIFT %d -> %s", &first, &amount, &target)
// 				wires[target] = wires[first] << amount
// 				// fmt.Println("LSHIFT", amount, target, wires[target])

// 			case strings.Contains(ins, "RSHIFT"):
// 				fmt.Sscanf(ins, "%s RSHIFT %d -> %s", &first, &amount, &target)
// 				wires[target] = wires[first] >> amount
// 				// fmt.Println("RSHIFT", amount, target, wires[target])

// 			case strings.HasPrefix(ins, "NOT"):
// 				fmt.Sscanf(ins, "NOT %s -> %s", &first, &target)
// 				wires[target] = int(^uint16(wires[first]))
// 				// fmt.Println("NOT", first, wires[first], target, wires[target])

// 			default:
// 				fmt.Sscanf(ins, "%d -> %s", &amount, &target)
// 				wires[target] = amount
// 				// fmt.Println("DEFAULT", amount, target)

// 				// remove this instruction once processed.
// 				// instructions = append(instructions[:index], instructions[index+1:]...)
// 			}
// 		}

// 		fmt.Printf("%#v\n", wires)
// 		// fmt.Println(len(instructions))
// 		time.Sleep(500 * time.Millisecond)
// 	}

// 	fmt.Println(">>>>> A:", wires["a"])

// 	return wires["a"]
// }

// // 2015-07-2:
// func (cfg *config) year2015day07part2(instructions []string) int {
// 	var out int

// 	return out
// }

// 2015-08-1: 1342
func (cfg *config) year2015day08part1(instructions []string) int {
	var countCode, countString int

	r := regexp.MustCompile(`\\x[0-9a-f]{2}`)

	for _, ins := range instructions {
		countCode += len(ins)

		// Remove the start and end quotes.
		newIns := ins[1 : len(ins)-1]

		newIns = strings.ReplaceAll(newIns, `\"`, `Q`)
		newIns = strings.ReplaceAll(newIns, `\\`, `S`)
		newIns = r.ReplaceAllString(newIns, `R`)

		countString += len(newIns)

		if cfg.debug {
			fmt.Println(ins, "---", newIns, "|", len(ins), "---", len(newIns))
		}
	}

	return countCode - countString
}

// 2015-08-2: 2074
func (cfg *config) year2015day08part2(instructions []string) int {
	var countCode, countString int

	for _, ins := range instructions {
		countCode += len(ins)

		newIns := strings.ReplaceAll(ins, `\`, `\\`)
		newIns = strings.ReplaceAll(newIns, `"`, `\"`)
		newIns = `"` + newIns + `"`

		countString += len(newIns)

		if cfg.debug {
			fmt.Println(ins, "---", newIns, "|", len(ins), "---", len(newIns))
		}
	}

	return countString - countCode
}
