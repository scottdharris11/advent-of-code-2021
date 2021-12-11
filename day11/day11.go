package day11

import (
	"fmt"
	"log"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	oGrid := parseInput()
	start := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		oGrid.ExecuteStep()
	}
	ans := oGrid.FlashCount()
	end := time.Now().UnixMilli()
	log.Printf("Day 11, Part 1 (%dms): Flash Count = %d", end-start, ans)
}

func solvePart2() {
	oGrid := parseInput()
	start := time.Now().UnixMilli()
	step := 0
	for {
		step++
		if oGrid.ExecuteStep() {
			break
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 11, Part 2 (%dms): All Flash Step = %d", end-start, step)
}

func parseInput() *OctoGrid {
	lines := utils.ReadLines("day11", "day-11-input.txt")
	/*lines = []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}*/

	grid := utils.ReadIntegerGrid(lines)
	var oGrid [][]*Octo
	for i := 0; i < len(grid); i++ {
		var oGridRow []*Octo
		for j := 0; j < len(grid[i]); j++ {
			oGridRow = append(oGridRow, &Octo{energyLevel: grid[i][j]})
		}
		oGrid = append(oGrid, oGridRow)
	}
	return &OctoGrid{grid: oGrid}
}

type OctoGrid struct {
	grid [][]*Octo
}

func (og OctoGrid) ExecuteStep() bool {
	// increase octopus levels
	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			og.increaseOctoLevel(i, j)
		}
	}

	// finalize step
	allFlashed := true
	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			if !og.grid[i][j].FinishStep() {
				allFlashed = false
			}
		}
	}
	return allFlashed
}

func (og OctoGrid) String() string {
	var sb strings.Builder
	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			sb.WriteString(og.grid[i][j].String())
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (og OctoGrid) FlashCount() int {
	flashes := 0
	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			flashes += og.grid[i][j].flashCount
		}
	}
	return flashes
}

func (og OctoGrid) increaseOctoLevel(row int, col int) {
	// return if past edges
	if row < 0 || row >= len(og.grid) || col < 0 || col >= len(og.grid[0]) {
		return
	}

	// increase level at point and if flashed...increase level of all
	// surrounding octopus
	if og.grid[row][col].IncreaseLevel() {
		for i := row - 1; i <= row+1; i++ {
			for j := col - 1; j <= col+1; j++ {
				if i == row && j == col {
					continue
				}
				og.increaseOctoLevel(i, j)
			}
		}
	}
}

type Octo struct {
	energyLevel int
	flashCount  int
}

func (o *Octo) IncreaseLevel() bool {
	o.energyLevel++
	if o.energyLevel == 10 {
		o.flashCount++
		return true
	}
	return false
}

func (o *Octo) FinishStep() bool {
	if o.energyLevel > 9 {
		o.energyLevel = 0
		return true
	}
	return false
}

func (o Octo) String() string {
	pLevel := o.energyLevel
	if pLevel > 9 {
		pLevel = 0
	}
	return fmt.Sprintf("%d", o.energyLevel)
}
