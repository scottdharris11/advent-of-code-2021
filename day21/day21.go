package day21

import (
	"log"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day21", "day-21-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 1 (%dms): TBD = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 2 (%dms): TBD = %d", end-start, ans)
	return ans
}
