package day10

import (
	"log"
	"sort"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	lines := parseInput()
	start := time.Now().UnixMilli()
	score := illegalSyntaxScore(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 1 (%dms): %d", end-start, score)
}

func solvePart2() {
	lines := parseInput()
	start := time.Now().UnixMilli()
	score := autoCompleteScore(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 2 (%dms): %d", end-start, score)
}

func parseInput() []string {
	lines := utils.ReadLines("day10", "day-10-input.txt")
	/*lines = []string{
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
	}*/
	return lines
}

func illegalSyntaxScore(commands []string) int {
	score := 0
	for _, command := range commands {
		idx, _ := illegalCharIdx(command)
		if idx >= 0 {
			switch command[idx] {
			case ')':
				score += 3
			case ']':
				score += 57
			case '}':
				score += 1197
			case '>':
				score += 25137
			}
		}
	}
	return score
}

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
			switch stack[i] {
			case ')':
				s += 1
			case ']':
				s += 2
			case '}':
				s += 3
			case '>':
				s += 4
			}
		}
		scores = append(scores, s)
	}
	sort.Ints(scores)
	middle := len(scores) / 2
	return scores[middle]
}

func illegalCharIdx(cmd string) (int, []rune) {
	var stack []rune
	for i, c := range cmd {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')
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
