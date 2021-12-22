package day22

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day22", "day-22-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	cuboids := parseInput(lines)
	start := time.Now().UnixMilli()
	reactor := Reactor{}
	reactor.Initialize(cuboids)
	ans := reactor.OnCubes()
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 1 (%dms): On Cubes After Init = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	cuboids := parseInput(lines)
	start := time.Now().UnixMilli()
	reactor := Reactor{}
	reactor.Reboot(cuboids)
	ans := reactor.OnCubes()
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 2 (%dms): On Cubes After Reboot = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) []*Cuboid {
	matcher := regexp.MustCompile(`^(.+) x=(.+)\.\.(.+),y=(.+)\.\.(.+),z=(.+)\.\.(.+)$`)
	var cuboids []*Cuboid
	for _, line := range lines {
		groups := matcher.FindStringSubmatch(line)
		cuboids = append(cuboids, &Cuboid{
			on:     groups[1] == "on",
			xRange: [2]int{utils.Number(groups[2]), utils.Number(groups[3])},
			yRange: [2]int{utils.Number(groups[4]), utils.Number(groups[5])},
			zRange: [2]int{utils.Number(groups[6]), utils.Number(groups[7])},
		})
	}
	return cuboids
}

type Reactor struct {
	cuboids []*Cuboid
}

func (r *Reactor) Cubes() int {
	cnt := 0
	for _, c := range r.cuboids {
		cnt += c.Cubes()
	}
	return cnt
}

func (r *Reactor) OnCubes() int {
	on := 0
	for _, c := range r.cuboids {
		if c.on {
			on += c.Cubes()
		}
	}
	return on
}

func (r *Reactor) Initialize(cuboids []*Cuboid) {
	for _, c := range cuboids {
		if c.xRange[0] < -50 || c.xRange[1] > 50 ||
			c.yRange[0] < -50 || c.yRange[1] > 50 ||
			c.zRange[0] < -50 || c.zRange[1] > 50 {
			return
		}
		r.applyCuboid(c)
	}
}

func (r *Reactor) Reboot(cuboids []*Cuboid) {
	for _, c := range cuboids {
		r.applyCuboid(c)
	}
}

func (r *Reactor) applyCuboid(c *Cuboid) {
	startIdx := 0
	for {
		done, idx := r.handleIntersections(c, startIdx)
		startIdx = idx
		if done {
			break
		}
	}
	r.cuboids = append(r.cuboids, c)
}

func (r *Reactor) handleIntersections(c *Cuboid, startIdx int) (bool, int) {
	l := len(r.cuboids)
	for i := startIdx; i < l; i++ {
		cuboid := r.cuboids[i]
		if !cuboid.Intersects(c) {
			continue
		}

		if c.Overlays(cuboid) {
			// new cuboid overlays this whole existing cube, just remove it by
			// replacing it with the last slice entry and then trimming the slice
			r.cuboids[i] = r.cuboids[len(r.cuboids)-1]
			r.cuboids = r.cuboids[:len(r.cuboids)-1]
			return false, i
		}

		// new cuboid intersects with this one, break into cuboids that don't intersect
		nCuboids := make([]*Cuboid, 0, len(r.cuboids)+10)
		nCuboids = append(nCuboids, r.cuboids[0:i]...)
		nCuboids = append(nCuboids, r.breakCuboid(cuboid, c)...)
		restartIdx := len(nCuboids)
		nCuboids = append(nCuboids, r.cuboids[i+1:]...)
		r.cuboids = nCuboids
		return false, restartIdx
	}
	return true, 0
}

func (r Reactor) breakCuboid(c *Cuboid, avoid *Cuboid) []*Cuboid {
	var cuboids []*Cuboid

	// find ranges on each axis that don't overlap with cuboid at all
	xNoOverlaps := r.noOverlapRanges(c.xRange, avoid.xRange)
	yNoOverlaps := r.noOverlapRanges(c.yRange, avoid.yRange)
	zNoOverlaps := r.noOverlapRanges(c.zRange, avoid.zRange)

	// for each non-overlap range on x-axis, add cuboid with it and current y, z ranges.
	// track a new min, max for the x-axis to compensate if adding items later based on
	// the y and z axis.
	xMin := c.xRange[0]
	xMax := c.xRange[1]
	for _, x := range xNoOverlaps {
		cuboids = append(cuboids, &Cuboid{on: c.on, xRange: x, yRange: c.yRange, zRange: c.zRange})
		if x[1] < xMax {
			xMin = x[1] + 1
		}
		if x[0] > xMin {
			xMax = x[0] - 1
		}
	}

	// for each non-overlap on y-axis, add cuboid with it, current z ranges, and the adjusted x-axis range
	// track a new min, max for y-axis to compensate if adding items later based on z axis
	yMin := c.yRange[0]
	yMax := c.yRange[1]
	for _, y := range yNoOverlaps {
		cuboids = append(cuboids, &Cuboid{on: c.on, xRange: [2]int{xMin, xMax}, yRange: y, zRange: c.zRange})
		if y[1] < yMax {
			yMin = y[1] + 1
		}
		if y[0] > yMin {
			yMax = y[0] - 1
		}
	}

	// for each non-overlap on z-axis, add cuboid with it and adjusted x, y ranges
	for _, z := range zNoOverlaps {
		cuboids = append(cuboids, &Cuboid{on: c.on, xRange: [2]int{xMin, xMax}, yRange: [2]int{yMin, yMax}, zRange: z})
	}

	return cuboids
}

func (Reactor) noOverlapRanges(check [2]int, avoid [2]int) [][2]int {
	var ranges [][2]int
	if check[0] < avoid[0] {
		ranges = append(ranges, [2]int{check[0], avoid[0] - 1})
	}
	if check[1] > avoid[1] {
		ranges = append(ranges, [2]int{avoid[1] + 1, check[1]})
	}
	return ranges
}

type Cuboid struct {
	xRange [2]int
	yRange [2]int
	zRange [2]int
	on     bool
}

func (c Cuboid) Cubes() int {
	return c.rangeSize(c.xRange) * c.rangeSize(c.yRange) * c.rangeSize(c.zRange)
}

func (c Cuboid) Intersects(c2 *Cuboid) bool {
	return c.rangeIntersects(c.xRange, c2.xRange) &&
		c.rangeIntersects(c.yRange, c2.yRange) &&
		c.rangeIntersects(c.zRange, c2.zRange)
}

func (c Cuboid) Overlays(c2 *Cuboid) bool {
	return c.rangeOverlays(c.xRange, c2.xRange) &&
		c.rangeOverlays(c.yRange, c2.yRange) &&
		c.rangeOverlays(c.zRange, c2.zRange)
}

func (c Cuboid) String() string {
	return fmt.Sprintf("%t %d x=%d..%d, y=%d..%d, z=%d..%d",
		c.on, c.Cubes(), c.xRange[0], c.xRange[1], c.yRange[0], c.yRange[1], c.zRange[0], c.zRange[1])
}

func (Cuboid) rangeSize(r [2]int) int {
	return int(math.Abs(float64(r[0]-r[1]))) + 1
}

func (c Cuboid) rangeIntersects(r1 [2]int, r2 [2]int) bool {
	return (r1[0] >= r2[0] && r1[0] <= r2[1]) ||
		(r1[1] >= r2[0] && r1[1] <= r2[1]) ||
		c.rangeOverlays(r1, r2)
}

func (Cuboid) rangeOverlays(r1 [2]int, r2 [2]int) bool {
	return r1[0] <= r2[0] && r1[1] >= r2[1]
}
