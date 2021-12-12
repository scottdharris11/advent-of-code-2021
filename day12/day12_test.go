package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues1 = []string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}
var testValues2 = []string{"dc-end", "HN-start", "start-kj", "dc-start", "dc-HN", "LN-dc", "HN-end", "kj-sa", "kj-HN", "kj-dc"}
var testValues3 = []string{"fs-end", "he-DX", "fs-he", "start-DX", "pj-DX", "end-zg", "zg-sl", "zg-pj", "pj-he", "RW-he", "fs-DX", "pj-RW", "zg-RW", "start-pj", "he-WI", "zg-he", "pj-fs", "start-RW"}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 10, solvePart1(testValues1))
	assert.Equal(t, 19, solvePart1(testValues2))
	assert.Equal(t, 226, solvePart1(testValues3))
	assert.Equal(t, 4720, solvePart1(utils.ReadLines("day12", "day-12-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 36, solvePart2(testValues1))
	assert.Equal(t, 103, solvePart2(testValues2))
	assert.Equal(t, 3509, solvePart2(testValues3))
	assert.Equal(t, 147848, solvePart2(utils.ReadLines("day12", "day-12-input.txt")))
}
