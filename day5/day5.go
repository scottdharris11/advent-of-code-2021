package day5

import (
	"advent-of-code-2021/utils"
	"log"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	lines := parseInput()
	grid := newGrid(lines, false)
	overlaps := grid.overlaps()
	log.Printf("Day 5, Part 1: overlapping spots %d", overlaps)
}

func solvePart2() {
	lines := parseInput()
	grid := newGrid(lines, true)
	overlaps := grid.overlaps()
	log.Printf("Day 5, Part 2: overlapping spots %d", overlaps)
}

func parseInput() []Line {
	sLines := []string{
		"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1",
		"7,0 -> 7,4", "6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4",
		"0,0 -> 8,8", "5,5 -> 8,2",
	}
	sLines = utils.ReadLines("day5", "day-5-input.txt")

	var lines []Line
	for _, sLine := range sLines {
		pieces := strings.Split(sLine, " ")
		if len(pieces) != 3 {
			log.Fatalln("unexpected line format:", sLine)
		}
		line := Line{start: parseCoordinate(pieces[0]), end: parseCoordinate(pieces[2])}
		lines = append(lines, line)
	}
	return lines
}

func parseCoordinate(s string) Coordinate {
	pieces := strings.Split(s, ",")
	x, err := strconv.Atoi(pieces[0])
	if err != nil {
		log.Fatalln("unexpected non-numeric value", x, err)
	}
	y, err := strconv.Atoi(pieces[1])
	if err != nil {
		log.Fatalln("unexpected non-numeric value", y, err)
	}
	return Coordinate{x: x, y: y}
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Equal(c2 Coordinate) bool {
	return c.x == c2.x && c.y == c2.y
}

type Line struct {
	start Coordinate
	end   Coordinate
}

func newGrid(lines []Line, incDiags bool) *Grid {
	maxX, maxY := 0, 0
	for _, line := range lines {
		if line.start.x > maxX {
			maxX = line.start.x
		}
		if line.end.x > maxX {
			maxX = line.end.x
		}
		if line.start.y > maxY {
			maxY = line.start.y
		}
		if line.end.y > maxY {
			maxY = line.end.y
		}
	}
	maxX++
	maxY++

	points := make([][]int, maxY)
	for i := range points {
		points[i] = make([]int, maxX)
	}

	grid := Grid{points: points}
	for _, line := range lines {
		if !incDiags && !(line.start.x == line.end.x || line.start.y == line.end.y) {
			continue
		}
		grid.addLine(line)
	}
	return &grid
}

type Grid struct {
	points [][]int
}

func (g *Grid) addLine(line Line) {
	xInc := 0
	if line.start.x < line.end.x {
		xInc = 1
	}
	if line.start.x > line.end.x {
		xInc = -1
	}
	yInc := 0
	if line.start.y < line.end.y {
		yInc = 1
	}
	if line.start.y > line.end.y {
		yInc = -1
	}

	p := line.start
	for {
		g.points[p.x][p.y] += 1
		if p.Equal(line.end) {
			break
		}
		p = Coordinate{x: p.x + xInc, y: p.y + yInc}
	}
}

func (g *Grid) overlaps() int {
	overlaps := 0
	for _, row := range g.points {
		for _, point := range row {
			if point > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}
