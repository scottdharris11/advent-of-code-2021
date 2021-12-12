package day6

import (
	"log"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	values := utils.ReadIntegersFromLine("day6", "day-6-input.txt")
	solvePart1(values)
	solvePart2(values)
}

func solvePart1(values []int) int {
	fish := parseInput(values)
	start := time.Now().UnixMilli()
	fishCnt := simulateDays(80, &fish)
	end := time.Now().UnixMilli()
	log.Printf("Day 6, Part 1 (%dms): laternfish %d", end-start, fishCnt)
	return fishCnt
}

func solvePart2(values []int) int {
	fish := parseInput(values)
	start := time.Now().UnixMilli()
	fishCnt := simulateDays(256, &fish)
	end := time.Now().UnixMilli()
	log.Printf("Day 6, Part 2 (%dms): laternfish %d", end-start, fishCnt)
	return fishCnt
}

func parseInput(values []int) [9]int {
	fish := [9]int{}
	for _, value := range values {
		fish[value]++
	}
	return fish
}

func simulateDays(daysToSim int, fish *[9]int) int {
	totalFish := 0
	for i := 0; i < daysToSim; i++ {
		totalFish = 0
		spawnCnt := fish[0]
		for j := 1; j < 9; j++ {
			totalFish += fish[j]
			fish[j-1] = fish[j]
		}
		fish[6] += spawnCnt
		fish[8] = spawnCnt
		totalFish += spawnCnt * 2
	}
	return totalFish
}
