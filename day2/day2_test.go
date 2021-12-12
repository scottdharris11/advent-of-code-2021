package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 150, solvePart1(testValues))
	assert.Equal(t, 2027977, solvePart1(utils.ReadLines("day2", "day-2-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 900, solvePart2(testValues))
	assert.Equal(t, 1903644897, solvePart2(utils.ReadLines("day2", "day-2-input.txt")))
}
