package main

import (
	"adventofcode2021/input"
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	fmt.Println(calculatePart1(lines))
	fmt.Println(calculatePart2(lines))
}

func calculatePart2(lines []string) int {
	var total = 0
	for _, line := range lines {
		total += processLine(line)
	}
	return total
}

func processLine(line string) int {
	splits := strings.Split(line, " | ")
	var patterns = parseSignalPatterns(splits[0])
	alphabet := createInitialAlphabet(patterns)

	for _, pattern := range patterns {
		var size = pattern.size()
		switch size {
		case 5:
			var diff = alphabet[4].subtract(pattern)
			if diff.size() == 2 {
				alphabet[2] = pattern
			} else {
				diff = diff.subtract(alphabet[1])
				if diff.size() == 0 {
					alphabet[5] = pattern
				} else {
					alphabet[3] = pattern
				}
			}
		case 6:
			var diff = pattern.subtract(alphabet[1])
			if diff.size() == 5 {
				alphabet[6] = pattern
			} else {
				diff = diff.subtract(alphabet[4])
				if diff.size() == 2 {
					alphabet[9] = pattern
				} else {
					alphabet[0] = pattern
				}
			}
		}
	}
	return decode(alphabet, splits[1])
}

func decode(alphabet map[int]SignalPattern, output string) int {
	var value = 0
	for ndx, digit := range strings.Fields(output) {
		var pattern = createSignalPattern(digit)
		for k, v := range alphabet {
			if reflect.DeepEqual(v.segments, pattern.segments) {
				value += int(math.Pow10(3-ndx)) * k
			}
		}
	}
	return value
}

func createInitialAlphabet(patterns []SignalPattern) map[int]SignalPattern {
	var alphabet = make(map[int]SignalPattern)
	for _, pattern := range patterns {
		var size = pattern.size()
		switch size {
		case 2:
			alphabet[1] = pattern
		case 3:
			alphabet[7] = pattern
		case 4:
			alphabet[4] = pattern
		case 7:
			alphabet[8] = pattern
		}
	}
	return alphabet
}

func parseSignalPatterns(s string) []SignalPattern {
	var patterns []SignalPattern
	for _, pattern := range strings.Fields(s) {
		patterns = append(patterns, createSignalPattern(pattern))
	}
	return patterns
}

func calculatePart1(lines []string) int {
	var count = 0

	for _, line := range lines {
		digits := strings.Fields(strings.Split(line, " | ")[1])
		for _, digit := range digits {
			segments := len(digit)
			if segments == 2 || segments == 3 || segments == 4 || segments == 7 {
				count++
			}
		}
	}
	return count
}

func createSignalPattern(digit string) SignalPattern {
	var segments = make(map[int32]bool)
	for _, c := range digit {
		segments[c] = true
	}
	return SignalPattern{segments: segments}
}

type SignalPattern struct {
	segments map[int32]bool
}

func (p SignalPattern) size() int {
	return len(p.segments)
}

func (p SignalPattern) subtract(right SignalPattern) SignalPattern {
	var diff = make(map[int32]bool)
	for k := range p.segments {
		diff[k] = true
	}
	for k := range right.segments {
		delete(diff, k)
	}
	return SignalPattern{segments: diff}
}
