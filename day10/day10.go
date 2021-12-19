package day10

import (
	"log"
	"sort"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day10", "day-10-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	score := illegalSyntaxScore(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 1 (%dms): %d", end-start, score)
	return score
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	score := autoCompleteScore(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 2 (%dms): %d", end-start, score)
	return score
}

var illegalCharScores = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

func illegalSyntaxScore(commands []string) int {
	score := 0
	for _, command := range commands {
		idx, _ := illegalCharIdx(command)
		if idx >= 0 {
			score += illegalCharScores[rune(command[idx])]
		}
	}
	return score
}

var autoCompleteCharScores = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

func autoCompleteScore(commands []string) int {
	var scores []int
	for _, command := range commands {
		idx, stack := illegalCharIdx(command)
		if idx >= 0 {
			continue
		}
		l := len(stack)
		s := 0
		for i := l - 1; i >= 0; i-- {
			s *= 5
			s += autoCompleteCharScores[stack[i]]
		}
		scores = append(scores, s)
	}
	sort.Ints(scores)
	middle := len(scores) / 2
	return scores[middle]
}

var matchingChar = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

func illegalCharIdx(cmd string) (int, []rune) {
	var stack []rune
	for i, c := range cmd {
		switch c {
		case '(', '[', '{', '<':
			stack = append(stack, matchingChar[c])
		default:
			n := len(stack) - 1
			validClose := stack[n]
			stack = stack[:n]
			if c != validClose {
				return i, nil
			}
		}
	}
	return -1, stack
}
