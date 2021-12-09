package day9

import (
	"log"
	"sort"
	"strconv"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	caveMap := parseInput()
	start := time.Now().UnixMilli()
	risk := caveMap.riskOfLowPoints()
	end := time.Now().UnixMilli()
	log.Printf("Day 9, Part 1 (%dms): Risk Score %d", end-start, risk)
}

func solvePart2() {
	caveMap := parseInput()
	start := time.Now().UnixMilli()
	lBasins := caveMap.largestBasinsSize()
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 2 (%dms): Large Basin %d", end-start, lBasins)
}

func parseInput() *CaveMap {
	lines := utils.ReadLines("day9", "day-9-input.txt")
	//lines = []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"}
	var hMap [][]int
	for _, line := range lines {
		var hMapRow []int
		for _, c := range line {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalln("unexpected bad map value", c, err)
			}
			hMapRow = append(hMapRow, v)
		}
		hMap = append(hMap, hMapRow)
	}
	return &CaveMap{heightMap: hMap}
}

type Point struct {
	row   int
	col   int
	value int
}

type Basin struct {
	points []Point
}

type CaveMap struct {
	heightMap [][]int
}

func (m CaveMap) riskOfLowPoints() int {
	lowPoints := m.lowPoints()
	risk := 0
	for _, lowPoint := range lowPoints {
		risk += m.riskLevel(lowPoint.value)
	}
	return risk
}

func (m CaveMap) lowPoints() []Point {
	var points []Point
	for i, row := range m.heightMap {
		for j, value := range row {
			checkPoints := m.adjacentPoints(Point{row: i, col: j})
			allHigher := true
			for _, cp := range checkPoints {
				if value >= m.heightMap[cp.row][cp.col] {
					allHigher = false
					break
				}
			}
			if allHigher {
				points = append(points, Point{row: i, col: j, value: value})
			}
		}
	}
	return points
}

func (m CaveMap) adjacentPoints(p Point) []Point {
	rMax := len(m.heightMap) - 1
	cMax := len(m.heightMap[0]) - 1

	var points []Point
	checkPoints := [][]int{{p.row - 1, p.col}, {p.row, p.col - 1}, {p.row, p.col + 1}, {p.row + 1, p.col}}
	for _, cp := range checkPoints {
		if cp[0] < 0 || cp[0] > rMax || cp[1] < 0 || cp[1] > cMax {
			continue
		}
		points = append(points, Point{
			row: cp[0], col: cp[1], value: m.heightMap[cp[0]][cp[1]],
		})
	}
	return points
}

func (CaveMap) riskLevel(i int) int {
	return i + 1
}

func (m CaveMap) largestBasinsSize() int {
	// find basins and build array of sizes
	basins := m.basins()
	var sizes []int
	for _, b := range basins {
		sizes = append(sizes, len(b.points))
	}

	// sort and multiply high three (from bottom of array)
	sort.Ints(sizes)
	c, size := len(sizes)-3, 1
	for i := len(sizes) - 1; i >= c; i-- {
		size *= sizes[i]
	}
	return size
}

func (m CaveMap) basins() []Basin {
	lowPoints := m.lowPoints()
	var basins []Basin
	for _, lowPoint := range lowPoints {
		points := []Point{lowPoint}
		m.higherPoints(lowPoint, &points)
		basins = append(basins, Basin{points: points})
	}
	return basins
}

func (m CaveMap) higherPoints(p Point, hPoints *[]Point) {
	aPoints := m.adjacentPoints(p)
	for _, aPoint := range aPoints {
		if aPoint.value < 9 && p.value < aPoint.value && !m.pointExists(aPoint, hPoints) {
			*hPoints = append(*hPoints, aPoint)
			m.higherPoints(aPoint, hPoints)
		}
	}
}

func (CaveMap) pointExists(p Point, points *[]Point) bool {
	for _, point := range *points {
		if point.row == p.row && point.col == p.col {
			return true
		}
	}
	return false
}
