package day12

import (
	"advent-of-code-2021/utils"
	"log"
	"strings"
	"time"
	"unicode"
)

const startCaveName = "start"
const endCaveName = "end"

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day12", "day-12-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	sCave := parseInput(lines)
	start := time.Now().UnixMilli()
	routes := buildRoutes(sCave, false)
	routeCnt := len(routes)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 1 (%dms): Paths = %d", end-start, routeCnt)
	return routeCnt
}

func solvePart2(lines []string) int {
	sCave := parseInput(lines)
	start := time.Now().UnixMilli()
	routes := buildRoutes(sCave, true)
	routeCnt := len(routes)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 2 (%dms): Paths = %d", end-start, routeCnt)
	return routeCnt
}

func parseInput(lines []string) *Cave {
	var start *Cave
	caves := make(map[string]*Cave)
	for _, line := range lines {
		s := strings.Split(line, "-")
		c1 := matchOrCreateCave(s[0], caves)
		c2 := matchOrCreateCave(s[1], caves)
		c1.connections[c2.name] = c2
		c2.connections[c1.name] = c1
		if c1.name == startCaveName {
			start = c1
		}
		if c2.name == startCaveName {
			start = c2
		}
	}
	return start
}

func matchOrCreateCave(name string, caves map[string]*Cave) *Cave {
	cave := caves[name]
	if cave == nil {
		cave = &Cave{name: name}
		cave.small = allLowerCase(name)
		cave.connections = make(map[string]*Cave)
		caves[name] = cave
	}
	return cave
}

func allLowerCase(name string) bool {
	for _, c := range name {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

type Cave struct {
	name        string
	small       bool
	connections map[string]*Cave
}

type Route struct {
	path           []string
	allowDup       bool
	smallCaveDuped bool
}

func (r Route) PathWouldBeValid(c *Cave) bool {
	if c.name == startCaveName {
		return false
	}
	if c.name == endCaveName {
		return true
	}
	if c.small {
		if r.allowDup && !r.smallCaveDuped {
			return true
		}
		for _, p := range r.path {
			if p == c.name {
				return false
			}
		}
	}
	return true
}

func (r *Route) AddEntry(c *Cave) {
	if c.small {
		for _, p := range r.path {
			if p == c.name {
				r.smallCaveDuped = true
			}
		}
	}
	r.path = append(r.path, c.name)
}

func (r Route) String() string {
	return strings.Join(r.path, ",")
}

func buildRoutes(start *Cave, allowDup bool) []*Route {
	var routes []*Route
	route := &Route{
		path:     []string{start.name},
		allowDup: allowDup,
	}
	routes = followConnections(start, route, routes)
	return routes
}

func followConnections(c *Cave, r *Route, routes []*Route) []*Route {
	for _, connectingCave := range c.connections {
		if r.PathWouldBeValid(connectingCave) {
			nRoute := &Route{
				path:           r.path,
				allowDup:       r.allowDup,
				smallCaveDuped: r.smallCaveDuped,
			}
			nRoute.AddEntry(connectingCave)
			if connectingCave.name == endCaveName {
				routes = append(routes, nRoute)
				continue
			}
			routes = followConnections(connectingCave, nRoute, routes)
		}
	}
	return routes
}
