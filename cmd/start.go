package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/mod/modfile"
)

type Start struct {
	Year int `short:"y"  long:"year"  description:"(optional) The year to start a new solution for"`
	Day  int `short:"d"  long:"day"   description:"(optional) The day to start a new solution for"`
}

func (opts Start) Execute(args []string) error {
	now := time.Now()

	year, day, err := validateYearAndDay(opts, now)
	if err != nil {
		return err
	}

	solutionDir := filepath.Join("solutions", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day))
	err = os.MkdirAll(solutionDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create solution directory: %v", err)
	}

	readmePath := filepath.Join(solutionDir, "README.md")
	err = generateReadme(year, day, readmePath)
	if err != nil {
		return fmt.Errorf("failed to generate README: %v", err)
	}

	inputFiles := []string{"input.txt", "input.test.txt"}
	for _, filename := range inputFiles {
		inputPath := filepath.Join(solutionDir, filename)
		file, err := os.Create(inputPath)
		if err != nil {
			return fmt.Errorf("failed to create %s: %v", filename, err)
		}
		file.Close()
	}

	solutionPath := filepath.Join(solutionDir, "solution.go")
	err = generateSolutionGo(day, solutionPath)
	if err != nil {
		return fmt.Errorf("failed to generate solution.go: %v", err)
	}

	fmt.Printf("Created solution template for Advent of Code %d Day %d\n", year, day)
	return nil
}

func isSolutionExists(year int, day int) bool {
	solutionPath := filepath.Join("solutions", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day))
	_, err := os.Stat(solutionPath)
	return err == nil
}

func validateYearAndDay(opts Start, now time.Time) (int, int, error) {
	if opts.Year == 0 {
		if now.Month() == time.December {
			opts.Year = now.Year()
		} else {
			opts.Year = now.Year() - 1
		}
	}

	if opts.Day == 0 {
		for opts.Day = 1; opts.Day <= 25; opts.Day++ {
			if !isSolutionExists(opts.Year, opts.Day) {
				return opts.Year, opts.Day, nil
			}
		}
		return 0, 0, fmt.Errorf("all solutions for year %d already exist", opts.Year)
	}

	if opts.Day < 1 || opts.Day > 25 {
		opts.Day = 1
	}

	return opts.Year, opts.Day, nil
}

func generateReadme(year, day int, path string) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch problem description: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch problem description: HTTP %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read problem description: %v", err)
	}

	return os.WriteFile(path, content, 0644)
}

func generateSolutionGo(day int, path string) error {
	moduleName, err := getCurrentModuleName()
	if err != nil {
		return fmt.Errorf("failed to get current module name: %v", err)
	}

	content := fmt.Sprintf(`package day%02d

import "%s/solutions"

type Solution struct{}

// Part1 is the function you will implement for the AoC part 1 challenge.
func (s Solution) Part1(input solutions.Input) int {
	return 0
}

// Part2 is the function you will implement for the AoC part 2 challenge.
func (s Solution) Part2(input solutions.Input) int {
	return 0
}

// Answer1 is an optional method for you to save your answer for part 1 after you've solved the challenge.
func (s Solution) Answer1() int {
	return 0
}

// Answer2 is an optional method for you to save your answer for part 2 after you've solved the challenge.
func (s Solution) Answer2() int {
	return 0
}
`, day, moduleName)

	return os.WriteFile(path, []byte(content), 0644)
}

func getCurrentModuleName() (string, error) {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %v", err)
	}

	module, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return "", fmt.Errorf("failed to parse go.mod: %v", err)
	}

	return module.Module.Mod.Path, nil
}
