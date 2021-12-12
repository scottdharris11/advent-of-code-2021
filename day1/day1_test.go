package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 7, solvePart1(testValues))
	assert.Equal(t, 1553, solvePart1(utils.ReadIntegers("day1", "day-1-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 5, solvePart2(testValues))
	assert.Equal(t, 1597, solvePart2(utils.ReadIntegers("day1", "day-1-input.txt")))
}
