package main

import (
	"flag"
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
		title, info, timing, reset string
	}
	// allDays bool
}

func main() {
	var cfg config
	cfg.timings.startup = time.Now()
	cfg.colours.title = "\u001b[32m"
	cfg.colours.info = "\u001b[33m"
	cfg.colours.timing = "\u001b[36m"
	// cfg.colours.reset = "\u001b[0m"

	flag.BoolVar(&cfg.debug, "d", cfg.debug, "Display debugging information")
	flag.BoolVar(&cfg.sample, "s", cfg.sample, "Use the sample input data")
	flag.BoolVar(&cfg.timing, "t", cfg.timing, "Display timing information")
	flag.IntVar(&cfg.year, "year", cfg.year, "Run just this day")
	flag.IntVar(&cfg.day, "day", cfg.day, "Run just this day")
	flag.Parse()

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
