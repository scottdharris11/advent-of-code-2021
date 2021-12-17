package day17

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	line := utils.ReadLines("day17", "day-17-input.txt")[0]
	solvePart1(line)
	solvePart2(line)
}

func solvePart1(line string) int {
	target := parseInput(line)
	start := time.Now().UnixMilli()
	xValues := findPossibleXValues(target.minX, target.maxX)
	yValues := findPossibleYValues(target.minY)
	ans := 0
	for _, x := range xValues {
		for _, y := range yValues {
			p := Probe{xVelocity: x, yVelocity: y}
			hit, maxHeight := p.Fire(target)
			if hit && maxHeight > ans {
				ans = maxHeight
			}
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 1 (%dms): Max Height = %d", end-start, ans)
	return ans
}

func solvePart2(line string) int {
	target := parseInput(line)
	start := time.Now().UnixMilli()
	xValues := findPossibleXValues(target.minX, target.maxX)
	yValues := findPossibleYValues(target.minY)
	ans := 0
	for _, x := range xValues {
		for _, y := range yValues {
			p := Probe{xVelocity: x, yVelocity: y}
			hit, _ := p.Fire(target)
			if hit {
				ans++
			}
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 2 (%dms): Combos = %d", end-start, ans)
	return ans
}

func parseInput(line string) *Area {
	matcher := regexp.MustCompile(`^target area: x=(.+)\.\.(.+), y=(.+)\.\.(.+)$`)
	matches := matcher.FindStringSubmatch(line)

	number := func(v string) int {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("Not a number: ", v)
		}
		return n
	}

	return &Area{
		minX: number(matches[1]),
		maxX: number(matches[2]),
		minY: number(matches[3]),
		maxY: number(matches[4]),
	}
}

func findPossibleXValues(tMin int, tMax int) []int {
	var possibleX []int
	for x := 1; x < tMin; x++ {
		sum := x
		for y := x - 1; y > 0; y-- {
			sum += y
			if sum >= tMin && sum <= tMax {
				possibleX = append(possibleX, x)
				break
			}
		}
	}
	for x := tMin; x <= tMax; x++ {
		possibleX = append(possibleX, x)
	}
	return possibleX
}

func findPossibleYValues(tMin int) []int {
	var possibleY []int
	cMax := int(math.Abs(float64(tMin))) - 1
	for x := cMax; x >= tMin; x-- {
		possibleY = append(possibleY, x)
	}
	return possibleY
}

type Area struct {
	minX int
	maxX int
	minY int
	maxY int
}

type Probe struct {
	xPos      int
	yPos      int
	xVelocity int
	yVelocity int
}

func (p *Probe) Fire(target *Area) (bool, int) {
	hitsTarget := false
	maxHeight := 0
	for {
		p.Step()
		if p.yPos > maxHeight {
			maxHeight = p.yPos
		}
		if p.In(target) {
			hitsTarget = true
			break
		}
		if p.NotReachable(target) {
			break
		}
	}
	return hitsTarget, maxHeight
}

func (p *Probe) Step() {
	p.xPos += p.xVelocity
	p.yPos += p.yVelocity
	if p.xVelocity != 0 {
		p.xVelocity--
	}
	p.yVelocity--
}

func (p *Probe) In(a *Area) bool {
	return p.xPos >= a.minX && p.xPos <= a.maxX && p.yPos >= a.minY && p.yPos <= a.maxY
}

func (p *Probe) NotReachable(a *Area) bool {
	return p.xPos > a.maxX || (p.xPos < a.minX && p.xVelocity == 0) || (p.yPos < a.minY && p.yVelocity <= 0)
}
