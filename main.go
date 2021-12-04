package main

import (
	"advent-of-code-2021/day1"
	"advent-of-code-2021/day2"
	"advent-of-code-2021/day3"
)

type Solver interface {
	Solve()
}

func main() {
	solvers := []Solver{
		day1.Puzzle{}, day2.Puzzle{}, day3.Puzzle{},
	}
	for _, solver := range solvers {
		solver.Solve()
	}
}
