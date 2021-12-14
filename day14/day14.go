package day14

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day14", "day-14-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	pt := parseInput(lines)
	start := time.Now().UnixMilli()
	for i := 0; i < 10; i++ {
		pt.Step()
	}
	ans := pt.Score()
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 1 (%dms): Score = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	pt := parseInput(lines)
	start := time.Now().UnixMilli()
	for i := 0; i < 40; i++ {
		pt.Step()
	}
	ans := pt.Score()
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 2 (%dms): Score = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) *PolymerTemplate {
	polymer := lines[0]

	rules := make(map[string]rune)
	for i := 2; i < len(lines); i++ {
		pieces := strings.Split(lines[i], " -> ")
		rules[pieces[0]] = rune(pieces[1][0])
	}

	return NewPolymerTemplate(polymer, rules)
}

func NewPolymerTemplate(polymer string, rules map[string]rune) *PolymerTemplate {
	pt := PolymerTemplate{rules: rules}
	counts := make(map[rune]int)
	for _, r := range polymer {
		counts[r]++
	}
	combos := make(map[string]int)
	for i := 1; i < len(polymer); i++ {
		combos[polymer[i-1:i+1]]++
	}
	pt.combos = combos
	pt.counts = counts
	return &pt
}

type PolymerTemplate struct {
	combos map[string]int
	counts map[rune]int
	rules  map[string]rune
}

func (pt *PolymerTemplate) Step() {
	nCombos := make(map[string]int)
	for combo, count := range pt.combos {
		r := pt.rules[combo]
		pt.counts[r] += count

		nCombo1 := string([]rune{rune(combo[0]), r})
		nCombos[nCombo1] += count
		nCombo2 := string([]rune{r, rune(combo[1])})
		nCombos[nCombo2] += count
	}
	pt.combos = nCombos
}

func (pt *PolymerTemplate) Score() int {
	max := 0
	min := 999999999999999
	for _, c := range pt.counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	return max - min
}
