package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"vaughany.com/advent_of_code_go/internal/aoc_config"
	"vaughany.com/advent_of_code_go/internal/loaders"
	"vaughany.com/advent_of_code_go/internal/output"
	"vaughany.com/advent_of_code_go/internal/support"
)

//go:embed inputs/*.txt
var embeddedFiles embed.FS

type config struct {
	sample, debug, timing, runAllDays bool
	day, year                         int
	efs                               embed.FS
	timings                           struct {
		startup  time.Time
		finished time.Time
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var cfg config
	cfg.timings.startup = time.Now()
	cfg.efs = embeddedFiles
	cfg.year = 2024
	cfg.day = time.Now().Day()

	flag.Usage = func() {
		output.Info(aoc_config.GetVersionInfo(cfg.year))
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&cfg.debug, "d", cfg.debug, "display debugging information")
	flag.BoolVar(&cfg.sample, "s", cfg.sample, "use the sample input file")
	flag.BoolVar(&cfg.timing, "t", cfg.timing, "display timing information")
	flag.IntVar(&cfg.day, "day", cfg.day, "run only this day")
	flag.BoolVar(&cfg.runAllDays, "all-days", cfg.runAllDays, "run all days (overrides '-day')")

	getInput := flag.Bool("get-input", false, "fetches today's input (requires -cookie)")
	sessionCookie := flag.String("cookie", "", "AoC session cookie, 128 chars")
	webserver := flag.Bool("webserver", false, "launch web server to see input files")
	version := flag.Bool("version", false, "version info")
	listFiles := flag.Bool("list-files", false, "list embedded files")
	flag.Parse()

	output.Title(cfg.year)

	if *webserver {
		address := "127.0.0.1:8080"

		err := support.ServeInputs(&cfg.efs, address)
		if err != nil {
			return err
		}

		return nil
	}

	if *version {
		output.Info(aoc_config.GetVersionInfo(cfg.year))

		return nil
	}

	if *listFiles {
		err := support.GetAllFilenames(&cfg.efs)
		if err != nil {
			return err
		}

		return nil
	}

	// Run with e.g. `go run ./cmd/2024/ -get-input` to use the cookie stored in `.cookie`
	//       or e.g. `go run ./cmd/2024/ -get-input -cookie aoc-session-cookie-goes-here`
	if *getInput {
		// Try loading the cookie from a file first.
		if len(*sessionCookie) == 0 {
			bytes, err := os.ReadFile(`.cookie`)
			if err != nil {
				return err
			}

			// Save the text from the file into the session cookie variable.
			*sessionCookie = strings.TrimSpace(string(bytes))
		}

		// If the cookie, no matter it's origin, is not exactly 128 chars, bail.
		if len(*sessionCookie) != 128 {
			return errors.New("sorry, a session cookie (128 chars) from the AoC website is required to use the -get-input function")
		}

		err := support.GetPuzzleInput(cfg.year, cfg.day, *sessionCookie)
		if err != nil {
			return err
		}

		return nil
	}

	if cfg.runAllDays {
		for j := 1; j <= 25; j++ {
			err := cfg.runDay(j)
			if err != nil {
				return err
			}
		}

		cfg.timings.finished = time.Now()
		output.TimeInfo(output.InfoTypeWholeRun, time.Since(cfg.timings.startup))

		return nil
	}

	err := cfg.runDay(cfg.day)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *config) runDay(day int) error {
	output.Subtitle(day)

	loader := loaders.NewLoader(cfg.efs, cfg.year, day, cfg.sample)

	var err error

	switch day {
	case 1:
		err = cfg.day01(loader)
	case 2:
		err = cfg.day02(loader)
	case 3:
		err = cfg.day03(loader)
	case 4:
		err = cfg.day04(loader)
	case 5:
		output.NotYetImplemented()
	case 6:
		output.NotYetImplemented()
	case 7:
		output.NotYetImplemented()
	case 8:
		output.NotYetImplemented()
	case 9:
		output.NotYetImplemented()
	case 10:
		output.NotYetImplemented()
	case 11:
		output.NotYetImplemented()
	case 12:
		output.NotYetImplemented()
	case 13:
		output.NotYetImplemented()
	case 14:
		output.NotYetImplemented()
	case 15:
		output.NotYetImplemented()
	case 16:
		output.NotYetImplemented()
	case 17:
		output.NotYetImplemented()
	case 18:
		output.NotYetImplemented()
	case 19:
		output.NotYetImplemented()
	case 20:
		output.NotYetImplemented()
	case 21:
		output.NotYetImplemented()
	case 22:
		output.NotYetImplemented()
	case 23:
		output.NotYetImplemented()
	case 24:
		output.NotYetImplemented()
	case 25:
		output.NotYetImplemented()
	}

	if err != nil {
		return err
	}

	return nil
}
