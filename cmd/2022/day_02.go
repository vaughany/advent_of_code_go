package main

import (
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day02(loader loaders.Loader) error {
	// timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		return err
	}
	// output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))

	// timingPartOne := time.Now()
	output.AnswerPart1(cfg.day02part1(instructions))
	// output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))

	// timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day02part2(instructions))
	// output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))

	// output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
	// output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))

	return nil
}

// 2022-02-1: 11841
func (cfg *config) day02part1(instructions []string) int {
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
func (cfg *config) day02part2(instructions []string) int {
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
