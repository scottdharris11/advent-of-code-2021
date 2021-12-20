package day19

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day19", "day-19-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	scanners := parseInput(lines)
	start := time.Now().UnixMilli()
	determineScannerPositions(scanners)
	ans := countUniqueDetectedBeacons(scanners)
	end := time.Now().UnixMilli()
	log.Printf("Day 19, Part 1 (%dms): Unique Beacons = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	scanners := parseInput(lines)
	start := time.Now().UnixMilli()
	determineScannerPositions(scanners)
	ans := largestManhattanDistance(scanners)
	end := time.Now().UnixMilli()
	log.Printf("Day 19, Part 2 (%dms): Largest Manhattan Distance = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) []*Scanner {
	var scanners []*Scanner
	coordinates := make([]Coordinate, 0)
	for _, line := range lines {
		if len(line) == 0 {
			scanners = append(scanners, &Scanner{detectedCoords: coordinates})
			coordinates = make([]Coordinate, 0)
			continue
		}
		if strings.HasPrefix(line, "---") {
			continue
		}
		values := strings.Split(line, ",")
		coordinates = append(coordinates, Coordinate{
			utils.Number(values[0]),
			utils.Number(values[1]),
			utils.Number(values[2]),
		})
	}
	scanners = append(scanners, &Scanner{detectedCoords: coordinates})
	return scanners
}

func determineScannerPositions(scanners []*Scanner) {
	scanners[0].position = &Coordinate{0, 0, 0}
	scanners[0].originRelativeCoords = scanners[0].detectedCoords

	rScanners := scanners[0:1]
	for {
		allPositioned, newPosScanners := checkScanners(rScanners, scanners)
		rScanners = newPosScanners
		if allPositioned {
			break
		}
	}
}

func checkScanners(relativeScanners []*Scanner, allScanners []*Scanner) (bool, []*Scanner) {
	var nPositionedScanners []*Scanner
	allPositioned := true
	for _, rScanner := range relativeScanners {
		for _, scanner := range allScanners {
			if scanner.position != nil {
				continue
			}
			if scanner.DeterminePosition(rScanner, 12) {
				nPositionedScanners = append(nPositionedScanners, scanner)
				continue
			}
			allPositioned = false
		}
	}
	return allPositioned, nPositionedScanners
}

func countUniqueDetectedBeacons(scanners []*Scanner) int {
	uniqueBeacons := make(map[Coordinate]int)
	for _, s := range scanners {
		for _, b := range s.originRelativeCoords {
			uniqueBeacons[b]++
		}
	}
	return len(uniqueBeacons)
}

func largestManhattanDistance(scanners []*Scanner) int {
	largest := 0
	for _, s1 := range scanners {
		for _, s2 := range scanners {
			if s1 == s2 {
				continue
			}
			md := s1.position.ManhattanDistance(*s2.position)
			if md > largest {
				largest = md
			}
		}
	}
	return largest
}

type Coordinate struct {
	x, y, z int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

func (c Coordinate) Subtract(c2 Coordinate) Coordinate {
	return Coordinate{c.x - c2.x, c.y - c2.y, c.z - c2.z}
}

func (c Coordinate) Add(c2 Coordinate) Coordinate {
	return Coordinate{c.x + c2.x, c.y + c2.y, c.z + c2.z}
}

func (c Coordinate) ManhattanDistance(c2 Coordinate) int {
	return int(math.Abs(float64(c.x-c2.x)) + math.Abs(float64(c.y-c2.y)) + math.Abs(float64(c.z-c2.z)))
}

type Scanner struct {
	position             *Coordinate
	detectedCoords       []Coordinate
	originRelativeCoords []Coordinate
	rFunctions           *[]func(Coordinate) Coordinate
}

func (s *Scanner) DeterminePosition(relative *Scanner, reqMatches int) bool {
	// if position already determined, just return
	if s.position != nil {
		return true
	}

	// build function list if not already built
	if s.rFunctions == nil {
		s.rFunctions = buildRotationFunctions()
	}

	// map potential coordinates by subtracting each beacon on each scanner,
	// looking for any coordinate that matched more than the required amount.
	// if found then we will compute the origin relative coordinates for each
	// detected position.
	for _, rFunc := range *s.rFunctions {
		potentialCoords := make(map[Coordinate]int, len(s.detectedCoords)^2)
		for _, s0c := range relative.originRelativeCoords {
			for _, s1c := range s.detectedCoords {
				potentialCoords[s0c.Subtract(rFunc(s1c))]++
			}
		}

		for c, cnt := range potentialCoords {
			if cnt >= reqMatches {
				match := c
				s.position = &match
				break
			}
		}

		if s.position != nil {
			s.originRelativeCoords = make([]Coordinate, len(s.detectedCoords))
			for i, c := range s.detectedCoords {
				s.originRelativeCoords[i] = s.position.Add(rFunc(c))
			}
			return true
		}
	}
	return false
}

func buildRotationFunctions() *[]func(Coordinate) Coordinate {
	// build rotation function combinations for the 3 axis
	var rFunctions []func(Coordinate) Coordinate
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				xTimes := x
				yTimes := y
				zTimes := z
				rFunctions = append(rFunctions, func(c Coordinate) Coordinate {
					return RotateZAxis(RotateYAxis(RotateXAxis(c, xTimes), yTimes), zTimes)
				})
			}
		}
	}

	// test functions with a coordinate and look for unique functions (should be 24 combos)
	c := Coordinate{1, 2, 3}
	rMap := make(map[Coordinate]int)
	for i, rFunc := range rFunctions {
		rMap[rFunc(c)] = i
	}
	var dFunctions []func(Coordinate) Coordinate
	for _, idx := range rMap {
		dFunctions = append(dFunctions, rFunctions[idx])
	}
	return &dFunctions
}

func RotateXAxis(c Coordinate, times int) Coordinate {
	result := c
	for i := 0; i < times; i++ {
		nY := -result.x
		result.x = result.y
		result.y = nY
	}
	return result
}

func RotateYAxis(c Coordinate, times int) Coordinate {
	result := c
	for i := 0; i < times; i++ {
		nZ := result.y
		result.y = -result.z
		result.z = nZ
	}
	return result
}

func RotateZAxis(c Coordinate, times int) Coordinate {
	result := c
	for i := 0; i < times; i++ {
		nZ := -result.x
		result.x = result.z
		result.z = nZ
	}
	return result
}
