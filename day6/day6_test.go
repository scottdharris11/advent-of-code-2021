package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []int{3, 4, 3, 1, 2}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 5934, solvePart1(testValues))
	assert.Equal(t, 362740, solvePart1(utils.ReadIntegersFromLine("day6", "day-6-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 26984457539, solvePart2(testValues))
	assert.Equal(t, 1644874076764, solvePart2(utils.ReadIntegersFromLine("day6", "day-6-input.txt")))
}
