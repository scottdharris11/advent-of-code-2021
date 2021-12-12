package day8

import (
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day8", "day-8-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	entries := parseInput(lines)
	start := time.Now().UnixMilli()
	counts := countDigits(entries)
	sum := sumCounts(counts, []int{1, 4, 7, 8})
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 1 (%dms): %d", end-start, sum)
	return sum
}

func solvePart2(lines []string) int {
	entries := parseInput(lines)
	start := time.Now().UnixMilli()
	ans := 0
	for _, entry := range entries {
		ans += determineValue(entry)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 2 (%dms): %d", end-start, ans)
	return ans
}

func parseInput(lines []string) []Entry {
	var entries []Entry
	for _, line := range lines {
		pieces := strings.Split(line, " | ")
		entries = append(entries, Entry{
			patterns: strings.Split(pieces[0], " "),
			values:   strings.Split(pieces[1], " "),
		})
	}
	return entries
}

type Entry struct {
	patterns []string
	values   []string
}

func sumCounts(digitCounts map[int]int, digits []int) int {
	sum := 0
	for _, digit := range digits {
		sum += digitCounts[digit]
	}
	return sum
}

func countDigits(entries []Entry) map[int]int {
	digitCounts := make(map[int]int)
	for _, entry := range entries {
		for _, value := range entry.values {
			length := len(value)
			number := 0
			switch length {
			case 2:
				number = 1
			case 3:
				number = 7
			case 4:
				number = 4
			case 7:
				number = 8
			}
			digitCounts[number]++
		}
	}
	return digitCounts
}

func determineValue(e Entry) int {
	// eliminate potential values based on positioning of known numbers
	values := [][]string{
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g"},
	}
	for _, pattern := range e.patterns {
		length := len(pattern)
		switch length {
		case 2:
			values = eliminate(pattern, []int{2, 5}, values) //1
		case 3:
			values = eliminate(pattern, []int{0, 2, 5}, values) //7
		case 4:
			values = eliminate(pattern, []int{1, 2, 3, 5}, values) //4
		}
	}

	// build patterns based on remaining possibilities and test
	// them against the different patterns to determine which
	// configuration is the right one (generates all good numbers)
	possible := buildPossiblePatterns(values)
	signalConfig := ""
	for _, p := range possible {
		allGood := true
		for _, s := range e.patterns {
			if toDigit(p, s) < 0 {
				allGood = false
				break
			}
		}
		if allGood {
			signalConfig = p
			break
		}
	}

	// Build number from digit
	sNum := ""
	for _, v := range e.values {
		sNum += strconv.Itoa(toDigit(signalConfig, v))
	}
	num, err := strconv.Atoi(sNum)
	if err != nil {
		log.Fatalln("Bad number conversion", err)
	}
	return num
}

func eliminate(s string, positions []int, values [][]string) [][]string {
	for i := 0; i < len(values); i++ {
		elimIn := true
		for _, position := range positions {
			if i == position {
				elimIn = false
				break
			}
		}
		var nValues []string
		for _, v := range values[i] {
			contains := strings.Contains(s, v)
			if !((!elimIn && !contains) || (elimIn && contains)) {
				nValues = append(nValues, v)
			}
		}
		values[i] = nValues
	}
	return values
}

func buildPossiblePatterns(values [][]string) []string {
	var possible []string
	for i := 0; i < len(values); i++ {
		// if first entry...just add the values
		if i == 0 {
			possible = append(possible, values[i]...)
			continue
		}

		// if there is only one value possible for position, append it to all existing values
		l := len(values[i])
		if l == 1 {
			pLen := len(possible)
			for j := 0; j < pLen; j++ {
				possible[j] += values[i][0]
			}
			continue
		}

		// there are multiple values for the position, so we need to create a copy of all previous
		// values for each option above one and then append the different possibilities here
		c := make([]string, len(possible))
		copy(c, possible)
		for j := 0; j < l; j++ {
			if j == 0 {
				pLen := len(possible)
				for k := 0; k < pLen; k++ {
					possible[k] += values[i][0]
				}
				continue
			}
			cLen := len(c)
			for k := 0; k < cLen; k++ {
				possible = append(possible, c[k]+values[i][j])
			}
		}

		// remove entries where same value is listed twice
		l = len(possible)
		var newP []string
		for j := 0; j < l; j++ {
			good := true
			for _, c := range possible[j] {
				if strings.Count(possible[j], string(c)) > 1 {
					good = false
					break
				}
			}
			if good {
				newP = append(newP, possible[j])
			}
		}
		possible = newP
	}

	return possible
}

var DIGITS = []string{
	"1110111", //0
	"0010010", //1
	"1011101", //2
	"1011011", //3
	"0111010", //4
	"1101011", //5
	"1101111", //6
	"1010010", //7
	"1111111", //8
	"1111011", //9
}

func toDigit(p string, s string) int {
	nStr := ""
	for _, c := range p {
		if strings.ContainsRune(s, c) {
			nStr += "1"
		} else {
			nStr += "0"
		}
	}

	n := -1
	for i, d := range DIGITS {
		if d == nStr {
			n = i
			break
		}
	}
	return n
}
