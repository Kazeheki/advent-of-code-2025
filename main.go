package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"

	"github.com/Kazeheki/advent-of-code-2025/solve"
	"github.com/Kazeheki/advent-of-code-2025/util"
)

type exitCode int

const (
	invalidArguments exitCode = iota + 1
)

func main() {

	args := os.Args[1:]

	days := determineDaysToSolve(args)

	runtime.GOMAXPROCS(4)
	doneCh := make(chan struct{})

	for _, day := range days {
		go executeDay(day, doneCh)
	}

	for range days {
		<-doneCh
	}
}

func executeDay(day int, doneCh chan<- struct{}) {
	defer func() { doneCh <- struct{}{} }()

	input, err := util.ReadFullDayInput(day)
	if err != nil {
		slog.Warn("issue on reading input", "day", day, "error", err)
		return
	}

	solver := createSolver(day)

	solver.SetInput(input)

	err = solveWithSolver(solver, solve.PART_1)
	if err != nil {
		slog.Warn("part 1 has failed", "day", day, "error", err)
		return
	}

	err = solveWithSolver(solver, solve.PART_2)
	if err != nil {
		slog.Warn("part 2 has failed", "day", day, "error", err)
		return
	}
}

func determineDaysToSolve(args []string) []int {
	if err := validateArgs(args); err != nil {
		fmt.Printf("the arguments cannot be used: %v\n", err)
		os.Exit(int(invalidArguments))
	}

	days := make([]int, 0, 12)
	if len(args) == 0 {
		// we want to run through all days if there is no argument given.
		for day := range 11 {
			days = append(days, day+1)
		}
	} else {
		for _, dayStr := range args {
			// ignoring error, the validation made sure this conversion to int is possible for all arguments.
			day, _ := strconv.Atoi(dayStr)
			days = append(days, day)
		}
	}
	return days
}

func createSolver(day int) solve.CanSolve {
	var solver solve.CanSolve
	switch day {
	case 1:
		solver = &solve.Day01Solver{}
	case 2:
		solver = &solve.Day02Solver{}
	case 3:
		solver = &solve.Day03Solver{}
	case 4:
		solver = &solve.Day04Solver{}
	case 5:
		solver = &solve.Day05Solver{}
	case 6:
		solver = &solve.Day06Solver{}
	case 7:
		solver = &solve.Day07Solver{}
	case 8:
		solver = &solve.Day08Solver{}
	case 9:
		solver = &solve.Day09Solver{}
	case 10:
		solver = &solve.Day10Solver{}
	case 11:
		solver = &solve.Day11Solver{}
	case 12:
		solver = &solve.Day12Solver{}
	default:
		panic("how did we get here?? we cannot be here, we validated!")
	}
	return solver
}

func solveWithSolver(solver solve.CanSolve, part solve.PuzzlePart) error {
	result, err := solver.Solve(part)
	if err != nil {
		return fmt.Errorf("solving part %v did not succeed: %w", part, err)
	}

	util.WriteResult(solver.GetDay(), part, result)
	return nil
}

func validateArgs(args []string) error {
	if len(args) == 0 {
		return nil
	}

	for _, arg := range args {
		asInt, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("argument must be a number, %v is no number", arg)
		}

		if asInt < 1 || asInt > 12 {
			return fmt.Errorf("argument must be between 1 and 12 (given: %d)", asInt)
		}
	}

	return nil
}
