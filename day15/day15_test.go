package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 40, solvePart1(testValues, true))
	assert.Equal(t, 707, solvePart1(utils.ReadLines("day15", "day-15-input.txt"), true))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 315, solvePart2(testValues, true))
	assert.Equal(t, 2942, solvePart2(utils.ReadLines("day15", "day-15-input.txt"), true))
}
