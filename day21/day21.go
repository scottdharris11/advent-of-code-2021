package day21

import (
	"fmt"
	"log"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day21", "day-21-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	game := parseInput(lines)
	start := time.Now().UnixMilli()
	game.Play(NewDeterministicDie(), 1000)
	ans := game.rollCnt * game.Loser().score
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 1 (%dms): Game Result = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	game := parseInput(lines)
	start := time.Now().UnixMilli()
	ans := playWithQuantumDie(*game, 21)
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 2 (%dms): Highest Universe Wins = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) *Game {
	return &Game{
		player1: Pawn{position: utils.Number(lines[0][28:])},
		player2: Pawn{position: utils.Number(lines[1][28:])},
	}
}

type Game struct {
	rollCnt int
	player1 Pawn
	player2 Pawn
	done    bool
	winner  int
	p2Turn  bool
	simCnt  int
}

func (g *Game) String() string {
	return fmt.Sprintf("Done: %t, Score: %d-%d", g.done, g.player1.score, g.player2.score)
}

func (g *Game) Play(die DeterministicDie, tillScore int) {
	for !g.done {
		g.TakeTurn(die.Roll(), tillScore)
	}
}

func (g *Game) TakeTurn(rollResult int, tillScore int) bool {
	g.rollCnt += 3
	pawn := &g.player1
	winner := 1
	if g.p2Turn {
		pawn = &g.player2
		winner = 2
		g.p2Turn = false
	} else {
		g.p2Turn = true
	}
	pawn.Move(rollResult)
	if pawn.score >= tillScore {
		g.winner = winner
		g.done = true
	}
	return g.done
}

func (g *Game) Winner() *Pawn {
	if !g.done {
		return nil
	}
	if g.winner == 1 {
		return &g.player1
	}
	return &g.player2
}

func (g *Game) Loser() *Pawn {
	if !g.done {
		return nil
	}
	if g.winner == 1 {
		return &g.player2
	}
	return &g.player1
}

type Pawn struct {
	position int
	score    int
}

func (p *Pawn) Move(positions int) {
	p.position += positions
	if p.position > 10 {
		p.position -= 10
	}
	p.score += p.position
}

func NewDeterministicDie() DeterministicDie {
	return DeterministicDie{nextRoll: 6}
}

type DeterministicDie struct {
	nextRoll int
}

func (d *DeterministicDie) Roll() int {
	rollAmt := d.nextRoll
	d.nextRoll--
	if d.nextRoll < 0 {
		d.nextRoll = 9
	}
	return rollAmt
}

var quantumGameCountsByRollResult = []int{0, 0, 0, 1, 3, 6, 7, 6, 3, 1}

func playWithQuantumDie(g Game, tillScore int) int {
	g.simCnt = 1
	games := []Game{g}
	activeGameCnt := 1
	winner := make([]int, 2)
	for activeGameCnt > 0 {
		activeIdx := 0
		nGames := make([]Game, 0, activeGameCnt)
		for i := 0; i < activeGameCnt; i++ {
			// for each currently active game, clone and simulate the results of 3-9 on next roll
			//   - when done, record winner count based on amount of universes the game is representing
			//   - when not done, schedule for next turn
			pGame := games[i]
			for m := 3; m <= 9; m++ {
				nGame := pGame
				count := quantumGameCountsByRollResult[m]
				nGame.simCnt *= count
				if nGame.TakeTurn(m, tillScore) {
					winner[nGame.winner-1] += nGame.simCnt
				} else {
					nGames = append(nGames, nGame)
					activeIdx++
				}
			}
		}
		games = nGames
		activeGameCnt = activeIdx
	}

	highestWinner := winner[0]
	if winner[1] > highestWinner {
		highestWinner = winner[1]
	}
	return highestWinner
}
