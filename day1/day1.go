package day1

import (
	"log"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	//values := []int{ 199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	values := utils.ReadIntegers("day1", "day-1-input.txt")
	start := time.Now().UnixMilli()
	increased := numTimesIncreased(values)
	end := time.Now().UnixMilli()
	log.Printf("Day 1, Part 1 (%dms): Increased Count %d", end-start, increased)
}

func solvePart2() {
	//values := []int{ 199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	values := utils.ReadIntegers("day1", "day-1-input.txt")

	start := time.Now().UnixMilli()
	valCount := len(values)
	var sections []int
	for i := 2; i < valCount; i++ {
		sections = append(sections, values[i]+values[i-1]+values[i-2])
	}
	increased := numTimesIncreased(sections)
	end := time.Now().UnixMilli()
	log.Printf("Day 1, Part 2 (%dms): Increased Count %d", end-start, increased)
}

func numTimesIncreased(values []int) int {
	valCount := len(values)
	increased := 0
	for i := 1; i < valCount; i++ {
		if values[i] > values[i-1] {
			increased++
		}
	}
	return increased
}
