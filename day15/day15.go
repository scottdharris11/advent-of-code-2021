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
	ans := grid.BestPath()
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 1 (%dms): Path Risk = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	grid := NewGrid(expandGrid(utils.ReadIntegerGrid(lines)))
	start := time.Now().UnixMilli()
	ans := grid.BestPath()
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 2 (%dms): Path Risk = %d", end-start, ans)
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

func (g *Grid) BestPath() int {
	search := utils.Search{Searcher: g}
	solution := search.Best(utils.SearchMove{State: g.nodes[0][0]})
	return solution.Cost
}

func (g *Grid) Goal(state interface{}) bool {
	var node = state.(Node)
	return node == g.nodes[g.bottomEdge][g.rightEdge]
}

func (g *Grid) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var n = state.(Node)
	var moves []utils.SearchMove
	checkPoints := [][]int{{n.gridX, n.gridY - 1}, {n.gridX - 1, n.gridY}, {n.gridX + 1, n.gridY}, {n.gridX, n.gridY + 1}}
	for _, point := range checkPoints {
		if point[0] >= 0 && point[0] <= g.rightEdge && point[1] >= 0 && point[1] <= g.bottomEdge {
			move := utils.SearchMove{
				Cost:  g.nodes[point[1]][point[0]].risk,
				State: g.nodes[point[1]][point[0]],
			}
			moves = append(moves, move)
		}
	}
	return moves
}

func (g *Grid) DistanceFromGoal(state interface{}) int {
	var node = state.(Node)
	return g.manhattanDistance(node, g.nodes[g.bottomEdge][g.rightEdge])
}

// compute estimate distance between two nodes based on distance between grid points
func (Grid) manhattanDistance(a Node, b Node) int {
	return int(math.Abs(float64(a.gridX-b.gridX)) + math.Abs(float64(a.gridY-b.gridY)))
}

type Node struct {
	gridX int
	gridY int
	risk  int
}
