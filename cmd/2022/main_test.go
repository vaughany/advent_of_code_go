package main

import (
	"testing"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func TestDay01(t *testing.T) {
	var (
		cfg           config
		part1expected = 70369
		part2expected = 203002
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 1, false)
	instructions, err := loaders.GetInts(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day01part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay01part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day01part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay01part2: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay02(t *testing.T) {
	var (
		cfg           config
		part1expected = 11841
		part2expected = 13022
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 2, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day02part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay02part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day02part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay02part2: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay03(t *testing.T) {
	var (
		cfg           config
		part1expected = 8018
		part2expected = 2518
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 3, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day03part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay03part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day03part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay03part2: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay04(t *testing.T) {
	var (
		cfg           config
		part1expected = 573
		part2expected = 867
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 4, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual, part2actual := cfg.day04bothParts(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay04part1: expected %d, actual %d", part1expected, part1actual)
	}

	if part2expected != part2actual {
		t.Errorf("TestDay04part1: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay05(t *testing.T) {
	var (
		cfg           config
		part1expected = "CFFHVVHNC"
		part2expected = "FSZWBPTBG"
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 5, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day05part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay05part1: expected %s, actual %s", part1expected, part1actual)
	}

	part2actual := cfg.day05part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay05part2: expected %s, actual %s", part2expected, part2actual)
	}
}

func TestDay06(t *testing.T) {
	var (
		cfg           config
		part1expected = 1175
		part2expected = 3217
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 6, false)
	instructions, err := loaders.GetString(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day06part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay06part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day06part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay06part2: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay07(t *testing.T) {
	t.Skip("Not yet implemented.")
}

func TestDay08(t *testing.T) {
	var (
		cfg           config
		part1expected = 1776
		part2expected = 234416
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 8, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day08part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay08part1: expected %d, actual %d", part1expected, part1actual)
	}

	part2actual := cfg.day08part2(instructions)

	if part2expected != part2actual {
		t.Errorf("TestDay08part2: expected %d, actual %d", part2expected, part2actual)
	}
}

func TestDay09(t *testing.T) {
	t.Skip("Not yet implemented.")
}

func TestDay10(t *testing.T) {
	var (
		cfg           config
		part1expected = 15680
		// 		part2expected = `#### #### ###  #### #  #  ##  #  # ###
		//    # #    #  # #    #  # #  # #  # #  #
		//   #  ###  ###  ###  #### #    #  # #  #
		//  #   #    #  # #    #  # # ## #  # ###
		// #    #    #  # #    #  # #  # #  # #
		// #### #    ###  #    #  #  ###  ##  #`
	)

	loader := loaders.NewLoader(embeddedFiles, 2022, 10, false)
	instructions, err := loaders.GetStrings(loader)
	if err != nil {
		t.Error(err)
	}

	part1actual := cfg.day10part1(instructions)

	if part1expected != part1actual {
		t.Errorf("TestDay10part1: expected %d, actual %d", part1expected, part1actual)
	}

	// part2actual := cfg.day10part2(instructions)
	//
	// if part2expected != part2actual {
	// t.Errorf("TestDay10part2: expected %d, actual %d", part2expected, part2actual)
	// }
}
