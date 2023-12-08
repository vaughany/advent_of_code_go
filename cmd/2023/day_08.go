package main

import (
	"regexp"
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

		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2023-08-1: 23147
func (cfg *config) day08part1(instructions []string) (steps int) {
	type node struct {
		left, right string
	}

	var (
		nodes      = make(map[string]node, len(instructions[2:]))
		re         = regexp.MustCompile(`\w{3}`)
		directions = instructions[0]
		zzzFound   bool
	)

	for _, instruction := range instructions[2:] {
		matches := re.FindAllString(instruction, 3)

		nodes[matches[0]] = node{left: matches[1], right: matches[2]}
	}

	workingNode := nodes["AAA"]
	for {
		for _, dir := range directions {
			var nodeID string

			switch string(dir) {
			case "L":
				nodeID = workingNode.left
			case "R":
				nodeID = workingNode.right
			}

			workingNode = nodes[nodeID]
			steps++

			if nodeID == "ZZZ" {
				zzzFound = true

				break
			}
		}

		if zzzFound {
			break
		}
	}

	return steps
}

// 2023-08-2:
func (cfg *config) day08part2(instructions []string) int {
	return 0
}
