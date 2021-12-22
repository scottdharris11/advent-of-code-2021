package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValues = []string{
	"on x=10..12,y=10..12,z=10..12",
	"on x=11..13,y=11..13,z=11..13",
	"off x=9..11,y=9..11,z=9..11",
	"on x=10..10,y=10..10,z=10..10",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 39, solvePart1(testValues))
	assert.Equal(t, 590784, solvePart1(utils.ReadLines("day22", "day-22-test.txt")))
	assert.Equal(t, 474140, solvePart1(utils.ReadLines("day22", "day-22-test2.txt")))
	assert.Equal(t, 655005, solvePart1(utils.ReadLines("day22", "day-22-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 2758514936282235, solvePart2(utils.ReadLines("day22", "day-22-test2.txt")))
	assert.Equal(t, 1125649856443608, solvePart2(utils.ReadLines("day22", "day-22-input.txt")))
}

func TestParseInput(t *testing.T) {
	cuboids := parseInput(testValues)

	assert.Equal(t, 4, len(cuboids))

	assert.True(t, cuboids[0].on)
	assert.ElementsMatch(t, [2]int{10, 12}, cuboids[0].xRange)
	assert.ElementsMatch(t, [2]int{10, 12}, cuboids[0].yRange)
	assert.ElementsMatch(t, [2]int{10, 12}, cuboids[0].zRange)

	assert.True(t, cuboids[0].on)
	assert.ElementsMatch(t, [2]int{11, 13}, cuboids[1].xRange)
	assert.ElementsMatch(t, [2]int{11, 13}, cuboids[1].yRange)
	assert.ElementsMatch(t, [2]int{11, 13}, cuboids[1].zRange)

	assert.False(t, cuboids[2].on)
	assert.ElementsMatch(t, [2]int{9, 11}, cuboids[2].xRange)
	assert.ElementsMatch(t, [2]int{9, 11}, cuboids[2].yRange)
	assert.ElementsMatch(t, [2]int{9, 11}, cuboids[2].zRange)

	assert.True(t, cuboids[3].on)
	assert.ElementsMatch(t, [2]int{10, 10}, cuboids[3].xRange)
	assert.ElementsMatch(t, [2]int{10, 10}, cuboids[3].yRange)
	assert.ElementsMatch(t, [2]int{10, 10}, cuboids[3].zRange)
}

func TestReactor_BreakCuboid(t *testing.T) {
	tests := []struct {
		name     string
		original Cuboid
		avoid    Cuboid
		expected []Cuboid
	}{
		{
			"1",
			Cuboid{on: true, xRange: [2]int{1, 3}, yRange: [2]int{4, 6}, zRange: [2]int{7, 9}},
			Cuboid{on: false, xRange: [2]int{2, 4}, yRange: [2]int{3, 5}, zRange: [2]int{6, 8}},
			[]Cuboid{
				{on: true, xRange: [2]int{1, 1}, yRange: [2]int{4, 6}, zRange: [2]int{7, 9}},
				{on: true, xRange: [2]int{2, 3}, yRange: [2]int{6, 6}, zRange: [2]int{7, 9}},
				{on: true, xRange: [2]int{2, 3}, yRange: [2]int{4, 5}, zRange: [2]int{9, 9}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oCuboids := Reactor{}.breakCuboid(&tt.original, &tt.avoid)
			var nCuboids []Cuboid
			for _, c := range oCuboids {
				nCuboids = append(nCuboids, *c)
			}
			assert.ElementsMatch(t, tt.expected, nCuboids)
		})
	}
}

func TestReactor_ClearRanges(t *testing.T) {
	tests := []struct {
		name     string
		check    [2]int
		avoid    [2]int
		expected [][2]int
	}{
		{"1", [2]int{1, 3}, [2]int{3, 5}, [][2]int{{1, 2}}},
		{"2", [2]int{3, 5}, [2]int{1, 3}, [][2]int{{4, 5}}},
		{"3", [2]int{3, 9}, [2]int{5, 7}, [][2]int{{3, 4}, {8, 9}}},
		{"4", [2]int{3, 9}, [2]int{3, 9}, [][2]int{}},
		{"5", [2]int{4, 8}, [2]int{3, 9}, [][2]int{}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.expected, Reactor{}.noOverlapRanges(tt.check, tt.avoid))
		})
	}
}

func TestCuboid_Cubes(t *testing.T) {
	tests := []struct {
		name     string
		cuboid   Cuboid
		expected int
	}{
		{"1", Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}}, 27},
		{"2", Cuboid{xRange: [2]int{-10, -12}, yRange: [2]int{-10, -12}, zRange: [2]int{-10, -12}}, 27},
		{"3", Cuboid{xRange: [2]int{-1, 1}, yRange: [2]int{-1, 1}, zRange: [2]int{-1, 1}}, 27},
		{"4", Cuboid{xRange: [2]int{11, 13}, yRange: [2]int{11, 20}, zRange: [2]int{11, 15}}, 150},
		{"5", Cuboid{xRange: [2]int{10, 10}, yRange: [2]int{10, 10}, zRange: [2]int{10, 10}}, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.cuboid.Cubes())
		})
	}
}

func TestCuboid_IntersectsAndOverlays(t *testing.T) {
	tests := []struct {
		name       string
		cuboid1    Cuboid
		cuboid2    Cuboid
		intersects bool
		over12     bool
		over21     bool
	}{
		{
			"1",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			true, true, true,
		},
		{
			"2",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{-10, -12}},
			false, false, false,
		},
		{
			"3",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{-10, -12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			false, false, false,
		},
		{
			"4",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{-10, -12}, zRange: [2]int{10, 12}},
			false, false, false,
		},
		{
			"5",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{9, 13}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			true, false, true,
		},
		{
			"6",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{11, 11}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			true, true, false,
		},
		{
			"7",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{11, 13}, yRange: [2]int{11, 13}, zRange: [2]int{11, 13}},
			true, false, false,
		},
		{
			"8",
			Cuboid{xRange: [2]int{10, 12}, yRange: [2]int{10, 12}, zRange: [2]int{10, 12}},
			Cuboid{xRange: [2]int{9, 11}, yRange: [2]int{9, 11}, zRange: [2]int{9, 11}},
			true, false, false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.intersects, tt.cuboid1.Intersects(&tt.cuboid2))
			assert.Equal(t, tt.intersects, tt.cuboid2.Intersects(&tt.cuboid1))
			assert.Equal(t, tt.over12, tt.cuboid1.Overlays(&tt.cuboid2))
			assert.Equal(t, tt.over21, tt.cuboid2.Overlays(&tt.cuboid1))
		})
	}
}

func TestCuboid_RangeSize(t *testing.T) {
	tests := []struct {
		name     string
		iRange   [2]int
		expected int
	}{
		{"1", [2]int{2, 5}, 4},
		{"2", [2]int{-5, -2}, 4},
		{"3", [2]int{-5, 2}, 8},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Cuboid{}.rangeSize(tt.iRange))
		})
	}
}

func TestCuboid_RangeIntersectsAndOverlays(t *testing.T) {
	tests := []struct {
		name       string
		range1     [2]int
		range2     [2]int
		intersects bool
		over12     bool
		over21     bool
	}{
		{"1", [2]int{2, 5}, [2]int{3, 6}, true, false, false},
		{"2", [2]int{2, 5}, [2]int{-2, 3}, true, false, false},
		{"3", [2]int{2, 5}, [2]int{6, 7}, false, false, false},
		{"4", [2]int{2, 5}, [2]int{-2, 30}, true, false, true},
		{"5", [2]int{2, 5}, [2]int{2, 4}, true, true, false},
		{"6", [2]int{2, 5}, [2]int{3, 5}, true, true, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.intersects, Cuboid{}.rangeIntersects(tt.range1, tt.range2))
			assert.Equal(t, tt.intersects, Cuboid{}.rangeIntersects(tt.range2, tt.range1))
			assert.Equal(t, tt.over12, Cuboid{}.rangeOverlays(tt.range1, tt.range2))
			assert.Equal(t, tt.over21, Cuboid{}.rangeOverlays(tt.range2, tt.range1))
		})
	}
}
