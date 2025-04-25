package main

type Solver interface {
	Init(input string)
	Part1() int
	Part2() int
}

var solvers map[string]Solver

func RegisterSolver(name string, solver Solver) {
	if solvers == nil {
		solvers = make(map[string]Solver)
	}
	solvers[name] = solver
}
