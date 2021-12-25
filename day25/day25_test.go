package day25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 58, solvePart1(utils.ReadLines("day25", "day-25-test.txt")))
	assert.Equal(t, 489, solvePart1(utils.ReadLines("day25", "day-25-input.txt")))
}

func TestSeaFloor_Step(t *testing.T) {
	tests := []struct {
		name    string
		bGrid   [][]rune
		aGrid   [][]rune
		moveCnt int
	}{
		{"1", [][]rune{[]rune("...>>>>>...")}, [][]rune{[]rune("...>>>>.>..")}, 1},
		{"2", [][]rune{[]rune("...>>>>.>..")}, [][]rune{[]rune("...>>>.>.>.")}, 2},
		{"3", [][]rune{[]rune("...>>>.>.>>")}, [][]rune{[]rune(">..>>.>.>>.")}, 3},
		{"4", [][]rune{[]rune(">..>>>.>.>>")}, [][]rune{[]rune(".>.>>.>.>>>")}, 3},
		{"5", [][]rune{
			[]rune(".........."),
			[]rune(".>v....v.."),
			[]rune(".......>.."),
			[]rune(".........."),
		}, [][]rune{
			[]rune(".........."),
			[]rune(".>........"),
			[]rune("..v....v>."),
			[]rune(".........."),
		}, 3},
		{"6", [][]rune{
			[]rune("..v......."),
			[]rune(".>v....v.."),
			[]rune(".......>.."),
			[]rune("..v......."),
		}, [][]rune{
			[]rune("..v......."),
			[]rune(".>........"),
			[]rune("..v....v>."),
			[]rune("..v......."),
		}, 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			sea := SeaFloor{grid: tt.bGrid, maxHeight: len(tt.bGrid) - 1, maxWidth: len(tt.bGrid[0]) - 1}
			moves := sea.step()
			fmt.Println(sea.String())
			assert.Equal(t, tt.moveCnt, moves)
			assert.Equal(t, tt.aGrid, sea.grid)
		})
	}
}
