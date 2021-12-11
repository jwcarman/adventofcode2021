package main

import (
	"adventofcode2021/input"
	"fmt"
)

const n = 10

func main() {
	lines, _ := input.ReadLines("./input.txt")
	fmt.Println(calculatePart1(&lines))
	fmt.Println(calculatePart2(&lines))

}

func calculatePart1(lines *[]string) int {
	grid := parseGrid(*lines)
	var flashes = 0
	for step := 0; step < 100; step++ {
		incrementAll(&grid)
		flashThoseGreaterThanNine(&grid)
		flashes += resetFlashed(&grid)
	}
	return flashes
}

func calculatePart2(lines *[]string) int {
	grid := parseGrid(*lines)
	var flashes = 0
	var step = 0
	for flashes != 100 {
		step++
		incrementAll(&grid)
		flashThoseGreaterThanNine(&grid)
		flashes = resetFlashed(&grid)
	}
	return step
}

func parseGrid(lines []string) []int {
	var grid []int
	for _, line := range lines {
		for _, c := range line {
			grid = append(grid, int(c-'0'))
		}
	}
	return grid
}

func resetFlashed(grid *[]int) int {
	var flashes = 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			coord := coord(row, col)
			if (*grid)[coord] > 9 {
				(*grid)[coord] = 0
				flashes++
			}
		}
	}
	return flashes
}

func flashThoseGreaterThanNine(grid *[]int) {
	var flashed = make([]bool, n*n)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if (*grid)[coord(row, col)] > 9 {
				flash(grid, &flashed, row, col)
			}
		}
	}
}

func incrementAll(grid *[]int) {
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			(*grid)[coord(row, col)]++
		}
	}
}

func flash(grid *[]int, flashed *[]bool, row int, col int) {
	if (*flashed)[coord(row, col)] {
		return
	}
	(*flashed)[coord(row, col)] = true

	for x := row - 1; x <= row+1; x++ {
		for y := col - 1; y <= col+1; y++ {
			if x >= 0 && x < n && y >= 0 && y < n {
				coord := coord(x, y)

				(*grid)[coord]++
				if (*grid)[coord] > 9 {
					flash(grid, flashed, x, y)
				}
			}
		}
	}
}
func coord(row int, col int) int {
	return row*n + col
}
