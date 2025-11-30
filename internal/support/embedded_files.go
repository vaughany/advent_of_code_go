package support

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"vaughany.com/advent_of_code_go/internal/output"
)

// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc?permalink_comment_id=4405804#gistcomment-4405804
func GetAllFilenames(efs *embed.FS) error {
	var files []string

	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		files = append(files, fmt.Sprintf("%s\t%d", path[7:], info.Size()))

		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("Files:", len(files))

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	for _, file := range files {
		fmt.Fprintln(tw, file)
	}
	tw.Flush()

	return nil
}

func loadCookie() (string, error) {
	// Load cookie from file
	cookieBytes, err := os.ReadFile(".cookie")
	if err != nil {
		return "", fmt.Errorf("failed to read .cookie file: %w", err)
	}
	cookie := strings.TrimSpace(string(cookieBytes))

	// Validate cookie length
	if len(cookie) != 128 {
		return "", errors.New("session cookie must be exactly 128 characters")
	}

	return cookie, nil
}

// GetPuzzleInput downloads a puzzle input file for the given year and day.
// It reads the session cookie from a .cookie file in the current directory.
// The session cookie must be exactly 128 characters.
func GetPuzzleInput(year, day int) error {
	// Load that cookie.
	cookie, err := loadCookie()
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	output.Info(fmt.Sprintf("Fetching %s...", url))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected 200 OK, received %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Make the inputs folder for the specified year, if it doesn't exist.
	err = os.MkdirAll(fmt.Sprintf("./cmd/%d/inputs", year), os.ModePerm)
	if err != nil {
		return err
	}

	// This path references the path within the code repository and will not work if using a compiled binary.
	filename := fmt.Sprintf("./cmd/%d/inputs/%.2d.txt", year, day)
	err = os.WriteFile(filename, bodyBytes, 0664)
	if err != nil {
		return err
	}

	output.Info(fmt.Sprintf("%d bytes written to %s.", len(bodyBytes), filename))

	return nil
}

// ServeInputs creates a http web server on the supplied address, and serves the inputs from the supplied embedded file system.
func ServeInputs(efs *embed.FS, address string) error {
	serverRoot, err := fs.Sub(efs, "inputs")
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.FS(serverRoot)))

	output.Info("Running web server on http://" + address)

	err = http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetAllPuzzleInputs downloads Advent of Code input files for a range of years and days.
// It reads the session cookie from a .cookie file in the current directory.
// startYear and endYear define the range of years to download (inclusive).
// skipExisting will skip files that already exist if true.
func GetAllPuzzleInputs(startYear, endYear int, skipExisting bool) error {
	failedDownloads := 0

	for year := startYear; year <= endYear; year++ {
		output.Info(fmt.Sprintf("Year %d", year))

		for day := 1; day <= 25; day++ {
			filename := fmt.Sprintf("./cmd/%d/inputs/%02d.txt", year, day)

			// Skip if file exists and skipExisting is true
			if skipExisting {
				if fi, err := os.Stat(filename); err == nil && fi.Size() > 0 {
					output.Info(fmt.Sprintf("skipping %d day %02d (already exists and is not empty)", year, day))
					continue
				}
			}

			output.Info(fmt.Sprintf("Downloading %d day %02d", year, day))
			err := GetPuzzleInput(year, day)
			if err != nil {
				// Check if it's a 404 error (day not available)
				errStr := err.Error()
				if strings.Contains(errStr, "404") || strings.Contains(errStr, "Not Found") {
					output.Info(fmt.Sprintf("day %02d not available (404)", day))
					continue
				}

				// For other errors, log and count as failure
				output.Info(fmt.Sprintf("failed to download: %v", err))
				failedDownloads++

				continue
			}

			// Rate limiting delay
			sleepSeconds := 5 + rand.Intn(6)
			output.Info(fmt.Sprintf("sleeping for %d seconds", sleepSeconds))
			time.Sleep(time.Second * time.Duration(sleepSeconds))
		}
	}

	if failedDownloads > 0 {
		return fmt.Errorf("completed with %d failed download(s)", failedDownloads)
	}

	output.Info("all downloads completed successfully!")
	return nil
}
