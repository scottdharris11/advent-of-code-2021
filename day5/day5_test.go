package day5

import (
	"advent-of-code-2021/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []string{
	"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1",
	"7,0 -> 7,4", "6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4",
	"0,0 -> 8,8", "5,5 -> 8,2",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 5, solvePart1(testValues))
	assert.Equal(t, 4728, solvePart1(utils.ReadLines("day5", "day-5-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 12, solvePart2(testValues))
	assert.Equal(t, 17717, solvePart2(utils.ReadLines("day5", "day-5-input.txt")))
}
