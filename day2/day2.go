package day2

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day2", "day-2-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	commands := parseCommands(lines)
	sub := Submarine{}
	applyCommands(&sub, commands)
	ans := sub.depth * sub.horizontal
	end := time.Now().UnixMilli()
	log.Printf("Day 2, Part 1 (%dms): Horizontal = %d, Depth = %d, Answer = %d", end-start, sub.horizontal, sub.depth, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	commands := parseCommands(lines)
	sub := AdvancedSubmarine{}
	applyCommands(&sub, commands)
	ans := sub.depth * sub.horizontal
	end := time.Now().UnixMilli()
	log.Printf("Day 2, Part 1 (%dms): Horizontal = %d, Depth = %d, Answer = %d", end-start, sub.horizontal, sub.depth, ans)
	return ans
}

func parseCommands(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		commands = append(commands, parseCommand(line))
	}
	return commands
}

func parseCommand(s string) Command {
	pieces := strings.Split(s, " ")
	unit := utils.Number(pieces[1])
	return Command{dir: pieces[0], unit: unit}
}

func applyCommands(cf CommandFollower, commands []Command) {
	for _, command := range commands {
		cf.applyCommand(command)
	}
}

type Command struct {
	dir  string
	unit int
}

type CommandFollower interface {
	applyCommand(cmd Command)
}

type Submarine struct {
	horizontal int
	depth      int
}

func (sub *Submarine) applyCommand(cmd Command) {
	switch cmd.dir {
	case "forward":
		sub.horizontal += cmd.unit
	case "down":
		sub.depth += cmd.unit
	case "up":
		sub.depth -= cmd.unit
	}
}

type AdvancedSubmarine struct {
	horizontal int
	depth      int
	aim        int
}

func (sub *AdvancedSubmarine) applyCommand(cmd Command) {
	switch cmd.dir {
	case "forward":
		sub.horizontal += cmd.unit
		sub.depth += sub.aim * cmd.unit
	case "down":
		sub.aim += cmd.unit
	case "up":
		sub.aim -= cmd.unit
	}
}
