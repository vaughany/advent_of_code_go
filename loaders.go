package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GetFilename creates the path and filename based on day and real / sample data.
func (cfg *config) getFilename(year, day int) string {
	var suffix string
	if cfg.sample {
		suffix = "-sample"
	}

	return fmt.Sprintf("inputs/%d/%02d%s.txt", year, day, suffix)
}

// Read a file with many lines and return an array (of strings).
func loadFile(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) < 1 {
		fmt.Printf("Input file '%s' has no lines.\n", filename)
		os.Exit(1)
	}

	return lines
}

// Takes a file where every line is an int, and returns a slice of ints.
func (cfg *config) getInputAsInts(year, day int) []int {
	var output []int

	for _, line := range loadFile(cfg.getFilename(year, day)) {
		int, _ := strconv.Atoi(line)
		output = append(output, int)
	}

	return output
}

// // Takes a one-line file and returns an int.
// func (cfg *config) getInputAsInt() int {
// 	return cfg.getInputAsInts()[0]
// }

// Takes a file where every line is a string, and returns a slice of strings.
func (cfg *config) getInputAsStrings(year, day int) []string {
	return loadFile(cfg.getFilename(year, day))
}

// Takes a one-line file and returns a string.
func (cfg *config) getInputAsString(year, day int) string {
	return loadFile(cfg.getFilename(year, day))[0]
}

func (cfg *config) getInputAsBytes(year, day int) []byte {
	return []byte(loadFile(cfg.getFilename(year, day))[0])
}

// // Takes a one-line file where each data point is a comma-separated int, and returns a slice of ints.
// func (cfg *config) getCommaSeparatedInputAsInts() []int {
// 	var output []int

// 	split := strings.Split(loadFile(cfg.getFilename())[0], ",")
// 	for _, s := range split {
// 		int, _ := strconv.Atoi(s)
// 		output = append(output, int)
// 	}

// 	return output
// }
