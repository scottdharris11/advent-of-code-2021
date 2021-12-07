package day7

import (
	"log"
	"math"
	"sort"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	positions := parseInput()
	fuelUsed := leastFuelUsed(positions, 0)
	log.Printf("Day 7, Part 1: Least Fuel Used %d", fuelUsed)
}

func solvePart2() {
	positions := parseInput()
	fuelUsed := leastFuelUsed(positions, 1)
	log.Printf("Day 7, Part 2: Least Fuel Used %d", fuelUsed)
}

func parseInput() []int {
	values := utils.ReadIntegersFromLine("day7", "day-7-input.txt")
	//values = []int{16,1,2,0,4,2,7,1,2,14}
	return values
}

func leastFuelUsed(positions []int, costAdjust int) int {
	// before a binary search to limit the potential position solution down to a range of 2 values
	workPositions := positions[:]
	sort.Ints(workPositions)
	length := len(workPositions)
	for length != 2 {
		// Find the middle position and the next position forward with diff value
		midIdx := int(length/2) - 1
		posIdx1 := midIdx
		posIdx2 := midIdx + 1
		for ; posIdx2 < length; posIdx2++ {
			if workPositions[posIdx1] != workPositions[posIdx2] {
				break
			}
		}

		// If we run to end and still have the same value, look backwards for diff value
		if posIdx2 == length {
			posIdx2 = posIdx1
			posIdx1 = posIdx2 - 1
			for ; posIdx1 >= 0; posIdx1-- {
				if workPositions[posIdx1] != workPositions[posIdx2] {
					break
				}
				posIdx2--
			}
		}

		// All values must be equal, we have our answer
		if posIdx1 < 0 {
			return fuelUsed(workPositions[posIdx2], costAdjust, positions)
		}

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

	// Limited range to two values, check each value between the range for the lowest
	fuel := fuelUsed(workPositions[0], costAdjust, positions)
	workVal := workPositions[0] + 1
	endVal := workPositions[1]
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
		stepCost := 1
		for i := 0; i < steps; i++ {
			fuel += stepCost
			stepCost += costAdjust
		}
	}
	return fuel
}
