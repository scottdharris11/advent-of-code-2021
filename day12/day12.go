package day12

import (
	"advent-of-code-2021/utils"
	"log"
	"strings"
	"time"
	"unicode"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	sCave := parseInput()
	start := time.Now().UnixMilli()
	routes := buildRoutes(sCave, false)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 1 (%dms): Paths = %d", end-start, len(routes))
}

func solvePart2() {
	sCave := parseInput()
	start := time.Now().UnixMilli()
	routes := buildRoutes(sCave, true)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 2 (%dms): Paths = %d", end-start, len(routes))
}

func parseInput() *Cave {
	lines := utils.ReadLines("day12", "day-12-input.txt")
	//lines = []string{"start-A","start-b","A-c","A-b","b-d","A-end","b-end"}
	//lines = []string{"dc-end","HN-start","start-kj","dc-start","dc-HN","LN-dc","HN-end","kj-sa","kj-HN","kj-dc"}
	//lines = []string{"fs-end","he-DX","fs-he","start-DX","pj-DX","end-zg","zg-sl","zg-pj","pj-he","RW-he","fs-DX","pj-RW","zg-RW","start-pj","he-WI","zg-he","pj-fs","start-RW"}

	var start *Cave
	caves := make(map[string]*Cave)
	for _, line := range lines {
		s := strings.Split(line, "-")
		c1 := matchOrCreateCave(s[0], caves)
		c2 := matchOrCreateCave(s[1], caves)
		c1.connections[c2.name] = c2
		c2.connections[c1.name] = c1
		if c1.name == "start" {
			start = c1
		}
		if c2.name == "start" {
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
	if c.name == "start" {
		return false
	}
	if c.name == "end" {
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
			if connectingCave.name == "end" {
				routes = append(routes, nRoute)
				continue
			}
			routes = followConnections(connectingCave, nRoute, routes)
		}
	}
	return routes
}
