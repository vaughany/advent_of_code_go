package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type config struct {
	sample, debug, timing bool
	day, year             int
	timings               struct {
		startup  time.Time
		finished time.Time
	}
	colours struct {
		title, info, timing, reset, nyi string
	}
	getInput      bool
	sessionCookie string
	// allDays bool
}

func main() {
	var cfg config
	cfg.timings.startup = time.Now()
	cfg.colours.title = "\u001b[32m"
	cfg.colours.info = "\u001b[33m"
	cfg.colours.timing = "\u001b[36m"
	cfg.colours.nyi = "\u001b[90m"
	// cfg.colours.reset = "\u001b[0m"

	flag.BoolVar(&cfg.debug, "d", cfg.debug, "Display debugging information")
	flag.BoolVar(&cfg.sample, "s", cfg.sample, "Use the sample input file")
	flag.BoolVar(&cfg.timing, "t", cfg.timing, "Display timing information")
	flag.IntVar(&cfg.year, "year", cfg.year, "Run just this year")
	flag.IntVar(&cfg.day, "day", cfg.day, "Run just this day")
	flag.BoolVar(&cfg.getInput, "get-today", cfg.getInput, "fetches today's input (requires -session)")
	flag.StringVar(&cfg.sessionCookie, "session", cfg.sessionCookie, "AoC session Cookie")
	flag.Parse()

	if cfg.getInput && len(cfg.sessionCookie) != 128 {
		fmt.Println("Sorry, a session cookie (128 chars) is required to use the -get-input function.")
		os.Exit(1)
	}

	if cfg.getInput && len(cfg.sessionCookie) == 128 {
		// Run with go run . -get-today -session=aoc-session-cookie-goes-here

		cfg.title(time.Now().Year(), 0)

		client := &http.Client{}

		now := time.Now()
		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", now.Year(), now.Day())
		fmt.Printf("Fetching %s...\n", url)

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}

		cookie := &http.Cookie{
			Name:  "session",
			Value: cfg.sessionCookie,
		}
		req.AddCookie(cookie)

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Expected 200 OK, received %d %s.\n", resp.StatusCode, http.StatusText(resp.StatusCode))
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		filename := fmt.Sprintf("inputs/%d/%.2d.txt", now.Year(), now.Day())
		err = os.WriteFile(filename, bodyBytes[:len(bodyBytes)-1], 0644) // 0600
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d bytes writtern to %s.\n", len(bodyBytes), filename)

		os.Exit(0)
	}

	// if cfg.day == 0 {
	// 	cfg.allDays = true
	// }

	cfg.run()

	cfg.timings.finished = time.Now()
}

func (cfg *config) run() {
	if cfg.year == 2015 || cfg.year == 0 {
		cfg.run2015()
	}

	// if cfg.year == 2016 || cfg.year == 0 {
	// 	cfg.run2016()
	// }

	// if cfg.year == 2017 || cfg.year == 0 {
	// 	cfg.run2017()
	// }

	// if cfg.year == 2018 || cfg.year == 0 {
	// 	cfg.run2018()
	// }

	// if cfg.year == 2019 || cfg.year == 0 {
	// 	cfg.run2019()
	// }

	// if cfg.year == 2020 || cfg.year == 0 {
	// 	cfg.run2020()
	// }

	// if cfg.year == 2021 || cfg.year == 0 {
	// 	cfg.run2021()
	// }

	if cfg.year == 2022 || cfg.year == 0 {
		cfg.run2022()
	}
}
