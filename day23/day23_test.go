package day23

import (
	"advent-of-code-2021/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 12521, solvePart1(utils.ReadLines("day23", "day-23-test.txt")))
	assert.Equal(t, 14148, solvePart1(utils.ReadLines("day23", "day-23-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 44169, solvePart2(utils.ReadLines("day23", "day-23-test2.txt")))
	assert.Equal(t, 43814, solvePart2(utils.ReadLines("day23", "day-23-input2.txt")))
}

func TestBurrow_Space(t *testing.T) {
	burrow := Burrow{
		diagram: [][]rune{
			[]rune("####"),
			[]rune("#..#"),
			[]rune("#.##"),
			[]rune("#.#"),
			[]rune("##"),
		},
	}

	tests := []struct {
		name  string
		x     int
		y     int
		space bool
	}{
		{"1", -1, 0, false},
		{"2", 0, -1, false},
		{"3", 3, 3, false},
		{"4", 0, 5, false},
		{"5", 0, 0, false},
		{"6", 1, 1, true},
		{"7", 1, 3, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.space, burrow.Space(tt.x, tt.y))
		})
	}
}

func TestBurrow_Room(t *testing.T) {
	burrow := Burrow{
		diagram: [][]rune{
			[]rune("#####"),
			[]rune("#...#"),
			[]rune("#.#.#"),
			[]rune("#.##"),
			[]rune("####"),
		},
		roomCoords: map[rune][2]int{
			'A': {1, 2},
			'B': {3, 2},
		},
	}

	tests := []struct {
		name     string
		x        int
		y        int
		aType    rune
		room     bool
		sameType bool
		bottom   bool
	}{
		{"1", 1, 1, 'A', false, false, false},
		{"2", 1, 2, 'A', true, true, false},
		{"3", 1, 3, 'A', true, true, true},
		{"4", 3, 2, 'A', true, false, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			room, sameType, bottom := burrow.Room(tt.x, tt.y, tt.aType)
			assert.Equal(t, tt.room, room)
			assert.Equal(t, tt.sameType, sameType)
			assert.Equal(t, tt.bottom, bottom)
		})
	}
}

func TestBurrowState_Goal(t *testing.T) {
	burrow := &Burrow{
		roomCoords: map[rune][2]int{
			'A': {1, 4},
			'B': {5, 4},
		},
	}

	tests := []struct {
		name  string
		state BurrowState
		goal  bool
	}{
		{"1", BurrowState{burrow: burrow, amphipodCnt: 2, amphipods: [16]Amphipod{
			{aType: 'A', locX: 1, locY: 4},
			{aType: 'B', locX: 5, locY: 5},
		}}, true},
		{"2", BurrowState{burrow: burrow, amphipodCnt: 2, amphipods: [16]Amphipod{
			{aType: 'A', locX: 1, locY: 4},
			{aType: 'B', locX: 5, locY: 4},
		}}, true},
		{"3", BurrowState{burrow: burrow, amphipodCnt: 2, amphipods: [16]Amphipod{
			{aType: 'A', locX: 7, locY: 4},
			{aType: 'B', locX: 5, locY: 5},
		}}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.goal, tt.state.Goal())
		})
	}
}

func TestBurrowState_PossibleMoves(t *testing.T) {
	burrow := &Burrow{
		diagram: [][]rune{
			[]rune("#########"),
			[]rune("#.......#"),
			[]rune("###.#.###"),
			[]rune("  #.#.#"),
			[]rune("  #####"),
		},
		roomCoords: map[rune][2]int{
			'A': {3, 2},
			'B': {5, 2},
		},
	}

	tests := []struct {
		name     string
		state    BurrowState
		possible []Move
	}{
		{"1", BurrowState{burrow: burrow, amphipodCnt: 2, amphipods: [16]Amphipod{
			{aType: 'A', locX: 5, locY: 2, energyPerStep: 1},
			{aType: 'A', locX: 5, locY: 3, energyPerStep: 1},
		}}, []Move{
			{energyUsed: 5, state: BurrowState{burrow: burrow, amphipodCnt: 2, amphipods: [16]Amphipod{
				{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
				{aType: 'A', locX: 5, locY: 3, energyPerStep: 1},
			}}},
		}},
		{"2", BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
			{aType: 'A', locX: 1, locY: 1, energyPerStep: 1},
			{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
			{aType: 'B', locX: 7, locY: 1, energyPerStep: 10},
		}}, []Move{
			{energyUsed: 3, state: BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
				{aType: 'A', locX: 3, locY: 2, energyPerStep: 1},
				{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
				{aType: 'B', locX: 7, locY: 1, energyPerStep: 10},
			}}},
			{energyUsed: 40, state: BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
				{aType: 'A', locX: 1, locY: 1, energyPerStep: 1},
				{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
				{aType: 'B', locX: 5, locY: 3, energyPerStep: 10},
			}}},
		}},
		{"3", BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
			{aType: 'A', locX: 4, locY: 1, energyPerStep: 1},
			{aType: 'A', locX: 5, locY: 3, energyPerStep: 1},
			{aType: 'B', locX: 6, locY: 1, energyPerStep: 10},
		}}, []Move{
			{energyUsed: 3, state: BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
				{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
				{aType: 'A', locX: 5, locY: 3, energyPerStep: 1},
				{aType: 'B', locX: 6, locY: 1, energyPerStep: 10},
			}}},
		}},
		{"4", BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
			{aType: 'A', locX: 1, locY: 1, energyPerStep: 1},
			{aType: 'A', locX: 6, locY: 1, energyPerStep: 1},
			{aType: 'B', locX: 7, locY: 1, energyPerStep: 10},
		}}, []Move{
			{energyUsed: 4, state: BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
				{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
				{aType: 'A', locX: 6, locY: 1, energyPerStep: 1},
				{aType: 'B', locX: 7, locY: 1, energyPerStep: 10},
			}}},
			{energyUsed: 5, state: BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
				{aType: 'A', locX: 1, locY: 1, energyPerStep: 1},
				{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
				{aType: 'B', locX: 7, locY: 1, energyPerStep: 10},
			}}},
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			states := tt.state.PossibleMoves()
			assert.ElementsMatch(t, tt.possible, states)
		})
	}
}

func TestBurrowState_DistanceFromGoal(t *testing.T) {
	burrow := &Burrow{
		roomCoords: map[rune][2]int{
			'A': {1, 2},
			'B': {5, 2},
		},
	}
	amphipods := [16]Amphipod{
		{aType: 'A', locX: 3, locY: 2, energyPerStep: 1}, //4
		{aType: 'A', locX: 1, locY: 3, energyPerStep: 1},
		{aType: 'B', locX: 1, locY: 3, energyPerStep: 10}, //7
		{aType: 'B', locX: 6, locY: 1, energyPerStep: 10}, //2
	}
	state := BurrowState{burrow: burrow, amphipodCnt: 4, amphipods: amphipods}

	assert.Equal(t, 94, state.DistanceFromGoal())
}

func TestBurrowState_Copy(t *testing.T) {
	burrow := &Burrow{}
	amphipods := [16]Amphipod{
		{aType: 'A', locX: 3, locY: 4},
		{aType: 'A', locX: 1, locY: 5},
	}
	original := BurrowState{burrow: burrow, amphipodCnt: 2, amphipods: amphipods}
	copyA := original.Copy()

	assert.NotSame(t, original, copyA)
	assert.Same(t, original.burrow, copyA.burrow)
	assert.NotSame(t, original.amphipods, copyA.amphipods)
	assert.ElementsMatch(t, original.amphipods, copyA.amphipods)

	copyA.amphipods[0].locX = 10
	copyA.amphipods[0].locY = 9
	assert.Equal(t, 10, copyA.amphipods[0].locX)
	assert.Equal(t, 9, copyA.amphipods[0].locY)
	assert.Equal(t, 3, original.amphipods[0].locX)
	assert.Equal(t, 4, original.amphipods[0].locY)
}

func TestBurrowState_Occupied(t *testing.T) {
	amphipods := [16]Amphipod{
		{aType: 'A', locX: 3, locY: 4},
		{aType: 'A', locX: 1, locY: 5},
	}
	state := BurrowState{amphipodCnt: 2, amphipods: amphipods}

	tests := []struct {
		name     string
		x        int
		y        int
		occupied bool
		aPtr     *Amphipod
	}{
		{"1", 1, 2, false, nil},
		{"2", 3, 5, false, nil},
		{"3", 1, 5, true, &amphipods[1]},
		{"4", 1, 4, false, nil},
		{"5", 3, 4, true, &amphipods[0]},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			occupied, a := state.Occupied(tt.x, tt.y)
			assert.Equal(t, tt.occupied, occupied)
			assert.Equal(t, tt.aPtr, a)
		})
	}
}

func TestBurrowState_RoomAvailable(t *testing.T) {
	burrow := &Burrow{
		diagram: [][]rune{
			[]rune("#########"),
			[]rune("#.......#"),
			[]rune("###.#.###"),
			[]rune("  #.#.#"),
			[]rune("  #.#.#"),
			[]rune("  #####"),
		},
		roomCoords: map[rune][2]int{
			'A': {3, 2},
			'B': {5, 2},
		},
	}

	tests := []struct {
		name      string
		state     BurrowState
		aType     rune
		available bool
		firstIdx  int
	}{
		{"1", BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
			{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
			{aType: 'A', locX: 3, locY: 4, energyPerStep: 1},
			{aType: 'B', locX: 7, locY: 1, energyPerStep: 10},
		}}, 'A', true, 2},

		{"2", BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
			{aType: 'A', locX: 1, locY: 1, energyPerStep: 1},
			{aType: 'A', locX: 3, locY: 3, energyPerStep: 1},
			{aType: 'B', locX: 3, locY: 4, energyPerStep: 10},
		}}, 'A', false, -1},

		{"3", BurrowState{burrow: burrow, amphipodCnt: 3, amphipods: [16]Amphipod{
			{aType: 'A', locX: 1, locY: 1, energyPerStep: 1},
			{aType: 'A', locX: 1, locY: 2, energyPerStep: 1},
			{aType: 'B', locX: 5, locY: 4, energyPerStep: 10},
		}}, 'A', true, 4},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			avail, idx := tt.state.roomAvailable(tt.aType, tt.state.burrow.roomCoords[tt.aType])
			assert.Equal(t, tt.available, avail)
			assert.Equal(t, tt.firstIdx, idx)
		})
	}
}
