package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 15, solvePart1(testValues))
	assert.Equal(t, 631, solvePart1(utils.ReadLines("day9", "day-9-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 1134, solvePart2(testValues))
	assert.Equal(t, 821560, solvePart2(utils.ReadLines("day9", "day-9-input.txt")))
}
