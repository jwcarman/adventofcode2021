package main

import (
	"adventofcode2021/input"
	"fmt"
	"strconv"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	fmt.Println(calculatePart1(lines))
	fmt.Println(calculatePart2(lines))
}

func calculatePart2(lines []string) int {
	return countIncreases(toWindows(toDepths(lines)))
}

func calculatePart1(lines []string) int {
	return countIncreases(toDepths(lines))
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
	for ndx, depth := range depths {
		if ndx >= 2 {
			windows = append(windows, depth+depths[ndx-1]+depths[ndx-2])
		}
	}
	return windows
}

func toDepths(lines []string) []int {
	var depths []int
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		depths = append(depths, i)
	}
	return depths
}
