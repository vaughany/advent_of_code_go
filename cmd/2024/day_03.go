package main

import (
	"regexp"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day03() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day03part1,
		cfg.day03part2,
	)
}

// 2024-03-1: 174336360
func (cfg *config) day03part1(instructions []string) (total int) {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, ins := range instructions {
		results := regex.FindAllStringSubmatch(ins, -1)

		for _, res := range results {
			a := atoi3(res[1])
			b := atoi3(res[2])

			total += a * b
		}
	}

	return
}

func atoi3(s string) int {
	n := 0

	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}

	return n
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
					a := atoi3(res[1])
					b := atoi3(res[2])

					total += a * b
				}
			}
		}
	}

	return
}
