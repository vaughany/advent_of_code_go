package main

import (
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day02() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day02part1,
		cfg.day02part2,
	)
}

// 2024-02-1: 510
func (cfg *config) day02part1(instructions []string) (safe int) {
	for _, ins := range instructions {
		report := day2ProcessInput(ins)
		if day2Safe(report) {
			safe++
		}
	}

	return
}

// 2024-02-2: 553
func (cfg *config) day02part2(instructions []string) (safe int) {
	for _, ins := range instructions {
		report := day2ProcessInput(ins)

		if day2Safe(report) || day2CheckLevels(report) {
			safe++
		}
	}

	return
}

func day2ProcessInput(in string) (out []int) {
	fields := strings.Fields(in)
	out = make([]int, 0, len(fields))

	for _, s := range fields {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out = append(out, v)
	}

	return
}

func day2CheckLevels(in []int) bool {
	n := len(in)
	if n <= 2 {
		// Removing any one level always leaves at least one element.
		return true
	}

	// Reusable buffer.
	buf := make([]int, n-1)

	for skip := 0; skip < n; skip++ {
		// Build buf = in with in[skip] removed
		pos := 0
		for i := 0; i < n; i++ {
			if i == skip {
				continue
			}
			buf[pos] = in[i]
			pos++
		}

		if day2Safe(buf[:pos]) {
			return true
		}
	}

	return false
}

func day2Safe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	var inc, dec bool

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff == 0 {
			return false
		}
		if diff > 0 {
			inc = true
		} else {
			dec = true
		}

		if diff < -3 || diff > 3 {
			return false
		}

		if inc && dec {
			return false
		}
	}

	return true
}
