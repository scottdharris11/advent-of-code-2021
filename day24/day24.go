package day24

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day24", "day-24-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := MaximumModelNumber(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 1 (%dms): Largest Model Number = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := MinimumModelNumber(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 2 (%dms): Smallest Model Number = %d", end-start, ans)
	return ans
}

func MaximumModelNumber(allInst []string) int {
	digitCombos := *ValidModelDigitCombinations(allInst)
	digits := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	sectionStack := []int{13}
	searchStack := []int{0}
	for {
		section := sectionStack[len(sectionStack)-1]
		if section == -1 {
			break
		}
		sectionDigits := digitCombos[section]
		digitIdx := 13 - section

		searchVal := searchStack[len(searchStack)-1]
		found := false
		for j := digits[digitIdx]; j >= 0; j-- {
			if digitMap, ok := sectionDigits[j]; ok {
				if nSearch, ok := digitMap[searchVal]; ok {
					searchStack = append(searchStack, nSearch)
					sectionStack = append(sectionStack, section-1)
					digits[digitIdx] = j
					found = true
					break
				}
			}
		}

		if !found {
			n := len(sectionStack) - 1
			sectionStack = sectionStack[:n]
			n = len(searchStack) - 1
			searchStack = searchStack[:n]
		}
	}

	// convert digits into number
	sb := strings.Builder{}
	for _, d := range digits {
		sb.WriteString(strconv.Itoa(d))
	}
	return utils.Number(sb.String())
}

func MinimumModelNumber(allInst []string) int {
	digitCombos := *ValidModelDigitCombinations(allInst)
	digits := [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	sectionStack := []int{13}
	searchStack := []int{0}
	for {
		section := sectionStack[len(sectionStack)-1]
		if section == -1 {
			break
		}
		sectionDigits := digitCombos[section]
		digitIdx := 13 - section

		searchVal := searchStack[len(searchStack)-1]
		found := false
		for j := digits[digitIdx]; j <= 9; j++ {
			if digitMap, ok := sectionDigits[j]; ok {
				if nSearch, ok := digitMap[searchVal]; ok {
					searchStack = append(searchStack, nSearch)
					sectionStack = append(sectionStack, section-1)
					digits[digitIdx] = j
					found = true
					break
				}
			}
		}

		if !found {
			n := len(sectionStack) - 1
			sectionStack = sectionStack[:n]
			n = len(searchStack) - 1
			searchStack = searchStack[:n]
		}
	}

	// convert digits into number
	sb := strings.Builder{}
	for _, d := range digits {
		sb.WriteString(strconv.Itoa(d))
	}
	return utils.Number(sb.String())
}

func ValidModelDigitCombinations(allInst []string) *[]map[int]map[int]int {
	// read map from cache if available
	combos := make([]map[int]map[int]int, 14)
	cachePath := utils.FilePath("day24", "cache.json")
	data, err := ioutil.ReadFile(cachePath)
	if err == nil {
		err = json.Unmarshal(data, &combos)
		if err == nil {
			return &combos
		}
		log.Println("Unable to unmarshal combination map: ", err.Error())
	}

	// Build map of valid combinations by walking backwards from goal of zero to the z value
	// inputs that could achieve that value, and then up the different calculation sections
	// until we have a complete map
	log.Println("Constructing map of valid digit combinations...")
	validValues := map[int]int{0: 0}
	for s := 1; s < 15; s++ {
		sectionValidValues := make(map[int]int, 100)
		sectionInst := allInst[len(allInst)-(18*s) : len(allInst)-(18*(s-1))]
		sectionMap := make(map[int]map[int]int, 9)
		combos = append(combos, sectionMap)
		for w := 1; w <= 9; w++ {
			for z := 0; z <= 500000; z++ {
				acl := NewArithmeticLogicUnit(0, 0, 0, z)
				acl.RunProgram(sectionInst, w)

				zVal := *acl.z
				if _, ok := validValues[zVal]; ok {
					sectionValidValues[z] = zVal
					zValues, ok := sectionMap[w]
					if !ok {
						zValues = make(map[int]int, 100)
						sectionMap[w] = zValues
					}
					zValues[z] = zVal
				}
			}
		}
		log.Printf("Mapped values for section %d, valid z values: %d", 15-s, len(sectionValidValues))
		validValues = sectionValidValues
	}

	// Write data to cache file
	data, err = json.Marshal(combos)
	if err != nil {
		log.Println("Unable to marshal cache of digit combinations: ", err.Error())
	}
	err = ioutil.WriteFile(cachePath, data, 0644)
	if err != nil {
		log.Println("Unable to write cache of digit combinations: ", err.Error())
	}
	return &combos
}

type Monad struct {
	instructions []string
}

func (m *Monad) Valid(digits [14]int) bool {
	alu := NewArithmeticLogicUnit(0, 0, 0, 0)
	alu.RunProgram(m.instructions, digits[:]...)
	return *alu.z == 0
}

func (m *Monad) ValidModelNumber(modelNo int) bool {
	s := strconv.Itoa(modelNo)
	if len(s) != 14 {
		return false
	}
	if strings.ContainsAny(s, "0") {
		return false
	}

	digits := [14]int{}
	for i, d := range s {
		digits[i] = int(d - '0')
	}

	alu := NewArithmeticLogicUnit(0, 0, 0, 0)
	alu.RunProgram(m.instructions, digits[:]...)
	return *alu.z == 0
}

func NewArithmeticLogicUnit(w int, x int, y int, z int) *ArithmeticLogicUnit {
	return &ArithmeticLogicUnit{&w, &x, &y, &z}
}

type ArithmeticLogicUnit struct {
	w *int
	x *int
	y *int
	z *int
}

func (alu *ArithmeticLogicUnit) RunProgram(instructions []string, inputs ...int) {
	inputIdx := 0
	for _, instruction := range instructions {
		input := 0
		if strings.HasPrefix(instruction, "inp") {
			input = inputs[inputIdx]
			inputIdx++
		}
		alu.applyInstruction(instruction, input)
	}
}

func (alu *ArithmeticLogicUnit) applyInstruction(i string, v int) {
	pieces := strings.Split(i, " ")
	switch pieces[0] {
	case "inp":
		alu.assign(alu.valuePointer(pieces[1]), v)
	case "add":
		alu.add(alu.valuePointer(pieces[1]), alu.valuePointer(pieces[2]))
	case "mul":
		alu.multiply(alu.valuePointer(pieces[1]), alu.valuePointer(pieces[2]))
	case "div":
		alu.divide(alu.valuePointer(pieces[1]), alu.valuePointer(pieces[2]))
	case "mod":
		alu.mod(alu.valuePointer(pieces[1]), alu.valuePointer(pieces[2]))
	case "eql":
		alu.equal(alu.valuePointer(pieces[1]), alu.valuePointer(pieces[2]))
	}
}

func (alu *ArithmeticLogicUnit) valuePointer(s string) *int {
	switch s {
	case "w":
		return alu.w
	case "x":
		return alu.x
	case "y":
		return alu.y
	case "z":
		return alu.z
	default:
		value := utils.Number(s)
		return &value
	}
}

func (ArithmeticLogicUnit) assign(a *int, b int) {
	*a = b
}

func (ArithmeticLogicUnit) add(a *int, b *int) {
	*a += *b
}

func (ArithmeticLogicUnit) multiply(a *int, b *int) {
	*a *= *b
}

func (ArithmeticLogicUnit) divide(a *int, b *int) {
	*a /= *b
}

func (ArithmeticLogicUnit) mod(a *int, b *int) {
	*a %= *b
}

func (ArithmeticLogicUnit) equal(a *int, b *int) {
	if *a == *b {
		*a = 1
	} else {
		*a = 0
	}
}
