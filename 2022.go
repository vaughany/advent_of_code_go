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
