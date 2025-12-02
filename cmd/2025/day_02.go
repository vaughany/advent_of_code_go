package main

import (
	"strconv"
	"strings"
	"time"

	"vaughany.com/advent_of_code_go/internal/output"
)

func (cfg *config) day02() error {
	timingStart := time.Now()

	// 'instructions' can vary in type, depending on if we're dealing with ints, strings, bytes etc.
	instructions, err := cfg.loader.GetString()
	if err != nil {
		return err
	}
	if cfg.timing {
		output.TimeInfo(output.InfoTypeSetup, time.Since(timingStart))
	}

	timingPartOne := time.Now()
	output.AnswerPart1(cfg.day02part1(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeOne, time.Since(timingPartOne))
	}

	timingPartTwo := time.Now()
	output.AnswerPart2(cfg.day02part2(instructions))
	if cfg.timing {
		output.TimeInfo(output.InfoTypeTwo, time.Since(timingPartTwo))
	}

	if cfg.timing {
		output.TimeInfo(output.InfoTypeBoth, time.Since(timingPartOne))
		output.TimeInfo(output.InfoTypeEverything, time.Since(timingStart))
	}

	return nil
}

// 2025-02-1: 30608905813
func (cfg *config) day02part1(instructions string) int {
	invalidIDs := 0

	for r := range strings.SplitSeq(instructions, ",") {
		rangeSlice := strings.Split(r, "-")

		// Probs should handle these errors.
		first, _ := strconv.Atoi(rangeSlice[0])
		last, _ := strconv.Atoi(rangeSlice[1])

		for i := first; i <= last; i++ {
			invalidIDs += cfg.day02CheckInvalid(i)
		}
	}

	return invalidIDs
}

// invalidIDsany ID which is made only of some sequence of digits repeated twice. So, 55 (5 twice), 6464 (64 twice),
// and 123123 (123 twice) would all be invalid IDs.
func (cfg *config) day02CheckInvalid(in int) int {
	inStr := strconv.Itoa(in)

	strLen := len(inStr)

	// Ignore strings with an odd length.
	if strLen%2 == 1 {
		return 0
	}

	middle := strLen / 2
	if inStr[:middle] == inStr[middle:] {
		return in
	}

	return 0
}

// 2025-02-2: 31898925685
func (cfg *config) day02part2(instructions string) int {
	invalidIDs := 0

	for r := range strings.SplitSeq(instructions, ",") {
		rangeSlice := strings.Split(r, "-")

		// Probs should handle these errors.
		first, _ := strconv.Atoi(rangeSlice[0])
		last, _ := strconv.Atoi(rangeSlice[1])

		for i := first; i <= last; i++ {
			invalidIDs += cfg.day02CheckInvalid2(i)
		}
	}

	return invalidIDs
}

var blockPatterns = map[int][]int{
	2:  {2},
	3:  {3},
	4:  {2, 4},
	5:  {5},
	6:  {2, 3},
	7:  {7},
	8:  {2, 4, 8},
	9:  {3, 9},
	10: {2, 5, 10},
}

// Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice.
// So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times),
// and 1111111 (1 seven times) are all invalid IDs.
func (cfg *config) day02CheckInvalid2(in int) int {
	inStr := strconv.Itoa(in)

	strLen := len(inStr)

	if strLen < 2 || strLen > 10 {
		return 0
	}

	// Are all digits the same?
	if allSame(inStr) {
		return in
	}

	// Any of the allowed repeated-block patterns for this length?
	if blocks, ok := blockPatterns[strLen]; ok {
		for _, b := range blocks {
			if repeatedBlocks(inStr, b) {
				return in
			}
		}
	}

	return 0
}

// Returns true if all characters in 'in' are identical.
func allSame(in string) bool {
	if len(in) == 0 {
		return false
	}

	for i := 1; i < len(in); i++ {
		if in[i] != in[0] {
			return false
		}
	}

	return true
}

// Reports whether s can be split into `blocks` equal-sized contiguous chunks and all chunks are identical.
func repeatedBlocks(in string, blocks int) bool {
	n := len(in)
	if blocks <= 0 || n%blocks != 0 {
		return false
	}

	chunkLen := n / blocks
	first := in[:chunkLen]

	for i := chunkLen; i < n; i += chunkLen {
		if in[i:i+chunkLen] != first {
			return false
		}
	}

	return true
}
