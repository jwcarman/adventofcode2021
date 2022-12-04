package main

import (
	"aoc2021/input"
	"aoc2021/util"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	var positions = util.StringsToInts(strings.Split(lines[0], ","))
	fmt.Println(calculatePart1(positions))
	fmt.Println(calculatePart2(positions))
}

func calculatePart1(positions []int) int {
	var median = medianOf(positions)
	var totalFuel = 0
	for _, position := range positions {
		totalFuel += abs(position - median)
	}
	return totalFuel
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func calculatePart2(positions []int) int {
	var mean = meanOf(positions)
	var totalFuel = 0
	for _, position := range positions {
		var dx = abs(position - mean)

		totalFuel += (dx * (dx + 1)) / 2
	}
	return totalFuel
}

func meanOf(values []int) int {
	var sum = 0.0
	for _, v := range values {
		sum += float64(v)
	}
	return int(math.Floor(sum / float64(len(values))))
}

func medianOf(values []int) int {
	sort.Ints(values)
	return values[len(values)/2]
}
