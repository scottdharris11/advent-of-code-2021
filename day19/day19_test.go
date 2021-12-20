package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 79, solvePart1(utils.ReadLines("day19", "day-19-test.txt")))
	assert.Equal(t, 326, solvePart1(utils.ReadLines("day19", "day-19-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3621, solvePart2(utils.ReadLines("day19", "day-19-test.txt")))
	assert.Equal(t, 10630, solvePart2(utils.ReadLines("day19", "day-19-input.txt")))
}

func TestCoordinateSubtract(t *testing.T) {
	tests := []struct {
		name     string
		c1       Coordinate
		c2       Coordinate
		expected Coordinate
	}{
		{"1", Coordinate{4, 1, 2}, Coordinate{-1, -1, -1}, Coordinate{5, 2, 3}},
		{"2", Coordinate{0, 2, 1}, Coordinate{-5, 0, -2}, Coordinate{5, 2, 3}},
		{"3", Coordinate{3, 3, 3}, Coordinate{-2, 1, 0}, Coordinate{5, 2, 3}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c1.Subtract(tt.c2))
		})
	}
}

func TestCoordinateAdd(t *testing.T) {
	tests := []struct {
		name     string
		c1       Coordinate
		c2       Coordinate
		expected Coordinate
	}{
		{"1", Coordinate{5, 2, 3}, Coordinate{-1, -1, -1}, Coordinate{4, 1, 2}},
		{"2", Coordinate{5, 2, 3}, Coordinate{-5, 0, -2}, Coordinate{0, 2, 1}},
		{"3", Coordinate{5, 2, 3}, Coordinate{-2, 1, 0}, Coordinate{3, 3, 3}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c1.Add(tt.c2))
		})
	}
}

func TestCoordinateManhattanDistance(t *testing.T) {
	c := Coordinate{1105, -1205, 1229}
	c2 := Coordinate{-92, -2380, -20}
	md := c.ManhattanDistance(c2)
	assert.Equal(t, 3621, md)
	md = c2.ManhattanDistance(c)
	assert.Equal(t, 3621, md)
}

func TestCoordinateRotate(t *testing.T) {
	c := Coordinate{2, 1, 3}
	xLeft := RotateXAxis(c, 1)
	assert.Equal(t, Coordinate{1, -2, 3}, xLeft)
	xBack := RotateXAxis(c, 2)
	assert.Equal(t, Coordinate{-2, -1, 3}, xBack)
	xRight := RotateXAxis(c, 3)
	assert.Equal(t, Coordinate{-1, 2, 3}, xRight)
	xBegin := RotateXAxis(c, 4)
	assert.Equal(t, Coordinate{2, 1, 3}, xBegin)

	top := RotateZAxis(c, 1)
	assert.Equal(t, Coordinate{3, 1, -2}, top)
	bottom := RotateZAxis(c, 3)
	assert.Equal(t, Coordinate{-3, 1, 2}, bottom)

	yLeft := RotateYAxis(c, 1)
	assert.Equal(t, Coordinate{2, -3, 1}, yLeft)
	yBack := RotateYAxis(c, 2)
	assert.Equal(t, Coordinate{2, -1, -3}, yBack)
	yRight := RotateYAxis(c, 3)
	assert.Equal(t, Coordinate{2, 3, -1}, yRight)
	yBegin := RotateYAxis(c, 4)
	assert.Equal(t, Coordinate{2, 1, 3}, yBegin)
}

func TestScannerDeterminePosition_SimpleOrientation(t *testing.T) {
	// build scanners (origin scanner)
	s1 := &Scanner{
		detectedCoords: []Coordinate{
			{-2, -2, -2},
			{0, 2, 1},
			{4, 1, 2},
			{3, 3, 3},
		},
	}
	s1.position = &Coordinate{0, 0, 0}
	s1.originRelativeCoords = s1.detectedCoords

	s2 := &Scanner{
		detectedCoords: []Coordinate{
			{-1, -1, -1},
			{-5, 0, -2},
			{-2, 1, 0},
			{3, 2, 2},
		},
	}

	// determine position and validate results
	determined := s2.DeterminePosition(s1, 3)
	assert.True(t, determined)
	assert.NotNil(t, s2.position)
	assert.Equal(t, Coordinate{5, 2, 3}, *s2.position)
	assert.NotNil(t, s2.originRelativeCoords)
	assert.Equal(t, 4, len(s2.originRelativeCoords))
	assert.Equal(t, Coordinate{4, 1, 2}, s2.originRelativeCoords[0])
	assert.Equal(t, Coordinate{0, 2, 1}, s2.originRelativeCoords[1])
	assert.Equal(t, Coordinate{3, 3, 3}, s2.originRelativeCoords[2])
	assert.Equal(t, Coordinate{8, 4, 5}, s2.originRelativeCoords[3])
}

func TestScannerDeterminePosition_DiffOrientation(t *testing.T) {
	// build scanners (origin scanner)
	s1 := &Scanner{
		detectedCoords: []Coordinate{
			{-2, -2, -2},
			{0, 2, 1},
			{4, 1, 2},
			{3, 3, 3},
		},
	}
	s1.position = &Coordinate{0, 0, 0}
	s1.originRelativeCoords = s1.detectedCoords

	s2 := &Scanner{
		detectedCoords: []Coordinate{
			{-1, 1, -1},
			{0, 5, -2},
			{1, 2, 0},
			{2, -3, 2},
		},
	}

	// determine position and validate results
	determined := s2.DeterminePosition(s1, 3)
	assert.True(t, determined)
	assert.NotNil(t, s2.position)
	if s2.position != nil {
		assert.Equal(t, Coordinate{5, 2, 3}, *s2.position)
		assert.NotNil(t, s2.originRelativeCoords)
		assert.Equal(t, 4, len(s2.originRelativeCoords))
		assert.Equal(t, Coordinate{4, 1, 2}, s2.originRelativeCoords[0])
		assert.Equal(t, Coordinate{0, 2, 1}, s2.originRelativeCoords[1])
		assert.Equal(t, Coordinate{3, 3, 3}, s2.originRelativeCoords[2])
		assert.Equal(t, Coordinate{8, 4, 5}, s2.originRelativeCoords[3])
	}
}

func TestScannerDeterminePosition_TestInput(t *testing.T) {
	// build scanners (origin scanner)
	lines := utils.ReadLines("day19", "day-19-test.txt")
	scanners := parseInput(lines)
	s0 := scanners[0]
	s0.position = &Coordinate{0, 0, 0}
	s0.originRelativeCoords = s0.detectedCoords
	s1 := scanners[1]

	// determine position and validate results
	determined := s1.DeterminePosition(s0, 12)
	assert.True(t, determined)
	assert.NotNil(t, s1.position)
	if s1.position != nil {
		assert.Equal(t, Coordinate{68, -1246, -43}, *s1.position)
		assert.NotNil(t, s1.originRelativeCoords)

		s4 := scanners[4]
		determined := s4.DeterminePosition(s1, 12)
		assert.True(t, determined)
		assert.NotNil(t, s4.position)
		if s4.position != nil {
			assert.Equal(t, Coordinate{-20, -1133, 1061}, *s4.position)
			assert.NotNil(t, s4.originRelativeCoords)
		}
	}
}

func TestBuildRotationFunctions(t *testing.T) {
	c := Coordinate{1, 2, 3}
	rFunctions := *buildRotationFunctions()
	assert.Equal(t, 24, len(rFunctions))
	rMap := make(map[Coordinate]map[int]int)
	for i, rFunc := range rFunctions {
		value := rFunc(c)
		fMap := rMap[value]
		if fMap == nil {
			fMap = make(map[int]int)
			rMap[value] = fMap
		}
		fMap[i]++
	}
	assert.Equal(t, 24, len(rMap))
}
