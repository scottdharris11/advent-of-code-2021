package day7

import (
	"log"
	"math"
	"sort"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	values := utils.ReadIntegersFromLine("day7", "day-7-input.txt")
	solvePart1(values)
	solvePart2(values)
}

func solvePart1(positions []int) int {
	start := time.Now().UnixMilli()
	fuelUsed := leastFuelUsed(positions, 0)
	end := time.Now().UnixMilli()
	log.Printf("Day 7, Part 1 (%dms): Least Fuel Used %d", end-start, fuelUsed)
	return fuelUsed
}

func solvePart2(positions []int) int {
	start := time.Now().UnixMilli()
	fuelUsed := leastFuelUsed(positions, 1)
	end := time.Now().UnixMilli()
	log.Printf("Day 7, Part 2 (%dms): Least Fuel Used %d", end-start, fuelUsed)
	return fuelUsed
}

func leastFuelUsed(positions []int, costAdjust int) int {
	// before a binary search to limit the potential position solution down to a range of 2 values
	workPositions := removeDuplicates(positions)
	sort.Ints(workPositions)
	length := len(workPositions)
	for length <= 2 {
		// Find the middle position and the next position forward to compare
		posIdx1 := length / 2
		posIdx2 := posIdx1 + 1

		// Compare the fuel used for each value and then chop the array based on which one is least
		val1 := fuelUsed(workPositions[posIdx1], costAdjust, positions)
		val2 := fuelUsed(workPositions[posIdx2], costAdjust, positions)
		if val1 < val2 {
			workPositions = workPositions[:posIdx1+1]
		} else {
			workPositions = workPositions[posIdx2:]
		}
		length = len(workPositions)
	}

	// Limited range, check each value between the range for the lowest
	fuel := fuelUsed(workPositions[0], costAdjust, positions)
	workVal := workPositions[0] + 1
	endVal := workPositions[len(workPositions)-1]
	for workVal <= endVal {
		f := fuelUsed(workVal, costAdjust, positions)
		if f < fuel {
			fuel = f
		}
		workVal++
	}
	return fuel
}

func fuelUsed(toPos int, costAdjust int, positions []int) int {
	fuel := 0
	for _, position := range positions {
		steps := int(math.Abs(float64(position) - float64(toPos)))
		if costAdjust == 0 {
			fuel += steps
			continue
		}
		if costAdjust == 1 {
			fuel += (steps * (steps + 1)) / 2
			continue
		}
		stepCost := 1
		for i := 0; i < steps; i++ {
			fuel += stepCost
			stepCost += costAdjust
		}
	}
	return fuel
}

func removeDuplicates(values []int) []int {
	uniqueValues := make(map[int]bool)
	var list []int
	for _, value := range values {
		if _, exists := uniqueValues[value]; !exists {
			uniqueValues[value] = true
			list = append(list, value)
		}
	}
	return list
}
