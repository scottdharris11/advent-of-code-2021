package day4

import (
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day4", "day-4-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	game := parseGame(lines)
	start := time.Now().UnixMilli()
	winner := game.Play()
	score := winner.unmarkedSum * game.lastCalled
	end := time.Now().UnixMilli()
	log.Printf("Day 4, Part 1 (%dms): winning score %d", end-start, score)
	return score
}

func solvePart2(lines []string) int {
	game := parseGame(lines)
	start := time.Now().UnixMilli()
	winner := game.LastWinner()
	score := winner.unmarkedSum * game.lastCalled
	end := time.Now().UnixMilli()
	log.Printf("Day 4, Part 2 (%dms): last winning score %d", end-start, score)
	return score
}

func parseGame(lines []string) *Game {
	game := Game{}
	sNums := strings.Split(lines[0], ",")
	for _, sNum := range sNums {
		num, err := strconv.Atoi(sNum)
		if err != nil {
			log.Fatalln("Invalid number detected", sNum)
		}
		game.numbers = append(game.numbers, num)
	}

	sIdx := 2
	for {
		eIdx := sIdx + 4
		if len(lines) < eIdx {
			break
		}
		game.boards = append(game.boards, parseBoard(lines[sIdx:eIdx+1]))
		sIdx += 6
	}

	return &game
}

func parseBoard(lines []string) *Board {
	var numbers [5][5]int
	for i, line := range lines {
		sNums := strings.Split(line, " ")
		idx := 0
		for _, sNum := range sNums {
			if sNum == "" {
				continue
			}
			num, err := strconv.Atoi(sNum)
			if err != nil {
				log.Fatalln("Invalid number detected", sNum)
			}
			numbers[i][idx] = num
			idx++
		}
	}
	return NewBoard(numbers)
}

type Game struct {
	numbers    []int
	boards     []*Board
	lastCalled int
}

func (g *Game) Play() Board {
	var winner *Board
	for _, number := range g.numbers {
		for _, board := range g.boards {
			if board.NumberCalled(number) {
				winner = board
			}
		}
		g.lastCalled = number
		if winner != nil {
			break
		}
	}
	return *winner
}

func (g *Game) LastWinner() Board {
	var winner *Board
	for _, number := range g.numbers {
		allWinners := true
		for _, board := range g.boards {
			if !board.Winner() {
				allWinners = false
				if board.NumberCalled(number) {
					winner = board
				}
			}
		}
		if allWinners {
			break
		}
		g.lastCalled = number
	}
	return *winner
}

type Board struct {
	markedSum   int
	unmarkedSum int
	numbers     [5][5]int
	marked      [5][5]bool
}

func NewBoard(numbers [5][5]int) *Board {
	board := Board{numbers: numbers}
	for _, row := range numbers {
		for _, number := range row {
			board.unmarkedSum += number
		}
	}
	return &board
}

func (b *Board) NumberCalled(number int) bool {
	for i, row := range b.numbers {
		for j, col := range row {
			if col == number {
				b.marked[i][j] = true
				b.markedSum += number
				b.unmarkedSum -= number
				return b.Winner()
			}
		}
	}
	return false
}

func (b *Board) Winner() bool {
	// check rows
	for r := 0; r < 5; r++ {
		allMarked := true
		for c := 0; c < 5; c++ {
			if !b.marked[r][c] {
				allMarked = false
				break
			}
		}
		if allMarked {
			return true
		}
	}

	// check cols
	for c := 0; c < 5; c++ {
		allMarked := true
		for r := 0; r < 5; r++ {
			if !b.marked[r][c] {
				allMarked = false
				break
			}
		}
		if allMarked {
			return true
		}
	}

	return false
}
