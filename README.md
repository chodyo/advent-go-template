# @chodyo's Go Advent of Code Project Template

[![Go](https://github.com/chodyo/advent-go-template/actions/workflows/go.yml/badge.svg)](https://github.com/chodyo/advent-go-template/actions/workflows/go.yml)

This is my Go project template for the [Advent of Code](https://adventofcode.com) puzzles. It handles creating stub solutions, input parsing, and printing your answer, letting you focus on the actual solve.

This project is a Go adaption of [xavdid's Python template](https://github.com/xavdid/advent-of-code-python-template).

## Quickstart

1. Install Go (check this project's go.mod for the required minimum version)
2. Create a new repo from this template ([docs](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template#creating-a-repository-from-a-template)) and clone it locally.
3. Start a new solution using `go run . start`
4. Edit the newly created file at `solutions/YEAR/day_01/solution.go`.
5. Run your code answers using `go run . advent`.
6. Repeat and enjoy!

## Commands

This repo has two main commands: `start` and `advent`.

### `start`

#### Usage

`go run . start [-h] [--year YEAR] [day]`

Scaffold files to start a new Advent of Code solution.

##### positional arguments

* `day` (optional) Which puzzle day to start, between `[1,25]`. Defaults to the next day **without** a folder (matching `day_NN`) in the specified year.

##### optional arguments

* `-h, --help` (optional): show this help message and exit
* `--year YEAR` (optional): Puzzle year. Defaults to current year if December has begun, otherwise previous year.

#### Examples

- `go run . start`
- `go run . start 2`
- `go run . start --year 2023 3`

### `advent`

#### Usage

`go run . advent [--year year] [--test-data] [--debug] [--profile] [--slow] [--time] [day]`

Run a specific day of Advent of Code.

##### informational flags

- `-h, --help`: show this help message and exit
- `--version`: print the version number and exit

##### positional arguments

* `day` (optional) Which puzzle day to run, between `[1,25]`. Defaults to the next day **with** a folder (matching `day_NN`) in the specified year.

##### optional flags

* `--year YEAR` (optional): Puzzle year. Defaults to current year if December has begun, otherwise previous year.
* `-t, --test-data`: read puzzle input from `input.test.txt` instead of `input.txt`.
* `--debug`: print normally-hidden debug statements (written with `log.Debug(...)`). See [debugging](#debugging).
* `-profile`: run solution through a performance profiler
* `--slow`: specify that long-running solutions (or those requiring manual input) should be run. They're skipped by default.
* `--time`: print information about how long solutions took to run. More useful than timing at a shell level, since this only starts the timer once all boilerplate has been run.

#### Examples

- `go run . advent`
- `go run . advent 2`
- `go run . advent --year 2023 5`
- `go run . advent --test-data 7`
- `go run . advent --debug 9`

## File Structure

```
solutions/
├── ...
└── 2024/
    ├── day01/
    │   ├── solution.go
    │   ├── input.txt
    │   ├── input.test.txt
    │   └── README.md
    ├── day02/
    │   ├── solution.go
    │   ├── ...
    └── ...
```

## Writing Solutions

### The `Solution` Interface

The `start` command generates a `Solution` stub for you to fill in.

```go
type Solution interface {
    Part1(input Input) int
    Part2(input Input) int

    Answer1() int
    Answer2() int
}
```

### Reading Input

The `Input` interface is used to provide the puzzle input in a variety of basic forms so that you don't have to worry about parsing the input yourself.

You can override the separator used to split the input for a single day by calling the `SetSeparator` method on the `Input` instance before calling `Lines` or `Ints`.

```go
type Input interface {
    // SetSeparator is used to override the default separator
    SetSeparator(string)

    // Lines returns the input as a slice of strings, default separator is newline
    Lines() []string

    // Text returns the input as a single solid block of text
    Text() string

    // Ints returns the input as a slice of ints, default separator is newline
    Ints() []int

    // Int returns the input as a single number
    Int() int
}
```

### Solution Functions

Each AoC puzzle has two parts, so there are two functions you need to implement: `Part1` and `Part2`. Each returns an `int` since that's typically the answer that AoC expects.

Sometimes it's easier to calculate both parts in a single function, but that's up to you.

### Saving Answers

Once you've solved the puzzle, you can optionally save your answer by calling the `Answer1` and `Answer2` methods on the `Solution` interface. This essentially hard-codes your answer and provides an interface for the `advent` runner to compare your answer to the expected answer.

```go
func (s *Solution) Answer1() int {
    return 1234
}
```

### Debugging

You can use the `log.Debug(...)` function to pretty print all manner of inputs. These only show up when the `--debug` flag is used, making it a convenient way to show debuugging info selectively.

### Linting

I recommend the following tools: 

- [golangci-lint](https://github.com/golangci/golangci-lint)

### Marking Slow Solutions

You can mark a solution as slow by calling the `Slow` method on the `Solution` interface.

```go   
// TODO for me: come up with a better way to mark slow solutions
func (s *Solution) Slow() bool {
    return true
}
```