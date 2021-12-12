package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1656, solvePart1(testValues))
	assert.Equal(t, 1667, solvePart1(utils.ReadLines("day11", "day-11-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 195, solvePart2(testValues))
	assert.Equal(t, 488, solvePart2(utils.ReadLines("day11", "day-11-input.txt")))
}
