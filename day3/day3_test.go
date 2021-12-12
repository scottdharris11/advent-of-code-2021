package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 198, solvePart1(testValues))
	assert.Equal(t, 3969000, solvePart1(utils.ReadLines("day3", "day-3-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 230, solvePart2(testValues))
	assert.Equal(t, 4267809, solvePart2(utils.ReadLines("day3", "day-3-input.txt")))
}
