package loaders

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

type Loader struct {
	fs        embed.FS
	year, day int
	sample    bool
	filename  string
}

func NewLoader(fs embed.FS, year, day int, sample bool) Loader {
	suffix := ""
	if sample {
		suffix = "-sample"
	}

	filename := fmt.Sprintf("inputs/%02d%s.txt", day, suffix)

	return Loader{
		fs:       fs,
		year:     year,
		day:      day,
		sample:   sample,
		filename: filename,
	}
}

// loadFile reads a file from embedded filesystem with one or more lines and return an array of strings.
func (l *Loader) loadFile() ([]string, error) {
	file, err := l.fs.Open(l.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) < 1 {
		return nil, fmt.Errorf("input file '%s' has no lines", l.filename)
	}

	return lines, nil
}

// GetStrings takes a file where every line is a string, and returns a slice of strings.
func GetStrings(loader Loader) ([]string, error) {
	return loader.loadFile()
}

// GetString takes a one-line file and returns a string.
func GetString(loader Loader) (string, error) {
	file, err := loader.fs.Open(loader.filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", fmt.Errorf("input file '%s' has no lines", loader.filename)
	}

	return scanner.Text(), scanner.Err()
}

// GetInts takes a file where every line is an int, and returns a slice of ints.
func GetInts(loader Loader) ([]int, error) {
	data, err := GetStrings(loader)
	if err != nil {
		return nil, err
	}

	output := make([]int, 0, len(data))
	for _, line := range data {
		// If the line is blank, return 0. TODO: this may bite me in the butt at some point.
		if line == "" {
			output = append(output, 0)
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		output = append(output, val)
	}

	return output, nil
}

// GetInt takes a one-line file and returns an int.
func GetInt(loader Loader) (int, error) {
	file, err := loader.fs.Open(loader.filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("input file '%s' has no lines", loader.filename)
	}

	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// If the line is blank, return 0. TODO: this may bite me in the butt at some point.
	if line == "" {
		return 0, nil
	}

	val, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// Get2DInts takes a file where every line is multiple strings separated by spaces, and returns [][]int (2D slice of ints).
func Get2DInts(loader Loader) ([][]int, error) {
	input, err := GetStrings(loader)
	if err != nil {
		return nil, err
	}

	output := make([][]int, 0, len(input))
	for _, lineStr := range input {
		lineSlice := strings.Fields(lineStr) // Fields handles multiple spaces better than Split

		intSlice := make([]int, 0, len(lineSlice))
		for _, ls := range lineSlice {
			val, err := strconv.Atoi(ls)
			if err != nil {
				return nil, err
			}

			intSlice = append(intSlice, val)
		}

		output = append(output, intSlice)
	}

	return output, nil
}

// GetBytes takes a one-line file and returns a slice of bytes.
func GetBytes(loader Loader) ([]byte, error) {
	output, err := GetString(loader)
	if err != nil {
		return nil, err
	}

	return []byte(output), nil
}

// GetCSInts takes a one-line file where each data point is a comma-separated int, and returns a slice of ints.
func GetCSInts(loader Loader) ([]int, error) {
	data, err := GetString(loader)
	if err != nil {
		return nil, err
	}

	split := strings.Split(data, ",")

	output := make([]int, 0, len(split))
	for _, s := range split {
		int, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		output = append(output, int)
	}

	return output, nil
}
