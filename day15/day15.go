package day15

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day15", "day-15-input.txt")
	solvePart1(lines, false)
	solvePart2(lines, false)
}

func solvePart1(lines []string, printPath bool) int {
	grid := NewGrid(utils.ReadIntegerGrid(lines))
	start := time.Now().UnixMilli()
	path := grid.BestPath()
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 1 (%dms): Path Risk = %d", end-start, path.Risk)
	if printPath {
		log.Printf("Path Details: %s", path)
	}
	return path.Risk
}

func solvePart2(lines []string, printPath bool) int {
	grid := NewGrid(expandGrid(utils.ReadIntegerGrid(lines)))
	start := time.Now().UnixMilli()
	path := grid.BestPath()
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 2 (%dms): Path Risk = %d", end-start, path.Risk)
	if printPath {
		log.Printf("Path Details: %s", path)
	}
	return path.Risk
}

func expandGrid(grid [][]int) [][]int {
	xLen := len(grid[0])
	yLen := len(grid)
	width := len(grid[0]) * 5
	height := len(grid) * 5

	eGrid := make([][]int, height)
	for i := 0; i < height; i++ {
		eGrid[i] = make([]int, width)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y := 0; y < yLen; y++ {
				for x := 0; x < xLen; x++ {
					val := grid[y][x] + i + j
					if val > 9 {
						val -= 9
					}
					eGrid[y+(i*yLen)][x+(j*xLen)] = val
				}
			}
		}
	}
	return eGrid
}

func NewGrid(riskGrid [][]int) *Grid {
	nodeRows := make([][]Node, len(riskGrid))
	for i, riskRow := range riskGrid {
		nodes := make([]Node, len(riskRow))
		for j, risk := range riskRow {
			nodes[j] = Node{gridX: j, gridY: i, risk: risk}
		}
		nodeRows[i] = nodes
	}
	rightEdge := len(nodeRows[0]) - 1
	bottomEdge := len(nodeRows) - 1
	return &Grid{nodes: nodeRows, rightEdge: rightEdge, bottomEdge: bottomEdge}
}

type Grid struct {
	nodes      [][]Node
	rightEdge  int
	bottomEdge int
}

// BestPath a-star search algorithm implementation:
//   process nodes from start to goal using a priority queue based on
//   cost (risk) to get to the node plus the estimated distance (using
//   manhattan distance) to the goal
func (g *Grid) BestPath() *NodePath {
	start := g.nodes[0][0]
	goal := g.nodes[g.bottomEdge][g.rightEdge]
	searchQueue := utils.PriorityQueue{}
	searchQueue.Queue(start, 0)
	from := make(map[Node]Node)
	cost := make(map[Node]int)
	visited := make(map[Node]bool)
	cost[start] = 0

	for !searchQueue.Empty() {
		var current = searchQueue.Next().(Node)
		visited[current] = true
		if current == goal {
			break
		}

		for _, next := range g.neighborNodes(current) {
			nCost := cost[current] + next.risk
			cCost, ok := cost[next]
			if !ok || nCost < cCost {
				cost[next] = nCost
				priority := nCost + g.manhattanDistance(next, goal)
				searchQueue.Queue(next, priority)
				from[next] = current
			}
		}
	}

	return &NodePath{
		start:    start,
		goal:     goal,
		from:     from,
		Risk:     cost[goal],
		Visited:  len(visited),
		GridSize: (g.rightEdge + 1) * (g.bottomEdge + 1),
	}
}

// find the nodes that neighbor the supplied without going off edge
func (g *Grid) neighborNodes(n Node) []Node {
	var nodes []Node
	checkPoints := [][]int{{n.gridX, n.gridY - 1}, {n.gridX - 1, n.gridY}, {n.gridX + 1, n.gridY}, {n.gridX, n.gridY + 1}}
	for _, point := range checkPoints {
		if point[0] >= 0 && point[0] <= g.rightEdge && point[1] >= 0 && point[1] <= g.bottomEdge {
			nodes = append(nodes, g.nodes[point[1]][point[0]])
		}
	}
	return nodes
}

// compute estimate distance between two nodes based on distance between grid points
func (Grid) manhattanDistance(a Node, b Node) int {
	return int(math.Abs(float64(a.gridX-b.gridX)) + math.Abs(float64(a.gridY-b.gridY)))
}

type NodePath struct {
	start    Node
	goal     Node
	from     map[Node]Node
	Risk     int
	Visited  int
	GridSize int
}

func (np *NodePath) ConstructPath() []Node {
	var path []Node
	current := np.goal
	for current != np.start {
		path = append(path, current)
		current = np.from[current]
	}
	path = append(path, np.start)
	return path
}

func (np *NodePath) String() string {
	sb := strings.Builder{}
	sb.WriteString("Path Risk: ")
	sb.WriteString(strconv.Itoa(np.Risk))
	sb.WriteString(", Nodes Visited: ")
	sb.WriteString(strconv.Itoa(np.Visited))
	sb.WriteRune('/')
	sb.WriteString(strconv.Itoa(np.GridSize))
	sb.WriteString(", Path: ")
	path := np.ConstructPath()
	l := len(path)
	for i := l - 1; i >= 0; i-- {
		if i < l-1 {
			sb.WriteString("->")
		}
		sb.WriteRune('(')
		sb.WriteString(strconv.Itoa(path[i].gridY))
		sb.WriteRune(',')
		sb.WriteString(strconv.Itoa(path[i].gridX))
		sb.WriteRune(')')
	}
	return sb.String()
}

type Node struct {
	gridX int
	gridY int
	risk  int
}
