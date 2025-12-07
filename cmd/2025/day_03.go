package main

import (
	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day03() error {
	return runDayWithInput(
		cfg,
		func(l loaders.Loader) ([]string, error) { return l.GetStrings() },
		cfg.day03part1,
		cfg.day03part2,
	)
}

// 2025-03-1: 17113
func (cfg *config) day03part1(instructions []string) int {
	total := 0

	for _, instruction := range instructions {
		var (
			length  = len(instruction)
			maxPair = 0
		)

		// Best digit we've seen so far to the right of the current index.
		bestDigit := -1

		// Scan from right to left. Apparently this is faster. (Note: yes it is: ms to µs.)
		for i := length - 1; i >= 0; i-- {
			digit := int(instruction[i] - '0')

			// If we already have at least one digit to the right, form a 2-digit number with this as the left digit.
			if bestDigit >= 0 {
				value := digit*10 + bestDigit
				if value > maxPair {
					maxPair = value
				}
			}

			// Update the best digit we've seen to the right.
			if digit > bestDigit {
				bestDigit = digit
			}
		}

		total += maxPair
	}

	return total
}

// 2025-03-2: 169709990062889
func (cfg *config) day03part2(instructions []string) int64 {
	var total int64

	for _, instruction := range instructions {
		length := len(instruction)

		var (
			start        = 0
			need         = 12
			chosenDigits = make([]int, 0, 12)
		)

		for need > 0 {
			maxStart := length - need
			bestDigit, bestIndex := -1, -1

			for i := start; i <= maxStart; i++ {
				d := int(instruction[i] - '0')

				// Update the best digit and it's index.
				if d > bestDigit {
					bestDigit = d
					bestIndex = i
				}
			}

			// Record it.
			chosenDigits = append(chosenDigits, bestDigit)

			// Update start position and decrement need.
			start = bestIndex + 1
			need--
		}

		total += cfg.day03DigitsToInt(chosenDigits)
	}

	return total
}

func (cfg *config) day03DigitsToInt(digits []int) int64 {
	var n int64 = 0

	for _, d := range digits {
		n = n*10 + int64(d)
	}

	return n
}
