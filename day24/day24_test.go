package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 91897399498995, solvePart1(utils.ReadLines("day24", "day-24-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 51121176121391, solvePart2(utils.ReadLines("day24", "day-24-input.txt")))
}

func TestArithmeticLogicUnit_RunProgram(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		inputs       []int
		expected     [4]int
	}{
		{"negate number", []string{
			"inp x", "mul x -1",
		}, []int{20}, [4]int{0, -20, 0, 0}},
		{"three times larger - yes", []string{
			"inp z", "inp x", "mul z 3", "eql z x",
		}, []int{3, 9}, [4]int{0, 9, 0, 1}},
		{"three times larger - no", []string{
			"inp z", "inp x", "mul z 3", "eql z x",
		}, []int{4, 9}, [4]int{0, 9, 0, 0}},
		{"binary conversion", []string{
			"inp w", "add z w", "mod z 2", "div w 2", "add y w", "mod y 2", "div w 2", "add x w", "mod x 2", "div w 2", "mod w 2",
		}, []int{5}, [4]int{0, 1, 0, 1}},
		{"binary conversion", []string{
			"inp w", "add z w", "mod z 2", "div w 2", "add y w", "mod y 2", "div w 2", "add x w", "mod x 2", "div w 2", "mod w 2",
		}, []int{9}, [4]int{1, 0, 0, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			alu := NewArithmeticLogicUnit(0, 0, 0, 0)
			alu.RunProgram(tt.instructions, tt.inputs...)
			assert.Equal(t, tt.expected[0], *alu.w)
			assert.Equal(t, tt.expected[1], *alu.x)
			assert.Equal(t, tt.expected[2], *alu.y)
			assert.Equal(t, tt.expected[3], *alu.z)
		})
	}
}

func TestArithmeticLogicUnit_ApplyInstruction(t *testing.T) {
	tests := []struct {
		name        string
		alu         *ArithmeticLogicUnit
		instruction string
		value       int
		expected    [4]int
	}{
		{"inp", NewArithmeticLogicUnit(1, 2, 3, 4), "inp x", 20, [4]int{1, 20, 3, 4}},
		{"add", NewArithmeticLogicUnit(1, 2, 3, 4), "add y z", 0, [4]int{1, 2, 7, 4}},
		{"add lit", NewArithmeticLogicUnit(1, 2, 3, 4), "add w 10", 0, [4]int{11, 2, 3, 4}},
		{"mul", NewArithmeticLogicUnit(1, 2, 3, 4), "mul z x", 0, [4]int{1, 2, 3, 8}},
		{"mul lit", NewArithmeticLogicUnit(1, 2, 3, 4), "mul x 3", 0, [4]int{1, 6, 3, 4}},
		{"div", NewArithmeticLogicUnit(1, 2, 3, 4), "div x y", 0, [4]int{1, 0, 3, 4}},
		{"div lit", NewArithmeticLogicUnit(1, 2, 3, 4), "div z 3", 0, [4]int{1, 2, 3, 1}},
		{"mod", NewArithmeticLogicUnit(1, 2, 3, 4), "mod z y", 0, [4]int{1, 2, 3, 1}},
		{"mod lit", NewArithmeticLogicUnit(1, 2, 3, 4), "mod z 3", 0, [4]int{1, 2, 3, 1}},
		{"eql yes", NewArithmeticLogicUnit(1, 2, 4, 4), "eql z y", 0, [4]int{1, 2, 4, 1}},
		{"eql no", NewArithmeticLogicUnit(1, 2, 3, 4), "eql y z", 0, [4]int{1, 2, 0, 4}},
		{"eql lit no", NewArithmeticLogicUnit(1, 2, 3, 4), "eql x 2", 0, [4]int{1, 1, 3, 4}},
		{"eql lit no", NewArithmeticLogicUnit(1, 2, 3, 4), "eql w 3", 0, [4]int{0, 2, 3, 4}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.alu.applyInstruction(tt.instruction, tt.value)
			assert.Equal(t, tt.expected[0], *tt.alu.w)
			assert.Equal(t, tt.expected[1], *tt.alu.x)
			assert.Equal(t, tt.expected[2], *tt.alu.y)
			assert.Equal(t, tt.expected[3], *tt.alu.z)
		})
	}

}

func TestArithmeticLogicUnit_Assign(t *testing.T) {
	a := 1
	b := 3
	ArithmeticLogicUnit{}.assign(&a, b)
	assert.Equal(t, 3, a)
	assert.Equal(t, 3, b)
}

func TestArithmeticLogicUnit_Add(t *testing.T) {
	a := 1
	b := 3
	ArithmeticLogicUnit{}.add(&a, &b)
	assert.Equal(t, 4, a)
	assert.Equal(t, 3, b)
}

func TestArithmeticLogicUnit_Multiply(t *testing.T) {
	a := 2
	b := 3
	ArithmeticLogicUnit{}.multiply(&a, &b)
	assert.Equal(t, 6, a)
	assert.Equal(t, 3, b)
}

func TestArithmeticLogicUnit_Divide(t *testing.T) {
	a := 7
	b := 3
	ArithmeticLogicUnit{}.divide(&a, &b)
	assert.Equal(t, 2, a)
	assert.Equal(t, 3, b)
}

func TestArithmeticLogicUnit_Mod(t *testing.T) {
	a := 7
	b := 3
	ArithmeticLogicUnit{}.mod(&a, &b)
	assert.Equal(t, 1, a)
	assert.Equal(t, 3, b)
}

func TestArithmeticLogicUnit_Equal(t *testing.T) {
	a := 7
	b := 3
	ArithmeticLogicUnit{}.equal(&a, &b)
	assert.Equal(t, 0, a)
	assert.Equal(t, 3, b)

	a = 3
	ArithmeticLogicUnit{}.equal(&a, &b)
	assert.Equal(t, 1, a)
	assert.Equal(t, 3, b)
}
