package main

import (
	"regexp"
	"strconv"
	"time"

	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day03(loader loaders.Loader) error {
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
	output.AnswerPart1(cfg.day03part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day03part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2024-03-1: 174336360
func (cfg *config) day03part1(instructions []string) (total int) {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, ins := range instructions {
		results := regex.FindAllStringSubmatch(ins, -1)

		for _, res := range results {
			a, _ := strconv.Atoi(res[1])
			b, _ := strconv.Atoi(res[2])

			total += a * b
		}
	}

	return
}

// 2024-03-2: 88802350
func (cfg *config) day03part2(instructions []string) (total int) {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	do := true

	for _, ins := range instructions {
		results := regex.FindAllStringSubmatch(ins, -1)

		for _, res := range results {

			switch res[0][:3] {
			case "do(":
				do = true
			case "don":
				do = false
			case "mul":
				if do {
					a, _ := strconv.Atoi(res[1])
					b, _ := strconv.Atoi(res[2])

					total += a * b
				}
			}
		}
	}

	return
}
