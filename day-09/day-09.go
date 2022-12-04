package main

import (
	"aoc2021/input"
	"fmt"
	"sort"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")

	var heightmap = parseHeightmap(&lines)
	fmt.Println(calculatePart1(&heightmap))
	fmt.Println(calculatePart2(&heightmap))
}

func calculatePart2(heightmap *[][]int) int {
	var rows = len(*heightmap)
	var cols = len((*heightmap)[0])
	var basinSizes []int
	var visited = createVisited(rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var value = (*heightmap)[row][col]
			var neighbors = neighborsOf(heightmap, row, col)
			if isLowPoint(value, neighbors) {
				visited[row][col] = true
				var basinSize = 1 +
					flowFrom(heightmap, &visited, row-1, col, value) +
					flowFrom(heightmap, &visited, row+1, col, value) +
					flowFrom(heightmap, &visited, row, col-1, value) +
					flowFrom(heightmap, &visited, row, col+1, value)
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
}

func createVisited(rows int, cols int) [][]bool {
	var visited [][]bool
	for row := 0; row < rows; row++ {
		visited = append(visited, make([]bool, cols))
	}
	return visited
}

func flowFrom(heightmap *[][]int, visited *[][]bool, row int, col int, lowPoint int) int {
	if row < 0 || row >= len(*heightmap) || col < 0 || col >= len((*heightmap)[0]) || (*visited)[row][col] {
		return 0
	}
	var value = (*heightmap)[row][col]
	if value == 9 || value <= lowPoint {
		return 0
	}
	(*visited)[row][col] = true
	return 1 +
		flowFrom(heightmap, visited, row-1, col, value) +
		flowFrom(heightmap, visited, row+1, col, value) +
		flowFrom(heightmap, visited, row, col-1, value) +
		flowFrom(heightmap, visited, row, col+1, value)
}

func calculatePart1(heightmap *[][]int) int {
	var n = len(*heightmap)
	var sum = 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			var value = (*heightmap)[row][col]
			var neighbors = neighborsOf(heightmap, row, col)
			if isLowPoint(value, neighbors) {
				sum += value + 1
			}
		}
	}
	return sum
}

func isLowPoint(value int, neighbors []int) bool {
	for _, neighbor := range neighbors {
		if value >= neighbor {
			return false
		}
	}
	return true
}
func neighborsOf(heightmap *[][]int, row int, col int) []int {
	var neighbors []int
	if row > 0 {
		neighbors = append(neighbors, (*heightmap)[row-1][col])
	}
	if row < len(*heightmap)-1 {
		neighbors = append(neighbors, (*heightmap)[row+1][col])
	}
	if col > 0 {
		neighbors = append(neighbors, (*heightmap)[row][col-1])
	}
	if col < len((*heightmap)[0])-1 {
		neighbors = append(neighbors, (*heightmap)[row][col+1])
	}
	return neighbors
}

func parseHeightmap(lines *[]string) [][]int {
	var heightmap [][]int
	for _, line := range *lines {
		heightmap = append(heightmap, parseRow(line))
	}
	return heightmap
}

func parseRow(line string) []int {
	var row []int
	for _, c := range line {
		row = append(row, int(c-'0'))
	}
	return row
}
