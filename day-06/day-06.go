package main

import (
	"aoc2021/input"
	"aoc2021/util"
	"fmt"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	var timers = util.StringsToInts(strings.Split(lines[0], ","))

	fmt.Println(calculatePart1(timers))
	fmt.Println(calculatePart2(timers))
}

func calculatePart1(timers []int) int {
	return countTotalDescendants(timers, 80)
}

func calculatePart2(timers []int) int {
	return countTotalDescendants(timers, 256)
}

func countTotalDescendants(timers []int, initialGenerations int) int {
	var count = 0
	var countCache = make(map[Lanternfish]int)

	for _, timer := range timers {
		count += countDescendents(Lanternfish{timer: timer, generations: initialGenerations}, &countCache)
	}
	return count
}

func countDescendents(fish Lanternfish, cacheMap *map[Lanternfish]int) int {
	if val, ok := (*cacheMap)[fish]; ok {
		return val
	}
	var count = 1
	for i := fish.generations - fish.timer - 1; i >= 0; i -= 7 {
		count += countDescendents(Lanternfish{timer: 6, generations: i - 2}, cacheMap)
	}
	(*cacheMap)[fish] = count
	return count
}

type Lanternfish struct {
	timer       int
	generations int
}
