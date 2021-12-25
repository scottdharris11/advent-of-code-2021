package day25

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day25", "day-25-input.txt")
	solvePart1(lines)
	solvePart2()
}

func solvePart1(lines []string) int {
	sea := parseInput(lines)
	start := time.Now().UnixMilli()
	ans := sea.StepsToClearing()
	end := time.Now().UnixMilli()
	log.Printf("Day 25, Part 1 (%dms): Steps required = %d", end-start, ans)
	return ans
}

func solvePart2() int {
	start := time.Now().UnixMilli()
	ans := 50
	end := time.Now().UnixMilli()
	log.Printf("Day 25, Part 2 (%dms): Mission ACCOMPLISHED, Stars Acquired = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) *SeaFloor {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return &SeaFloor{grid: grid, maxWidth: len(grid[0]) - 1, maxHeight: len(grid) - 1}
}

type SeaFloor struct {
	grid      [][]rune
	maxWidth  int
	maxHeight int
}

func (s *SeaFloor) String() string {
	sb := strings.Builder{}
	for _, row := range s.grid {
		sb.WriteString(string(row))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (s *SeaFloor) StepsToClearing() int {
	steps := 0
	for {
		steps++
		if s.step() == 0 {
			break
		}
	}
	return steps
}

func (s *SeaFloor) step() int {
	nGrid := make([][]rune, len(s.grid))
	moves := 0

	// process horizontal pack
	for y, row := range s.grid {
		nGrid[y] = make([]rune, s.maxWidth+1)
		for x, r := range row {
			if nGrid[y][x] == 0 {
				nGrid[y][x] = '.'
			}
			if r == '>' {
				nIdx := s.checkHorizontalIdx(x)
				if s.grid[y][nIdx] == '.' {
					moves++
					nGrid[y][nIdx] = '>'
				} else {
					nGrid[y][x] = '>'
				}
			}
		}
	}

	// process vertical pack
	for y, row := range s.grid {
		for x, r := range row {
			if r == 'v' {
				nIdx := s.checkVerticalIdx(y)
				if s.grid[nIdx][x] != 'v' && nGrid[nIdx][x] == '.' {
					moves++
					nGrid[nIdx][x] = 'v'
				} else {
					nGrid[y][x] = 'v'
				}
			}
		}
	}

	s.grid = nGrid
	return moves
}

func (s *SeaFloor) checkHorizontalIdx(currIdx int) int {
	i := currIdx + 1
	if i > s.maxWidth {
		i = 0
	}
	return i
}

func (s *SeaFloor) checkVerticalIdx(currIdx int) int {
	i := currIdx + 1
	if i > s.maxHeight {
		i = 0
	}
	return i
}
