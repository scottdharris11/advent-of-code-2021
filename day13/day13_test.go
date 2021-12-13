package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{
	"6,10",
	"0,14",
	"9,10",
	"0,3",
	"10,4",
	"4,11",
	"6,0",
	"6,12",
	"4,1",
	"0,13",
	"10,12",
	"3,4",
	"3,0",
	"8,4",
	"1,10",
	"2,14",
	"8,10",
	"9,0",
	"",
	"fold along y=7",
	"fold along x=5",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 17, solvePart1(testValues))
	assert.Equal(t, 618, solvePart1(utils.ReadLines("day13", "day-13-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	expected := "#####\n" +
		"#...#\n" +
		"#...#\n" +
		"#...#\n" +
		"#####\n" +
		".....\n" +
		".....\n"
	assert.Equal(t, expected, solvePart2(testValues))

	expected = ".##..#....###..####.#..#.####.#..#.#..#.\n" +
		"#..#.#....#..#.#....#.#..#....#.#..#..#.\n" +
		"#..#.#....#..#.###..##...###..##...#..#.\n" +
		"####.#....###..#....#.#..#....#.#..#..#.\n" +
		"#..#.#....#.#..#....#.#..#....#.#..#..#.\n" +
		"#..#.####.#..#.####.#..#.#....#..#..##..\n"
	assert.Equal(t, expected, solvePart2(utils.ReadLines("day13", "day-13-input.txt")))
}
