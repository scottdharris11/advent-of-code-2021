package day20

import (
	"advent-of-code-2021/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 35, solvePart1(utils.ReadLines("day20", "day-20-test.txt")))
	assert.Equal(t, 5203, solvePart1(utils.ReadLines("day20", "day-20-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3351, solvePart2(utils.ReadLines("day20", "day-20-test.txt")))
	assert.Equal(t, 18806, solvePart2(utils.ReadLines("day20", "day-20-input.txt")))
}

func TestParseInput(t *testing.T) {
	ip, img := parseInput(utils.ReadLines("day20", "day-20-test.txt"))
	assert.Equal(t, 512, len(ip.enhanceAlgo))
	assert.Equal(t, 5, img.height)
	assert.Equal(t, 5, img.width)
	assert.Equal(t, "#..#.\n#....\n##..#\n..#..\n..###\n", img.String())
}

func TestPixelAtIndex(t *testing.T) {
	img := &Image{
		width:  3,
		height: 3,
		pixels: [][]rune{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		},
		expanseBit: '?',
	}
	assert.Equal(t, '?', img.pixelAtIndex(-1, 0))
	assert.Equal(t, '?', img.pixelAtIndex(0, -1))
	assert.Equal(t, '?', img.pixelAtIndex(4, 0))
	assert.Equal(t, '?', img.pixelAtIndex(0, 4))
	assert.Equal(t, '1', img.pixelAtIndex(0, 0))
	assert.Equal(t, '5', img.pixelAtIndex(1, 1))
	assert.Equal(t, '9', img.pixelAtIndex(2, 2))
}

func TestPixelArrayToNumber(t *testing.T) {
	pixels := [3][3]rune{
		{'.', '.', '.'},
		{'#', '.', '.'},
		{'.', '#', '.'},
	}
	assert.Equal(t, 34, pixelArrayToNumber(pixels))

	pixels = [3][3]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}
	assert.Equal(t, 0, pixelArrayToNumber(pixels))

	pixels = [3][3]rune{
		{'#', '#', '#'},
		{'#', '#', '#'},
		{'#', '#', '#'},
	}
	assert.Equal(t, 511, pixelArrayToNumber(pixels))
}
