package day3

import (
	"log"

	"advent-of-code-2021/utils"
)

func SolveDay3Part1() {
	r := readDiagnosticReport("day-3-input.txt")
	p := r.powerRating()
	log.Printf("Day 3, Part 1: Power Rating = %d", p)
}

func SolveDay3Part2() {
	r := readDiagnosticReport("day-3-input.txt")
	o := r.oxygenRating()
	c := r.co2ScrubberRating()
	l := o * c
	log.Printf("Day 3, Part 2: Oxygen Rating = %d, CO2 Rating = %d, Life Support Rating = %d", o, c, l)
}

func readDiagnosticReport(file string) DiagnosticReport {
	//lines := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	lines := utils.ReadLines(file)
	return newDiagnosticReport(lines)
}

type DiagnosticReport struct {
	bits [][]bool
}

func newDiagnosticReport(lines []string) DiagnosticReport {
	var bits [][]bool
	for _, line := range lines {
		lineBits := make([]bool, len(line))
		for i, c := range line {
			lineBits[i] = string(c) == "1"
		}
		bits = append(bits, lineBits)
	}
	return DiagnosticReport{bits}
}

func (d DiagnosticReport) powerRating() int {
	gamma := 0
	epsilon := 0
	posCnt := len(d.bits[0])
	for i := 0; i < posCnt; i++ {
		bitIdx := posCnt - i - 1
		if mostCommonBit(d.bits, i) {
			gamma |= 1 << bitIdx
			//fmt.Printf("Gamma OR Mask: %16b, Gamma After: %16b\n", 1<<bitIdx, gamma)
		} else {
			epsilon |= 1 << bitIdx
			//fmt.Printf("Epsilon OR Mask: %16b, Epsilon After: %16b\n", 1<<bitIdx, epsilon)
		}
	}
	//fmt.Printf("%d(%16b) vs %d(%16b)\n", gamma, gamma, epsilon, epsilon)
	return gamma * epsilon
}

func (d DiagnosticReport) oxygenRating() int {
	posCnt := len(d.bits[0])
	activeLines := d.bits
	for i := 0; i < posCnt; i++ {
		cBit := mostCommonBit(activeLines, i)
		activeLines = linesWithBitMatch(activeLines, i, cBit)
		if len(activeLines) == 1 {
			break
		}
	}
	//fmt.Println(activeLines[0])
	return bitsToDecimal(activeLines[0])
}

func (d DiagnosticReport) co2ScrubberRating() int {
	posCnt := len(d.bits[0])
	activeLines := d.bits
	for i := 0; i < posCnt; i++ {
		cBit := !mostCommonBit(activeLines, i)
		activeLines = linesWithBitMatch(activeLines, i, cBit)
		if len(activeLines) == 1 {
			break
		}
	}
	//fmt.Println(activeLines[0])
	return bitsToDecimal(activeLines[0])
}

func mostCommonBit(bits [][]bool, pos int) bool {
	onCnt := 0
	offCnt := len(bits)
	for _, line := range bits {
		if line[pos] {
			onCnt++
			offCnt--
		}
	}
	return onCnt >= offCnt
}

func linesWithBitMatch(bits [][]bool, pos int, match bool) [][]bool {
	var matched [][]bool
	for _, line := range bits {
		if line[pos] == match {
			matched = append(matched, line)
		}
	}
	return matched
}

func bitsToDecimal(bits []bool) int {
	out := 0
	posCnt := len(bits)
	for i := 0; i < posCnt; i++ {
		bitIdx := posCnt - i - 1
		if bits[i] {
			out |= 1 << bitIdx
		}
	}
	return out
}
