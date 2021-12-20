package day13

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day13", "day-13-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	paper, folds := parseInput(lines)
	start := time.Now().UnixMilli()
	paper.FoldPaper(folds[0])
	ans := paper.DotCount()
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 1 (%dms): Dot Count = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) string {
	paper, folds := parseInput(lines)
	start := time.Now().UnixMilli()
	for _, f := range folds {
		paper.FoldPaper(f)
	}
	ans := paper.String()
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 2 (%dms): Code = \n%s", end-start, ans)
	return ans
}

func parseInput(lines []string) (*Paper, []Fold) {
	// parse dots into paper grid
	xMax := 0
	yMax := 0
	var dots []Point
	fStart := 0
	for fStart = 0; fStart < len(lines); fStart++ {
		if lines[fStart] == "" {
			break
		}
		sCoords := strings.Split(lines[fStart], ",")
		x := utils.Number(sCoords[0])
		y := utils.Number(sCoords[1])
		dots = append(dots, Point{x: x, y: y})
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	xMax++
	yMax++
	grid := make([][]bool, yMax)
	for r := range grid {
		grid[r] = make([]bool, xMax)
	}
	for _, d := range dots {
		grid[d.y][d.x] = true
	}

	paper := &Paper{grid: grid, xMax: xMax, yMax: yMax}

	// parse folds
	fStart++
	var folds []Fold
	for ; fStart < len(lines); fStart++ {
		pieces := strings.Split(lines[fStart], " ")
		foldCmd := strings.Split(pieces[2], "=")
		value := utils.Number(foldCmd[1])
		xVal, yVal := 0, 0
		switch foldCmd[0] {
		case "y":
			yVal = value
		case "x":
			xVal = value
		}
		folds = append(folds, Fold{x: xVal, y: yVal})
	}
	return paper, folds
}

type Point struct {
	x int
	y int
}

type Fold struct {
	x int
	y int
}

type Paper struct {
	grid [][]bool
	xMax int
	yMax int
}

func (p Paper) String() string {
	sb := strings.Builder{}
	for _, row := range p.grid {
		for _, dot := range row {
			if dot {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (p Paper) DotCount() int {
	count := 0
	for _, row := range p.grid {
		for _, dot := range row {
			if dot {
				count++
			}
		}
	}
	return count
}

func (p *Paper) FoldPaper(f Fold) {
	if f.x > 0 {
		p.FoldHorizontal(f.x)
	}
	if f.y > 0 {
		p.FoldVertical(f.y)
	}
}

func (p *Paper) FoldHorizontal(foldCol int) {
	for rowIdx, row := range p.grid {
		for i, j := foldCol+1, foldCol-1; i < p.xMax; i++ {
			if row[i] {
				row[j] = true
			}
			j--
		}
		p.grid[rowIdx] = row[:foldCol]
	}
	p.xMax = foldCol
}

func (p *Paper) FoldVertical(foldRow int) {
	for i, j := foldRow+1, foldRow-1; i < p.yMax; i++ {
		for c := 0; c < p.xMax; c++ {
			if p.grid[i][c] {
				p.grid[j][c] = true
			}
		}
		j--
	}
	p.grid = p.grid[:foldRow]
	p.yMax = foldRow
}
