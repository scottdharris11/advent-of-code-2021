package day18

import (
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day18", "day-18-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	pairs := parseInput(lines)
	start := time.Now().UnixMilli()
	result := addPairs(pairs)
	ans := result.Magnitude()
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 1 (%dms): Final Magnitude = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	pairs := parseInput(lines)
	start := time.Now().UnixMilli()
	ans := 0
	for _, p1 := range pairs {
		for _, p2 := range pairs {
			if p1 == p2 {
				continue
			}
			mag := p1.Add(p2).Magnitude()
			if mag > ans {
				ans = mag
			}
			mag = p2.Add(p1).Magnitude()
			if mag > ans {
				ans = mag
			}
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 2 (%dms): Largest Magnitude = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) []*Pair {
	var pairs []*Pair
	for _, line := range lines {
		pairs = append(pairs, NewPair(line, nil))
	}
	return pairs
}

func addPairs(pairs []*Pair) *Pair {
	result := pairs[0]
	for i, pair := range pairs {
		if i == 0 {
			continue
		}
		result = result.Add(pair)
	}
	return result
}

func NewPair(s string, parent *Pair) *Pair {
	p := &Pair{parent: parent}

	// find right and left sides
	work := s[1:findEndBraceIdx(s, 0)]
	splitIdx := findSplitIdx(work)
	lSide := work[0:splitIdx]
	rSide := work[splitIdx+1:]

	// expand left side values
	if lSide[0] == '[' {
		p.lPair = NewPair(lSide, p)
	} else {
		p.lLiteral = int(lSide[0] - '0')
	}

	// expand right side values
	if rSide[0] == '[' {
		p.rPair = NewPair(rSide, p)
	} else {
		p.rLiteral = int(rSide[0] - '0')
	}

	return p
}

func findEndBraceIdx(s string, fromIdx int) int {
	skipCnt := 0
	l := len(s)
	for i := fromIdx + 1; i < l; i++ {
		switch s[i] {
		case '[':
			skipCnt++
		case ']':
			if skipCnt == 0 {
				return i
			}
			skipCnt--
		}
	}
	return -1
}

func findSplitIdx(s string) int {
	skipCnt := 0
	for i, c := range s {
		switch c {
		case '[':
			skipCnt++
		case ']':
			skipCnt--
		case ',':
			if skipCnt == 0 {
				return i
			}
		}
	}
	return -1
}

type Pair struct {
	parent   *Pair
	lLiteral int
	lPair    *Pair
	rLiteral int
	rPair    *Pair
}

func (p *Pair) Magnitude() int {
	mag := 0
	if p.lPair != nil {
		mag += p.lPair.Magnitude() * 3
	}
	mag += p.lLiteral * 3
	if p.rPair != nil {
		mag += p.rPair.Magnitude() * 2
	}
	mag += p.rLiteral * 2
	return mag
}

func (p *Pair) Add(a *Pair) *Pair {
	nPair := &Pair{}
	nPair.lPair = p.Copy(nPair)
	nPair.rPair = a.Copy(nPair)
	nPair.reduce()
	return nPair
}

func (p *Pair) Copy(nParent *Pair) *Pair {
	cPair := &Pair{}
	cPair.parent = nParent
	cPair.lLiteral = p.lLiteral
	if p.lPair != nil {
		cPair.lPair = p.lPair.Copy(cPair)
	}
	cPair.rLiteral = p.rLiteral
	if p.rPair != nil {
		cPair.rPair = p.rPair.Copy(cPair)
	}
	return cPair
}

func (p *Pair) AddToLeftMostLiteral(v int) {
	if p.lPair == nil {
		p.lLiteral += v
		return
	}
	if p.rPair == nil {
		p.rLiteral += v
	}
}

func (p *Pair) AddToRightMostLiteral(v int) {
	if p.rPair == nil {
		p.rLiteral += v
		return
	}
	if p.lPair == nil {
		p.lLiteral += v
	}
}

func (p *Pair) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	if p.lPair != nil {
		sb.WriteString(p.lPair.String())
	} else {
		sb.WriteString(strconv.Itoa(p.lLiteral))
	}
	sb.WriteRune(',')
	if p.rPair != nil {
		sb.WriteString(p.rPair.String())
	} else {
		sb.WriteString(strconv.Itoa(p.rLiteral))
	}
	sb.WriteRune(']')
	return sb.String()
}

func (p *Pair) reduce() {
	for {
		// log.Printf("Before Explode ONLY Reduction: %s", p)
		rStack := &ReductionStack{depth: -1}
		if p.processReduction(rStack, false) {
			continue
		}

		// log.Printf("Before Split Reduction: %s", p)
		rStack = &ReductionStack{depth: -1}
		if !p.processReduction(rStack, true) {
			break
		}
	}
}

func (p *Pair) processReduction(stack *ReductionStack, doSplit bool) bool {
	// increment depth, if four, process explosion scenario
	stack.depth++
	if stack.depth == 4 {
		p.explode(stack)
		return true
	}

	// check left side
	if p.lPair != nil {
		if p.lPair.processReduction(stack, doSplit) {
			return true
		}
	} else {
		if doSplit && p.lLiteral > 9 {
			p.split(p.lLiteral, true)
			return true
		}
		stack.lastLiteralPair = p
	}

	// check right side
	if p.rPair != nil {
		if p.rPair.processReduction(stack, doSplit) {
			return true
		}
	} else {
		if doSplit && p.rLiteral > 9 {
			p.split(p.rLiteral, false)
			return true
		}
		stack.lastLiteralPair = p
	}

	// reduce depth and return with no reduction
	stack.depth--
	return false
}

func (p *Pair) explode(stack *ReductionStack) {
	// if literal value to left, add the left value to it
	if stack.lastLiteralPair != nil {
		stack.lastLiteralPair.AddToRightMostLiteral(p.lLiteral)
	}

	// if literal value to the right, add the right value to it
	nextLiteralPair := p.findLiteralToRight()
	if nextLiteralPair != nil {
		nextLiteralPair.AddToLeftMostLiteral(p.rLiteral)
	}

	// zero out pair value in parent
	if p.parent.lPair == p {
		p.parent.lPair = nil
		p.parent.lLiteral = 0
		return
	}
	p.parent.rPair = nil
	p.parent.rLiteral = 0
}

func (p *Pair) findLiteralToRight() *Pair {
	// if no parent, no next to be found
	if p.parent == nil {
		return nil
	}

	// if current pair is the right pair, recurse to the parent
	if p.parent.rPair == p {
		return p.parent.findLiteralToRight()
	}

	// if parent right value is a literal, then we found it
	if p.parent.rPair == nil {
		return p.parent
	}

	// if parent right value is a pair, search from there to find next
	return p.parent.rPair.findNextLiteralPair()
}

func (p *Pair) findNextLiteralPair() *Pair {
	if p.lPair == nil {
		return p
	}
	return p.lPair.findNextLiteralPair()
}

func (p *Pair) split(value int, left bool) {
	lValue := value / 2
	rValue := value / 2
	if rValue*2 < value {
		rValue++
	}

	nPair := &Pair{parent: p, lLiteral: lValue, rLiteral: rValue}
	if left {
		p.lPair = nPair
		p.lLiteral = 0
	} else {
		p.rPair = nPair
		p.rLiteral = 0
	}
}

type ReductionStack struct {
	depth           int
	lastLiteralPair *Pair
}
