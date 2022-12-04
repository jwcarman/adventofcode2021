package main

import (
	"aoc2021/input"
	"aoc2021/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")

	var points = make(map[Point2D]int)
	for _, line := range lines {
		if strings.Contains(line, ",") {
			var splits = strings.Split(line, ",")
			var coords = util.StringsToInts(splits)
			var point = Point2D{x: coords[0], y: coords[1]}
			points[point]++
		}
		if strings.Contains(line, "fold along x=") {
			x, _ := strconv.Atoi(strings.TrimPrefix(line, "fold along x="))
			for oldPoint, count := range points {
				var newPoint = oldPoint.foldX(x)
				delete(points, oldPoint)
				points[newPoint] += count
			}
		}
		if strings.Contains(line, "fold along y=") {
			y, _ := strconv.Atoi(strings.TrimPrefix(line, "fold along y="))
			for oldPoint, count := range points {
				var newPoint = oldPoint.foldY(y)
				delete(points, oldPoint)
				points[newPoint] += count
			}
		}
	}

	var y = 0
	var xs = selectXs(y, points)
	for len(xs) > 0 {
		var maxX = xs[len(xs)-1]
		for x := 0; x <= maxX; x++ {
			if xs[0] == x {
				fmt.Print("#")
				xs = xs[1:]
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
		y++
		xs = selectXs(y, points)
	}
}

func selectXs(y int, points map[Point2D]int) []int {
	var xs []int
	for point := range points {
		if point.y == y {
			xs = append(xs, point.x)
		}
	}
	sort.Ints(xs)
	return xs
}

type Point2D struct {
	x, y int
}

func (point Point2D) foldX(x int) Point2D {
	if point.x > x {
		return Point2D{x: 2*x - point.x, y: point.y}
	}
	return point
}

func (point Point2D) foldY(y int) Point2D {
	if point.y > y {
		return Point2D{x: point.x, y: 2*y - point.y}
	}
	return point
}
