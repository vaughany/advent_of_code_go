package main

import (
	"slices"
	"strconv"
	"strings"
	"time"

	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day05() error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := cfg.loader.GetStrings()
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day05part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day05part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

type Fresh struct {
	from, to int
}

// 2025-05-1: 739
func (cfg *config) day05part1(instructions []string) int {
	freshRanges := make([]Fresh, 0, 256)
	ingredients := make([]int, 0, 1024)

	for _, line := range instructions {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if fromStr, toStr, ok := strings.Cut(line, "-"); ok {
			from, _ := strconv.Atoi(fromStr)
			to, _ := strconv.Atoi(toStr)

			freshRanges = append(freshRanges, Fresh{from: from, to: to})

		} else {
			ingredient, _ := strconv.Atoi(line)

			ingredients = append(ingredients, ingredient)
		}
	}

	totalIngredients := 0
	for _, ingredient := range ingredients {
		if isFresh(freshRanges, ingredient) {
			totalIngredients++
		}
	}

	return totalIngredients
}

func isFresh(ranges []Fresh, ingredient int) bool {
	for _, fr := range ranges {
		if ingredient >= fr.from && ingredient <= fr.to {
			return true
		}
	}

	return false
}

// 2025-05-2: 344486348901788
func (cfg *config) day05part2(instructions []string) int {
	freshRanges := make([]Fresh, 0, 256)

	for _, line := range instructions {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if fromStr, toStr, ok := strings.Cut(line, "-"); ok {
			from, _ := strconv.Atoi(fromStr)
			to, _ := strconv.Atoi(toStr)

			freshRanges = append(freshRanges, Fresh{from: from, to: to})
		}
	}

	// Sorting.
	slices.SortFunc(freshRanges, func(a, b Fresh) int {
		if a.from != b.from {
			if a.from < b.from {
				return -1
			}

			return 1
		}

		if a.to < b.to {
			return -1
		}

		if a.to > b.to {
			return 1
		}

		return 0
	})

	writeIdx := 0
	cur := freshRanges[0]

	// Merge (ugh).
	for i := 1; i < len(freshRanges); i++ {
		r := freshRanges[i]
		if r.from <= cur.to+1 {
			// If ranges overlap or tough, add to / extend current.
			if r.to > cur.to {
				cur.to = r.to
			}

		} else {
			// No? Create a new range.
			freshRanges[writeIdx] = cur
			writeIdx++
			cur = r
		}
	}

	// Push the last ranges.
	freshRanges[writeIdx] = cur
	writeIdx++

	merged := freshRanges[:writeIdx]

	total := 0
	for _, fr := range merged {
		total += fr.to - fr.from + 1 // The +1 includes the 'end' of each range.
	}

	return total
}
