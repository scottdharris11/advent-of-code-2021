package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2021/utils"
)

var testValue = "target area: x=20..30, y=-10..-5"

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 45, solvePart1(testValue))
	assert.Equal(t, 11175, solvePart1(utils.ReadLines("day17", "day-17-input.txt")[0]))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 112, solvePart2(testValue))
	assert.Equal(t, 3540, solvePart2(utils.ReadLines("day17", "day-17-input.txt")[0]))
}

func TestProbeStep(t *testing.T) {
	p := Probe{xVelocity: 2, yVelocity: 2}
	p.Step()
	assert.Equal(t, 2, p.xPos)
	assert.Equal(t, 2, p.yPos)
	assert.Equal(t, 1, p.xVelocity)
	assert.Equal(t, 1, p.yVelocity)

	p.Step()
	assert.Equal(t, 3, p.xPos)
	assert.Equal(t, 3, p.yPos)
	assert.Equal(t, 0, p.xVelocity)
	assert.Equal(t, 0, p.yVelocity)

	p.Step()
	assert.Equal(t, 3, p.xPos)
	assert.Equal(t, 3, p.yPos)
	assert.Equal(t, 0, p.xVelocity)
	assert.Equal(t, -1, p.yVelocity)

	p.Step()
	assert.Equal(t, 3, p.xPos)
	assert.Equal(t, 2, p.yPos)
	assert.Equal(t, 0, p.xVelocity)
	assert.Equal(t, -2, p.yVelocity)

	p.Step()
	assert.Equal(t, 3, p.xPos)
	assert.Equal(t, 0, p.yPos)
	assert.Equal(t, 0, p.xVelocity)
	assert.Equal(t, -3, p.yVelocity)
}

func TestInArea(t *testing.T) {
	a := &Area{minX: 4, maxX: 7, minY: -5, maxY: 5}

	p := Probe{xPos: 3, yPos: 3}
	assert.False(t, p.In(a))
	p = Probe{xPos: 5, yPos: 3}
	assert.True(t, p.In(a))
	p = Probe{xPos: 5, yPos: -3}
	assert.True(t, p.In(a))
	p = Probe{xPos: 5, yPos: 7}
	assert.False(t, p.In(a))
}

func TestAreaNotReachable(t *testing.T) {
	a := &Area{minX: 4, maxX: 7, minY: -5, maxY: 5}

	p := Probe{xPos: 3, yPos: 3, xVelocity: 3, yVelocity: -2}
	assert.False(t, p.NotReachable(a))

	p = Probe{xPos: 8, yPos: 3, yVelocity: 0}
	assert.True(t, p.NotReachable(a))

	p = Probe{xPos: 5, yPos: -10, yVelocity: 0}
	assert.True(t, p.NotReachable(a))

	p = Probe{xPos: 3, yPos: -10, xVelocity: 3, yVelocity: 5}
	assert.False(t, p.NotReachable(a))

	p = Probe{xPos: 3, yPos: 10, xVelocity: 0, yVelocity: -2}
	assert.True(t, p.NotReachable(a))
}
