package day3

import (
	"log"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	r := readDiagnosticReport()
	start := time.Now().UnixMilli()
	p := r.powerRating()
	end := time.Now().UnixMilli()
	log.Printf("Day 3, Part 1 (%dms): Power Rating = %d", end-start, p)
}

func solvePart2() {
	r := readDiagnosticReport()
	start := time.Now().UnixMilli()
	o := r.oxygenRating()
	c := r.co2ScrubberRating()
	l := o * c
	end := time.Now().UnixMilli()
	log.Printf("Day 3, Part 2 (%dms): Oxygen Rating = %d, CO2 Rating = %d, Life Support Rating = %d", end-start, o, c, l)
}

func readDiagnosticReport() DiagnosticReport {
	lines := utils.ReadLines("day3", "day-3-input.txt")
	// lines = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
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
		} else {
			epsilon |= 1 << bitIdx
		}
	}
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
