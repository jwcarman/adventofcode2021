package main

import (
	"aoc2021/input"
	"aoc2021/util"
	"fmt"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	fmt.Println(calculatePart1(lines))
	fmt.Println(calculatePart2(lines))
}

func calculatePart2(lines []string) int {
	return countIncreases(toWindows(util.StringsToInts(lines)))
}

func calculatePart1(lines []string) int {
	return countIncreases(util.StringsToInts(lines))
}

func countIncreases(values []int) int {
	var count = 0

	for ndx, curr := range values {
		if ndx > 0 {
			prev := values[ndx-1]
			if curr > prev {
				count = count + 1
			}
		}
	}
	return count
}

func toWindows(depths []int) []int {
	var windows []int
	for ndx := 2; ndx < len(depths); ndx++ {

		windows = append(windows, depths[ndx]+depths[ndx-1]+depths[ndx-2])
	}
	return windows
}
