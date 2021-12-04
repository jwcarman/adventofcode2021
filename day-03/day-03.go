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

const MaximumLineLength = 12

func calculatePart2(lines []string) int64 {
	return oxygenGeneratorRatingOf(lines) * c02ScrubberRating(lines)
}

func oxygenGeneratorRatingOf(lines []string) int64 {
	var currentLines = lines
	for ndx := 0; ndx < MaximumLineLength; ndx++ {
		if len(currentLines) == 1 {
			return decimalValueOf(currentLines[0])
		}
		var sum = sumOfValuesAt(currentLines, ndx)
		if sum >= 0 {
			currentLines = keepWithDigitAt(currentLines, ndx, '1')
		} else {
			currentLines = keepWithDigitAt(currentLines, ndx, '0')
		}
	}
	return -1
}

func c02ScrubberRating(lines []string) int64 {
	var currentLines = lines
	for ndx := 0; ndx < MaximumLineLength; ndx++ {
		if len(currentLines) == 1 {
			return decimalValueOf(currentLines[0])
		}
		var sum = sumOfValuesAt(currentLines, ndx)
		if sum < 0 {
			currentLines = keepWithDigitAt(currentLines, ndx, '1')
		} else {
			currentLines = keepWithDigitAt(currentLines, ndx, '0')
		}
	}
	return -1
}

func keepWithDigitAt(lines []string, ndx int, digit uint8) []string {
	var validLines []string
	for _, line := range lines {
		if line[ndx] == digit {
			validLines = append(validLines, line)
		}
	}
	return validLines
}

func calculatePart1(lines []string) int64 {
	var gammaRate []uint8

	for ndx := 0; ndx < MaximumLineLength; ndx++ {
		gammaRate = append(gammaRate, mostCommonDigitAt(lines, ndx))
	}
	var epsilonRate = flipDigits(gammaRate)
	return decimalValueOf(string(gammaRate)) * decimalValueOf(string(epsilonRate))
}

func decimalValueOf(binaryString string) int64 {
	dec, _ := strconv.ParseInt(binaryString, 2, 0)
	return dec
}

func flipDigit(digit uint8) uint8 {
	if digit == '1' {
		return '0'
	}
	return '1'
}

func flipDigits(digits []uint8) []uint8 {
	var flipped []uint8
	for _, digit := range digits {
		flipped = append(flipped, flipDigit(digit))
	}
	return flipped
}

func sumOfValuesAt(lines []string, ndx int) int {
	var sum = 0
	for _, line := range lines {
		sum += valueOfDigitAt(line, ndx)
	}
	return sum
}

func mostCommonDigitAt(lines []string, ndx int) uint8 {
	var sum = 0
	for _, line := range lines {
		sum += valueOfDigitAt(line, ndx)
	}
	if sum > 0 {
		return '1'
	} else {
		return '0'
	}
}

func valueOfDigitAt(line string, ndx int) int {
	if line[ndx] == '1' {
		return 1
	} else {
		return -1
	}
}
