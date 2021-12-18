package day18

import (
	"advent-of-code-2021/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []string{
	"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
	"[[[5,[2,8]],4],[5,[[9,9],0]]]",
	"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
	"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
	"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
	"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
	"[[[[5,4],[7,7]],8],[[8,3],8]]",
	"[[9,3],[[9,9],[6,[4,9]]]]",
	"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
	"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 4140, solvePart1(testValues))
	assert.Equal(t, 4017, solvePart1(utils.ReadLines("day18", "day-18-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3993, solvePart2(testValues))
	assert.Equal(t, 4583, solvePart2(utils.ReadLines("day18", "day-18-input.txt")))
}

func TestNewPair(t *testing.T) {
	p := NewPair("[1,2]", nil)
	assert.Equal(t, 1, p.lLiteral)
	assert.Equal(t, 2, p.rLiteral)

	p = NewPair("[[1,2],[[3,4],5]]", nil)

	assert.NotNil(t, p.lPair)
	assert.NotNil(t, p.rPair)

	assert.Equal(t, p, p.lPair.parent)
	assert.Equal(t, 1, p.lPair.lLiteral)
	assert.Equal(t, 2, p.lPair.rLiteral)

	assert.Equal(t, p, p.rPair.parent)
	assert.NotNil(t, p.rPair.lPair)
	assert.Equal(t, 5, p.rPair.rLiteral)

	assert.Equal(t, p.rPair, p.rPair.lPair.parent)
	assert.Equal(t, 3, p.rPair.lPair.lLiteral)
	assert.Equal(t, 4, p.rPair.lPair.rLiteral)
}

func TestAdd(t *testing.T) {
	tests := []struct {
		value1   string
		value2   string
		expected string
	}{
		{"[1,1]", "[2,2]", "[[1,1],[2,2]]"},
		{"[[1,1],[2,2]]", "[3,3]", "[[[1,1],[2,2]],[3,3]]"},
		{"[[[1,1],[2,2]],[3,3]]", "[4,4]", "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", "[5,5]", "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", "[6,6]", "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
		{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
		{"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]", "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"},
		{"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]", "[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]", "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"},
		{"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]", "[7,[5,[[3,8],[1,4]]]]", "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]"},
		{"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]", "[[2,[2,2]],[8,[8,1]]]", "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]"},
		{"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]", "[2,9]", "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]"},
		{"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]", "[1,[[[9,3],9],[[9,0],[0,7]]]]", "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]"},
		{"[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]", "[[[5,[7,4]],7],1]", "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]"},
		{"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]", "[[[[4,2],2],6],[8,7]]", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
	}
	for _, tt := range tests {
		tt := tt

		t.Run("Scenario", func(t *testing.T) {
			p1 := NewPair(tt.value1, nil)
			p2 := NewPair(tt.value2, nil)
			result := p1.Add(p2)
			assert.Equal(t, tt.expected, result.String())
			assert.Equal(t, tt.value1, p1.String())
			assert.Equal(t, tt.value2, p2.String())
		})
	}
}

func TestAdd_Chain(t *testing.T) {
	var result *Pair
	result = NewPair(testValues[0], nil)
	for i, v := range testValues {
		if i == 0 {
			continue
		}
		result = result.Add(NewPair(v, nil))
	}
	assert.Equal(t, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", result.String())
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name     string
		pair     string
		expected int
	}{
		{"Just literals", "[1,2]", 7},
		{"Value 1", "[[1,2],[[3,4],5]]", 143},
		{"Value 2", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"Value 3", "[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"Value 4", "[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"Value 5", "[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"Value 6", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			p := NewPair(tt.pair, nil)
			assert.Equal(t, tt.expected, p.Magnitude())
		})
	}
}

func TestSplit(t *testing.T) {
	p := &Pair{lLiteral: 10, rLiteral: 1}
	p.split(p.lLiteral, true)
	assert.Equal(t, 1, p.rLiteral)
	assert.Equal(t, 0, p.lLiteral)
	assert.NotNil(t, p.lPair)
	assert.Equal(t, p, p.lPair.parent)
	assert.Equal(t, 5, p.lPair.lLiteral)
	assert.Equal(t, 5, p.lPair.rLiteral)

	p = &Pair{lLiteral: 11, rLiteral: 1}
	p.split(p.lLiteral, true)
	assert.Equal(t, 1, p.rLiteral)
	assert.Equal(t, 0, p.lLiteral)
	assert.NotNil(t, p.lPair)
	assert.Equal(t, p, p.lPair.parent)
	assert.Equal(t, 5, p.lPair.lLiteral)
	assert.Equal(t, 6, p.lPair.rLiteral)

	p = &Pair{lLiteral: 1, rLiteral: 13}
	p.split(p.rLiteral, false)
	assert.Equal(t, 1, p.lLiteral)
	assert.Equal(t, 0, p.rLiteral)
	assert.NotNil(t, p.rPair)
	assert.Equal(t, p, p.rPair.parent)
	assert.Equal(t, 6, p.rPair.lLiteral)
	assert.Equal(t, 7, p.rPair.rLiteral)
}

func TestExplode(t *testing.T) {
	var p *Pair
	var stack *ReductionStack

	p = NewPair("[[[[[9,8],1],2],3],4]", nil)
	stack = &ReductionStack{lastLiteralPair: nil}
	p.lPair.lPair.lPair.lPair.explode(stack)
	assert.Equal(t, "[[[[0,9],2],3],4]", p.String())

	p = NewPair("[7,[6,[5,[4,[3,2]]]]]", nil)
	stack = &ReductionStack{lastLiteralPair: p.rPair.rPair.rPair}
	p.rPair.rPair.rPair.rPair.explode(stack)
	assert.Equal(t, "[7,[6,[5,[7,0]]]]", p.String())

	p = NewPair("[[6,[5,[4,[3,2]]]],1]", nil)
	stack = &ReductionStack{lastLiteralPair: p.lPair.rPair.rPair}
	p.lPair.rPair.rPair.rPair.explode(stack)
	assert.Equal(t, "[[6,[5,[7,0]]],3]", p.String())

	p = NewPair("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", nil)
	stack = &ReductionStack{lastLiteralPair: p.lPair.rPair.rPair}
	p.lPair.rPair.rPair.rPair.explode(stack)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", p.String())

	p = NewPair("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", nil)
	stack = &ReductionStack{lastLiteralPair: p.rPair.rPair.rPair}
	p.rPair.rPair.rPair.rPair.explode(stack)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", p.String())
}
