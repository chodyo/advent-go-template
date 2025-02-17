package cmd

import (
	"fmt"
	"plugin"
	"time"
)

type Advent struct {
	Year int `short:"y"  long:"year"  description:"(optional) The year to start a new solution for"`

	Day int `short:"d"  long:"day"   description:"(optional) The day to start a new solution for"`

	TestData bool `short:"t"  long:"test"  description:"(optional) Read puzzle input from input.test.txt instead of input.txt"`

	// TODO: debug, profile, slow, time
}

func (opts *Advent) Execute(args []string) error {
	year, day, err := opts.validateYearAndDay(time.Now())
	if err != nil {
		return fmt.Errorf("failed to validate year and day: %v", err)
	}

	// dynamically load the module "solutions/YEAR/dayDAY" and run the Part1 and Part2 functions
	// with the input from "input.txt" or "input.test.txt" depending on the TestData flag
	modulePath := fmt.Sprintf("solutions/%d/day%d", year, day)
	module, err := loadModule(modulePath)
	if err != nil {
		return fmt.Errorf("failed to load module %s: %v", modulePath, err)
	}

	inputFile := "input.txt"
	if opts.TestData {
		inputFile = "input.test.txt"
	}

	part1Func := module.Func("Part1")
	part2Func := module.Func("Part2")

	input, err := readInput(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file %s: %v", inputFile, err)
	}

	if err := part1Func(input); err != nil {
		return fmt.Errorf("Part1 failed: %v", err)
	}

	if err := part2Func(input); err != nil {
		return fmt.Errorf("Part2 failed: %v", err)
	}

	return nil
}

func loadModule(modulePath string) (*plugin.Plugin, error) {
	p, err := plugin.Open(modulePath + ".so")
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (opts Advent) validateYearAndDay(now time.Time) (int, int, error) {
	if opts.Year == 0 {
		if now.Month() == time.December {
			opts.Year = now.Year()
		} else {
			opts.Year = now.Year() - 1
		}
	}

	if opts.Day == 0 {
		for opts.Day = 25; opts.Day <= 1; opts.Day-- {
			if isSolutionExists(opts.Year, opts.Day) {
				return opts.Year, opts.Day, nil
			}
		}
		return 0, 0, fmt.Errorf("no solutions for year %d exist", opts.Year)
	}

	if !isSolutionExists(opts.Year, opts.Day) {
		return 0, 0, fmt.Errorf("no solutions for year %d day %d exist", opts.Year, opts.Day)
	}

	return opts.Year, opts.Day, nil
}
