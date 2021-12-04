package day4

import (
	"log"
	"strconv"
	"strings"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	game := parseGame()
	winner := game.Play()
	score := winner.unmarkedSum * game.lastCalled
	log.Printf("Day 4, Part 1: winning score %d", score)
}

func solvePart2() {
	game := parseGame()
	winner := game.LastWinner()
	score := winner.unmarkedSum * game.lastCalled
	log.Printf("Day 4, Part 2: last winning score %d", score)
}

func parseGame() *Game {
	lines := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		"8  2 23  4 24",
		"21  9 14 16  7",
		"6 10  3 18  5",
		"1 12 20 15 19",
		"",
		"3 15  0  2 22",
		"9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		"2  0 12  3  7",
	}
	lines = utils.ReadLines("day4", "day-4-input.txt")

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
