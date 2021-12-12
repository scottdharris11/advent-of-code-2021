package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 7, solvePart1(testValues))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 5, solvePart2(testValues))
}
