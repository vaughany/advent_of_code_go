package loaders

import (
	"bufio"
	"embed"
	"errors"
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
	var lines []string

	file, err := l.fs.Open(l.filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 1 {
		return lines, errors.New(fmt.Sprintf("input file '%s' has no lines", l.filename))
	}

	return lines, nil
}

// GetStrings takes a file where every line is a string, and returns a slice of strings.
func GetStrings(loader Loader) ([]string, error) {
	var output []string

	output, err := loader.loadFile()
	if err != nil {
		return output, err
	}

	return output, nil
}

// GetString takes a one-line file and returns a string.
func GetString(loader Loader) (string, error) {
	var output []string

	output, err := GetStrings(loader)
	if err != nil {
		return "", err
	}

	return output[0], nil
}

// GetInts takes a file where every line is an int, and returns a slice of ints.
func GetInts(loader Loader) ([]int, error) {
	var output []int

	data, err := GetStrings(loader)
	if err != nil {
		return output, err
	}

	for _, line := range data {
		var int int

		// If the line is blank, return 0. TODO: this may bite me in the butt at some point.
		if line == "" {
			int = 0

		} else {
			int, err = strconv.Atoi(line)
			if err != nil {
				return output, err
			}
		}

		output = append(output, int)
	}

	return output, nil
}

// GetInt takes a one-line file and returns an int.
func GetInt(loader Loader) (int, error) {
	var output []int

	output, err := GetInts(loader)
	if err != nil {
		return 0, err
	}

	return output[0], nil
}

// GetBytes takes a one-line file and returns a slice of bytes.
func GetBytes(loader Loader) ([]byte, error) {
	var output string

	output, err := GetString(loader)
	if err != nil {
		return []byte{}, err
	}

	return []byte(output), nil
}

// GetCSInts takes a one-line file where each data point is a comma-separated int, and returns a slice of ints.
func GetCSInts(loader Loader) ([]int, error) {
	var output []int

	data, err := GetString(loader)
	if err != nil {
		return output, err
	}

	split := strings.Split(data, ",")
	for _, s := range split {
		int, err := strconv.Atoi(s)
		if err != nil {
			return output, err
		}

		output = append(output, int)
	}

	return output, nil
}
