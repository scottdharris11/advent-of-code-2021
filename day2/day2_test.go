package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 150, solvePart1(testValues))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 900, solvePart2(testValues))
}
