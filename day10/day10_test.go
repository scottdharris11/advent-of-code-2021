package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 26397, solvePart1(testValues))
	assert.Equal(t, 392043, solvePart1(utils.ReadLines("day10", "day-10-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 288957, solvePart2(testValues))
	assert.Equal(t, 1605968119, solvePart2(utils.ReadLines("day10", "day-10-input.txt")))
}
