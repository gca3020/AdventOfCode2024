package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	list := flag.Bool("list", false, "List all available solvers")
	solverName := flag.String("solver", "", "Specify a solver to run")
	part := flag.Int("part", 0, "Specify part 1 or part 2 to run (0 for both)")

	flag.Parse()

	if *list {
		listSolvers()
		return
	}

	if *solverName == "" {
		fmt.Println("Please specify a solver name using -solver flag.")
		return
	}

	solver, exists := solvers[*solverName]
	if !exists {
		fmt.Println("Solver not found:", *solverName)
		return
	}

	data, err := os.ReadFile(*solverName + ".txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	solve(*part, solver, string(data))
}

func listSolvers() {
	// List all available solvers
	for name := range solvers {
		fmt.Println(name)
	}
}

func solve(part int, solver Solver, input string) {
	startTime := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(startTime))
	}()
	if part != 2 {
		solver.Init(input)
		fmt.Println("Part 1:", solver.Part1())
	}
	if part != 1 {
		solver.Init(input)
		fmt.Println("Part 2:", solver.Part2())
	}
}
