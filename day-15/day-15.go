package main

import (
	"aoc2021/input"
	"fmt"
	"math"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	calculatePart1(lines)
	calculatePart2(lines)
}

func calculatePart1(lines []string) {
	fmt.Println(calculateLowestRisk(newCavern(lines, 1)))
}

func calculatePart2(lines []string) {
	fmt.Println(calculateLowestRisk(newCavern(lines, 5)))
}

func calculateLowestRisk(riskMap RiskMap) int {
	var totalSpaces = riskMap.size()
	var src = 0
	var dest = totalSpaces - 1

	var distances = make([]int, totalSpaces)
	var visited = make([]bool, totalSpaces)
	for i := 0; i < totalSpaces; i++ {
		distances[i] = math.MaxInt
		visited[i] = false
	}
	distances[src] = 0

	for i := 0; i < totalSpaces; i++ {
		var min = math.MaxInt
		var minPos = -1
		for pos := range distances {
			if !visited[pos] && distances[pos] < min {
				minPos = pos
				min = distances[pos]
			}
		}
		visited[minPos] = true
		for _, neighbor := range riskMap.adjacent(minPos) {
			if distances[minPos]+riskMap.riskAt(neighbor) < distances[neighbor] {
				distances[neighbor] = distances[minPos] + riskMap.riskAt(neighbor)
			}
		}
	}
	return distances[dest]
}

type RiskMap interface {
	riskAt(pos int) int
	adjacent(pos int) []int
	size() int
}

func newCavern(lines []string, scale int) *Cavern {
	var w = len(lines[0])
	var h = len(lines)
	var risks = make([]int, 0)
	for _, line := range lines {
		for _, c := range line {
			risks = append(risks, int(c-'0'))
		}
	}
	return &Cavern{originalWidth: w, originalHeight: h, scale: scale, risks: risks}
}

type Cavern struct {
	originalWidth  int
	originalHeight int
	scale          int
	risks          []int
}

func (cavern *Cavern) adjacent(pos int) []int {
	var adjacent []int
	var x = pos % cavern.width()
	if x > 0 {
		adjacent = append(adjacent, pos-1)
	}
	if x < cavern.width()-1 {
		adjacent = append(adjacent, pos+1)
	}

	var y = pos / cavern.height()
	if y > 0 {
		adjacent = append(adjacent, pos-cavern.width())
	}
	if y < cavern.height()-1 {
		adjacent = append(adjacent, pos+cavern.width())
	}
	return adjacent
}

func (cavern *Cavern) width() int {
	return cavern.originalWidth * cavern.scale
}

func (cavern *Cavern) height() int {
	return cavern.originalHeight * cavern.scale
}

func (cavern *Cavern) riskAt(pos int) int {
	var x = pos % cavern.width()
	var y = pos / cavern.height()
	var originalX = x % cavern.originalWidth
	var originalY = y % cavern.originalHeight
	var xHops = (x - originalX) / cavern.originalWidth
	var yHops = (y - originalY) / cavern.originalHeight
	var originalRisk = cavern.risks[originalY*cavern.originalWidth+originalX]
	return (originalRisk-1+yHops+xHops)%9 + 1
}

func (cavern *Cavern) size() int {
	return cavern.height() * cavern.width()
}
