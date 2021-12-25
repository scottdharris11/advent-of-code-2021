package main

import (
	"advent-of-code-2021/day1"
	"advent-of-code-2021/day10"
	"advent-of-code-2021/day11"
	"advent-of-code-2021/day12"
	"advent-of-code-2021/day13"
	"advent-of-code-2021/day14"
	"advent-of-code-2021/day15"
	"advent-of-code-2021/day16"
	"advent-of-code-2021/day17"
	"advent-of-code-2021/day18"
	"advent-of-code-2021/day19"
	"advent-of-code-2021/day2"
	"advent-of-code-2021/day20"
	"advent-of-code-2021/day21"
	"advent-of-code-2021/day22"
	"advent-of-code-2021/day23"
	"advent-of-code-2021/day24"
	"advent-of-code-2021/day25"
	"advent-of-code-2021/day3"
	"advent-of-code-2021/day4"
	"advent-of-code-2021/day5"
	"advent-of-code-2021/day6"
	"advent-of-code-2021/day7"
	"advent-of-code-2021/day8"
	"advent-of-code-2021/day9"
)

type Solver interface {
	Solve()
}

func main() {
	solvers := []Solver{
		day1.Puzzle{}, day2.Puzzle{}, day3.Puzzle{}, day4.Puzzle{}, day5.Puzzle{},
		day6.Puzzle{}, day7.Puzzle{}, day8.Puzzle{}, day9.Puzzle{}, day10.Puzzle{},
		day11.Puzzle{}, day12.Puzzle{}, day13.Puzzle{}, day14.Puzzle{}, day15.Puzzle{},
		day16.Puzzle{}, day17.Puzzle{}, day18.Puzzle{}, day19.Puzzle{}, day20.Puzzle{},
		day21.Puzzle{}, day22.Puzzle{}, day23.Puzzle{}, day24.Puzzle{}, day25.Puzzle{},
	}
	for _, solver := range solvers {
		solver.Solve()
	}
}
