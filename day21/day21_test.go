package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 0, solvePart1(utils.ReadLines("day21", "day-21-test.txt")))
	assert.Equal(t, 0, solvePart1(utils.ReadLines("day21", "day-21-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day21", "day-21-test.txt")))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day21", "day-21-input.txt")))
}
