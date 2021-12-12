package main

import (
	"adventofcode2021/input"
	"fmt"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	start := parseUndirectedGraph(&lines)
	fmt.Println(calculatePart1(start))
	fmt.Println(calculatePart2(start))
}

func calculatePart1(start *Cave) int {
	var visitor = Visitor{revisited: true}
	return countPaths(start, visitor)
}

func calculatePart2(start *Cave) int {
	var visitor = Visitor{}
	return countPaths(start, visitor)
}

func countPaths(cave *Cave, visitor Visitor) int {
	if cave.symbol == "end" {
		return 1
	}
	visitor.visit(cave)
	var paths = 0
	neigobors := visitor.neighborsOf(cave)
	for _, neighbor := range neigobors {
		paths += countPaths(neighbor, visitor)
	}
	return paths
}

func (visitor *Visitor) canVisit(cave *Cave) bool {
	if !isLowerCase(cave.symbol) {
		return true
	}
	if cave.symbol == "start" {
		return false
	}
	for _, symbol := range visitor.path {
		if symbol == cave.symbol && visitor.revisited {
			return false
		}
	}
	return true
}

func (visitor *Visitor) visit(cave *Cave) {
	if isLowerCase(cave.symbol) {
		for _, symbol := range visitor.path {
			if symbol == cave.symbol {
				visitor.revisited = true
			}
		}
	}
	visitor.path = append(visitor.path, cave.symbol)
}

func isLowerCase(s string) bool {
	return s == strings.ToLower(s)
}

func (visitor *Visitor) neighborsOf(cave *Cave) []*Cave {
	var unvisited []*Cave
	for _, neighbor := range cave.neighbors {
		if visitor.canVisit(neighbor) {
			unvisited = append(unvisited, neighbor)
		}
	}
	return unvisited
}

func parseUndirectedGraph(lines *[]string) *Cave {
	caves := make(map[string]*Cave)
	for _, line := range *lines {
		symbols := strings.Split(line, "-")
		left := getOrCreateCave(&caves, symbols[0])
		right := getOrCreateCave(&caves, symbols[1])
		left.neighbors = append(left.neighbors, right)
		right.neighbors = append(right.neighbors, left)
	}
	start := caves["start"]
	return start
}

func getOrCreateCave(caves *map[string]*Cave, symbol string) *Cave {
	cave, found := (*caves)[symbol]
	if found {
		return cave
	} else {
		cave = &Cave{symbol: symbol}
		(*caves)[symbol] = cave
		return cave
	}
}

type Cave struct {
	symbol    string
	neighbors []*Cave
}

type Visitor struct {
	revisited bool
	path      []string
}
