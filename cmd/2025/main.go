package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
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
	loader loaders.Loader
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
	cfg.year = 2025
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

	getInput := flag.Bool("get-input", false, "fetches today's input")
	getAllInputs := flag.Bool("get-all-inputs", false, "fetches all inputs for all years")
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

	if *getInput {
		err := support.GetPuzzleInput(cfg.year, cfg.day)
		if err != nil {
			return err
		}

		return nil
	}

	if *getAllInputs {
		err := support.GetAllPuzzleInputs(2015, cfg.year-1, true)
		if err != nil {
			return err
		}

		return nil
	}

	if cfg.runAllDays {
		// Just 12 days' worth of puzzles in 2025.
		for j := 1; j <= 12; j++ {
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

	cfg.loader = loaders.NewLoader(cfg.efs, cfg.year, day, cfg.sample)

	var err error

	switch day {
	case 1:
		err = cfg.day01()
	case 2:
		err = cfg.day02()
	case 3:
		err = cfg.day03()
	case 4:
		err = cfg.day04()
	case 5:
		err = cfg.day05()
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
	}

	if err != nil {
		return err
	}

	return nil
}
