package main

import (
	"aoc2021/input"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	var polymer = lines[0]
	var rules = make(map[string]string)
	for _, line := range lines[2:] {
		var splits = strings.Split(line, " -> ")
		rules[splits[0]] = splits[1]
	}

	fmt.Println(calculatePart1(polymer, rules))
	fmt.Println(calculatePart2(polymer, rules))
}

func calculatePart2(polymer string, rules map[string]string) int {
	return calculateAnswer(polymer, rules, 40)
}

func calculatePart1(polymer string, rules map[string]string) int {
	return calculateAnswer(polymer, rules, 10)
}

func calculateAnswer(template string, rules map[string]string, rounds int) int {
	var pairCounts = make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairCounts[string(template[i])+string(template[i+1])]++
	}

	for round := 0; round < rounds; round++ {
		pairCounts = performRound(&pairCounts, &rules)
	}

	counts := characterCounts(&pairCounts)
	counts[string(template[len(template)-1])]++

	min, max := minAndMaxCounts(counts)
	return max - min
}

func performRound(pairCounts *map[string]int, rules *map[string]string) map[string]int {
	newPairCounts := make(map[string]int)
	for pair, count := range *pairCounts {
		replacement := (*rules)[pair]
		newPairCounts[string(pair[0])+replacement] += count
		newPairCounts[replacement+string(pair[1])] += count
	}
	return newPairCounts
}

func minAndMaxCounts(counts map[string]int) (int, int) {
	var min = math.MaxInt
	var max = 0
	for _, count := range counts {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return min, max
}

func characterCounts(pairCounts *map[string]int) map[string]int {
	var counts = make(map[string]int)
	for pair, count := range *pairCounts {
		counts[string(pair[0])] += count
	}
	return counts
}
