package day6

import (
	"log"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	fish := parseInput()
	fishCnt := simulateDays(80, &fish)
	log.Printf("Day 6, Part 1: laternfish %d", fishCnt)
}

func solvePart2() {
	fish := parseInput()
	fishCnt := simulateDays(256, &fish)
	log.Printf("Day 6, Part 2: laternfish %d", fishCnt)
}

func parseInput() [9]int {
	values := utils.ReadIntegersFromLine("day6", "day-6-input.txt")
	//values = []int{3,4,3,1,2}

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
