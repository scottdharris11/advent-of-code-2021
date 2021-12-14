package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{
	"NNCB",
	"",
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1588, solvePart1(testValues))
	assert.Equal(t, 2590, solvePart1(utils.ReadLines("day14", "day-14-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 2188189693529, solvePart2(testValues))
	assert.Equal(t, 2875665202438, solvePart2(utils.ReadLines("day14", "day-14-input.txt")))
}
