package day15

import (
	"log"
	"math"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day15", "day-15-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	grid := NewGrid(utils.ReadIntegerGrid(lines))
	start := time.Now().UnixMilli()
	_, cost := search(grid, grid.nodes[0][0], grid.nodes[grid.bottomEdge][grid.rightEdge])
	ans := cost
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 1 (%dms): Path Cost = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	grid := NewGrid(expandGrid(utils.ReadIntegerGrid(lines)))
	start := time.Now().UnixMilli()
	_, cost := search(grid, grid.nodes[0][0], grid.nodes[grid.bottomEdge][grid.rightEdge])
	ans := cost
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 2 (%dms): Path Cost = %d", end-start, ans)
	return ans
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

type Node struct {
	gridX int
	gridY int
	risk  int
}

type NodeQueueItem struct {
	node     Node
	priority int
}

type NodeQueue struct {
	nodes []NodeQueueItem
}

func (q *NodeQueue) empty() bool {
	return len(q.nodes) == 0
}

func (q *NodeQueue) next() Node {
	n := len(q.nodes) - 1
	next := q.nodes[n]
	q.nodes = q.nodes[:n]
	return next.node
}

func (q *NodeQueue) queue(n Node, priority int) {
	nqi := NodeQueueItem{node: n, priority: priority}

	insertIdx := -1
	for idx, item := range q.nodes {
		if item.priority < priority {
			insertIdx = idx
			break
		}
	}

	if insertIdx == -1 {
		q.nodes = append(q.nodes, nqi)
	} else {
		q.nodes = append(q.nodes[:insertIdx+1], q.nodes[insertIdx:]...)
		q.nodes[insertIdx] = nqi
	}
}

// a* search algorithm implementation
func search(grid *Grid, start Node, goal Node) ([]Node, int) {
	searchQueue := NodeQueue{}
	searchQueue.queue(start, 0)
	from := make(map[Node]Node)
	cost := make(map[Node]int)
	cost[start] = 0

	for !searchQueue.empty() {
		current := searchQueue.next()
		if current == goal {
			break
		}

		for _, next := range grid.neighborNodes(current) {
			nCost := cost[current] + next.risk
			cCost, ok := cost[next]
			if !ok || nCost < cCost {
				cost[next] = nCost
				priority := nCost + heuristic(next, goal)
				searchQueue.queue(next, priority)
				from[next] = current
			}
		}
	}

	return constructPath(start, goal, from)
}

func heuristic(a Node, b Node) int {
	return int(math.Abs(float64(a.gridX-b.gridX)) + math.Abs(float64(a.gridY-b.gridY)))
}

func constructPath(start Node, goal Node, from map[Node]Node) ([]Node, int) {
	var path []Node
	cost := 0
	current := goal
	for current != start {
		path = append(path, current)
		cost += current.risk
		current = from[current]
	}
	path = append(path, start)
	return path, cost
}
