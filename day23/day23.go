package day23

import (
	"log"
	"math"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day23", "day-23-input.txt")
	solvePart1(lines)
	// solvePart2(lines)
}

func solvePart1(lines []string) int {
	state := parseInput(lines)
	start := time.Now().UnixMilli()
	ans := BestSolution(*state)
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 1 (%dms): Least energy = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	state := parseInput(lines)
	start := time.Now().UnixMilli()
	ans := BestSolution(*state)
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 2 (%dms): Least energy = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) *BurrowState {
	var diagram [][]rune
	typeGoalCoords := make(map[rune][2]int, 4)
	var amphipods [16]Amphipod

	roomType := 'A'
	amphipodsIdx := 0
	for y := 0; y < len(lines); y++ {
		locations := []rune(lines[y])
		for x, lType := range locations {
			stepEnergy := 1
			switch lType {
			case 'B':
				stepEnergy = 10
			case 'C':
				stepEnergy = 100
			case 'D':
				stepEnergy = 1000
			}

			switch lType {
			case 'A', 'B', 'C', 'D':
				amphipods[amphipodsIdx] = Amphipod{aType: lType, locX: x, locY: y, energyPerStep: stepEnergy}
				amphipodsIdx++
				if roomType <= 'D' {
					typeGoalCoords[roomType] = [2]int{x, y}
					roomType++
				}
				locations[x] = '.'
			}
		}
		diagram = append(diagram, locations)
	}
	return &BurrowState{
		burrow:      &Burrow{diagram: diagram, roomCoords: typeGoalCoords},
		amphipods:   amphipods,
		amphipodCnt: amphipodsIdx,
	}
}

type Amphipod struct {
	aType         rune
	locX          int
	locY          int
	energyPerStep int
}

type Burrow struct {
	diagram    [][]rune
	roomCoords map[rune][2]int
}

func (b *Burrow) Space(x int, y int) bool {
	if y < 0 || y >= len(b.diagram) {
		return false
	}
	if x < 0 || x >= len(b.diagram[y]) {
		return false
	}
	return b.diagram[y][x] == '.'
}

func (b *Burrow) Room(x int, y int, aType rune) (room bool, forType bool, bottom bool) {
	for t, pos := range b.roomCoords {
		if pos[0] == x && pos[1] <= y {
			room = true
			forType = t == aType
			bottom = !b.Space(x, y+1)
			break
		}
	}
	return
}

func (b *Burrow) RoomEntrance(x int, y int) bool {
	for _, pos := range b.roomCoords {
		if pos[0] == x && pos[1]-1 == y {
			return true
		}
	}
	return false
}

type BurrowState struct {
	burrow      *Burrow
	amphipods   [16]Amphipod
	amphipodCnt int
}

func (s *BurrowState) Goal() bool {
	for i := 0; i < s.amphipodCnt; i++ {
		a := s.amphipods[i]
		goalLoc := s.burrow.roomCoords[a.aType]
		if goalLoc[0] != a.locX {
			return false
		}
	}
	return true
}

func (s *BurrowState) PossibleMoves() []Move {
	var moves []Move
	for i := 0; i < s.amphipodCnt; i++ {
		a := s.amphipods[i]
		// determine if amphipod is currently in room or hall, when...
		//   a. in a room at bottom of same type -> nothing to do...at goal
		//   b. in a room of wrong type -> check to see if blocked -> when not, move to hall
		//   c. in a room at top of same type -> check to see if blocking a different type -> if so, move to hall
		//   d. in hall -> check for blockers on path to room -> when none, move to room
		room, sType, bottom := s.burrow.Room(a.locX, a.locY, a.aType)
		if room {
			if sType && bottom {
				continue
			}

			roomTop := s.burrow.roomCoords[a.aType][1]

			if !sType {
				y := a.locY - 1
				blocked := false
				for y >= roomTop {
					if s.amphipodAt(a.locX, y) != nil {
						blocked = true
						break
					}
					y--
				}
				if !blocked {
					moves = append(moves, s.hallFromRoomMoves(a, roomTop)...)
				}
				continue
			}

			if s.blocking(a) {
				moves = append(moves, s.hallFromRoomMoves(a, roomTop)...)
			}
			continue
		}
		if state := s.hallToRoomMove(a); state != nil {
			moves = append(moves, *state)
		}
	}
	return moves
}

func (s *BurrowState) MoveAmphipod(fromX int, fromY int, toX int, toY int) int {
	a := s.amphipodAt(fromX, fromY)
	steps := int(math.Abs(float64(fromX-toX)) + math.Abs(float64(fromY-toY)))
	energyUsed := steps * a.energyPerStep
	a.locX = toX
	a.locY = toY
	return energyUsed
}

func (s *BurrowState) DistanceFromGoal() int {
	distance := 0
	for i := 0; i < s.amphipodCnt; i++ {
		a := s.amphipods[i]
		goalLoc := s.burrow.roomCoords[a.aType]
		xDistance := int(math.Abs(float64(goalLoc[0] - a.locX)))
		yDistance := 0
		if a.locY < goalLoc[1] {
			yDistance = goalLoc[1] - a.locY
		}
		if xDistance > 0 && a.locY >= goalLoc[1] {
			yDistance = a.locY - goalLoc[1] + 2
		}
		if xDistance == 0 && s.blocking(a) {
			yDistance = goalLoc[1] - a.locY + 3
		}
		distance += (xDistance + yDistance) * a.energyPerStep
	}
	return distance
}

func (s *BurrowState) Copy() *BurrowState {
	dup := *s
	return &dup
}

func (s *BurrowState) Occupied(x int, y int) (bool, *Amphipod) {
	for i := 0; i < s.amphipodCnt; i++ {
		a := s.amphipods[i]
		if x == a.locX && y == a.locY {
			return true, &a
		}
	}
	return false, nil
}

func (s *BurrowState) String() string {
	sb := strings.Builder{}
	sb.WriteRune('\n')
	for y, row := range s.burrow.diagram {
		for x, loc := range row {
			occupied, a := s.Occupied(x, y)
			if occupied {
				sb.WriteRune(a.aType)
			} else {
				sb.WriteRune(loc)
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (s *BurrowState) amphipodAt(x int, y int) *Amphipod {
	for i := 0; i < s.amphipodCnt; i++ {
		a := s.amphipods[i]
		if x == a.locX && y == a.locY {
			return &s.amphipods[i]
		}
	}
	return nil
}

func (s *BurrowState) blocking(a Amphipod) bool {
	y := a.locY + 1
	for s.burrow.Space(a.locX, y) {
		aLoc := s.amphipodAt(a.locX, y)
		if aLoc != nil && aLoc.aType != a.aType {
			return true
		}
		y++
	}
	return false
}

func (s *BurrowState) hallFromRoomMoves(a Amphipod, roomTop int) []Move {
	// position on hall row
	hallY := roomTop - 1

	// walk to the left/right until wall or blocked and record valid states
	var moves []Move
	walkHall := func(xAdd int) {
		hallX := a.locX
		for {
			hallX += xAdd
			if !s.burrow.Space(hallX, hallY) || s.amphipodAt(hallX, hallY) != nil {
				break
			}

			if s.burrow.RoomEntrance(hallX, hallY) {
				continue
			}

			nState := s.Copy()
			energyUsed := nState.MoveAmphipod(a.locX, a.locY, hallX, hallY)
			moves = append(moves, Move{state: *nState, energyUsed: energyUsed})
		}
	}
	walkHall(1)
	walkHall(-1)
	return moves
}

func (s *BurrowState) hallToRoomMove(a Amphipod) *Move {
	roomCoords := s.burrow.roomCoords[a.aType]

	// check to see if the room is empty or only contains proper types
	roomY := roomCoords[1] - 1
	for {
		roomA := s.amphipodAt(roomCoords[0], roomY+1)
		if roomA != nil {
			if roomA.locY == roomCoords[1] {
				return nil
			}
			if roomA.aType != a.aType {
				return nil
			}
			break
		}
		roomY++
		if !s.burrow.Space(roomCoords[0], roomY+1) {
			break
		}
	}

	// walk path to the room entrance, if all clear, create new state
	hallX := a.locX
	xAdd := 1
	if a.locX > roomCoords[0] {
		xAdd = -1
	}
	for {
		hallX += xAdd
		if hallX == roomCoords[0] {
			nState := s.Copy()
			energyUsed := nState.MoveAmphipod(a.locX, a.locY, hallX, roomY)
			return &Move{state: *nState, energyUsed: energyUsed}
		}
		if s.amphipodAt(hallX, a.locY) != nil {
			break
		}
	}
	return nil
}

type Move struct {
	state      BurrowState
	energyUsed int
}

// BestSolution a-star search algorithm implementation:
//   process moves from start to goal using a priority queue based on
//   cost (energy used) to get to the current state plus the estimated
//   distance to the goal
func BestSolution(b BurrowState) int {
	searchQueue := utils.PriorityQueue{}
	searchQueue.Queue(b, 0)
	cost := make(map[BurrowState]int)
	from := make(map[BurrowState]BurrowState)
	cost[b] = 0
	var goal BurrowState
	for !searchQueue.Empty() {
		var current = searchQueue.Next().(BurrowState)
		// log.Println(current.String())
		if current.Goal() {
			goal = current
			break
		}

		for _, next := range current.PossibleMoves() {
			nCost := cost[current] + next.energyUsed
			cCost, ok := cost[next.state]
			if !ok || nCost < cCost {
				cost[next.state] = nCost
				// log.Println("Possible Next...", next.String())
				priority := nCost + next.state.DistanceFromGoal()
				searchQueue.Queue(next.state, priority)
				from[next.state] = current
			}
		}
	}

	// Print path
	/*
	   	current := goal
	       for current != b {
	           fmt.Println(current.String())
	           current = from[current]
	       }
	*/
	return cost[goal]
}
