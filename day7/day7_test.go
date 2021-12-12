package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 37, solvePart1(testValues))
	assert.Equal(t, 328187, solvePart1(utils.ReadIntegersFromLine("day7", "day-7-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 168, solvePart2(testValues))
	assert.Equal(t, 91257582, solvePart2(utils.ReadIntegersFromLine("day7", "day-7-input.txt")))
}
