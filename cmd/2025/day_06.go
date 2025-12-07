package main

import (
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day06() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day06part1,
		cfg.day06part2,
	)
}

// 2025-06-1: 6503327062445
func (cfg *config) day06part1(instructions []string) int {
	nLines := len(instructions)
	lastIdx := nLines - 1

	rows := make([][]int, 0, lastIdx)
	var operations []byte

	for i, instruction := range instructions {
		fields := strings.Fields(instruction)

		if i < lastIdx {
			// Numeric row.
			row := make([]int, len(fields))

			for j, field := range fields {
				val, _ := strconv.Atoi(field)

				row[j] = val
			}

			rows = append(rows, row)

		} else {
			// Operator row.
			operations = make([]byte, len(fields))

			for j, field := range fields {
				operations[j] = field[0]
			}
		}
	}

	var (
		numRows = len(rows)
		numCols = len(rows[0])
		total   = 0
	)

	for col := 0; col < numCols; col++ {
		op := operations[col]

		acc := rows[0][col]
		for row := 1; row < numRows; row++ {
			val := rows[row][col]

			if op == '*' {
				acc *= val
			} else {
				// Assume '+' for this puzzle; if you want to be explicit:
				// if op == '+' { acc += val }
				acc += val
			}
		}

		total += acc
	}

	return total
}

// 2025-06-2: 9640641878593
func (cfg *config) day06part2(instructions []string) int {
	rows := len(instructions)
	cols := len(instructions[0])

	total := 0

	// Set initial state.
	inProblem := false
	group := 0
	var currentOp byte // '*' or '+', 0 if not set yet

	for col := range cols {
		// Build the digit string for this column (excluding the operator row).
		var sb strings.Builder

		for row := 0; row < rows-1; row++ {
			ch := instructions[row][col]

			if ch >= '0' && ch <= '9' {
				sb.WriteByte(ch)
			}
		}

		digitStr := sb.String()
		hasDigits := len(digitStr) > 0

		// Operator lives on the last row (if any) for this column.
		bottom := instructions[rows-1][col]
		hasOp := bottom == '*' || bottom == '+'

		// No digits and no operator = separator column and end of this problem.
		if !hasDigits && !hasOp {
			if inProblem {
				total += group

				inProblem = false
				group = 0
				currentOp = 0
			}

			continue
		}

		// If this column has an operator, remember it for this problem.
		if hasOp {
			currentOp = bottom
		}

		// If there are digits in this column, turn them into an int.
		if hasDigits {
			val, err := strconv.Atoi(digitStr)
			if err != nil {
				// AoC input should be clean; you could panic here if you prefer.
				continue
			}

			if !inProblem {
				// First number in this problem.
				inProblem = true
				group = val

			} else {
				// Subsequent number; apply the current operator.
				switch currentOp {
				case '*':
					group *= val
				case '+':
					group += val
				}
			}
		}
	}

	// If we finished while still "in a problem", flush the last one.
	if inProblem {
		total += group
	}

	return total
}
