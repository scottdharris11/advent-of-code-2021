package day2

import (
	"log"
	"strconv"
	"strings"

	"advent-of-code-2021/utils"
)

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

func SolveDay2Part1() {
	/*commands := []Command{
		{dir: "forward", unit: 5},
		{dir: "down", unit: 5},
		{dir: "forward", unit: 8},
		{dir: "up", unit: 3},
		{dir: "down", unit: 8},
		{dir: "forward", unit: 2},
	}*/
	commands := readCommands("day-2-input.txt")

	sub := Submarine{}
	applyCommands(&sub, commands)

	ans := sub.depth * sub.horizontal
	log.Printf("Day 2, Part 1: Horizontal = %d, Depth = %d, Answer = %d", sub.horizontal, sub.depth, ans)
}

func SolveDay2Part2() {
	/*commands := []Command{
		{dir: "forward", unit: 5},
		{dir: "down", unit: 5},
		{dir: "forward", unit: 8},
		{dir: "up", unit: 3},
		{dir: "down", unit: 8},
		{dir: "forward", unit: 2},
	}*/
	commands := readCommands("day-2-input.txt")

	sub := AdvancedSubmarine{}
	applyCommands(&sub, commands)

	ans := sub.depth * sub.horizontal
	log.Printf("Day 2, Part 1: Horizontal = %d, Depth = %d, Answer = %d", sub.horizontal, sub.depth, ans)
}

func applyCommands(cf CommandFollower, commands []Command) {
	for _, command := range commands {
		cf.applyCommand(command)
	}
}

func readCommands(file string) []Command {
	lines := utils.ReadLines(file)
	var commands []Command
	for _, line := range lines {
		commands = append(commands, parseCommand(line))
	}
	return commands
}

func parseCommand(s string) Command {
	pieces := strings.Split(s, " ")
	unit, err := strconv.Atoi(pieces[1])
	if err != nil {
		log.Fatalln(err)
	}
	return Command{dir: pieces[0], unit: unit}
}
