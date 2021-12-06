package main

import (
	"adventofcode2021/input"
	"adventofcode2021/util"
	"fmt"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	var segments = parseSegments(&lines)
	fmt.Println(calculatePart1(&segments))
	fmt.Println(calculatePart2(&segments))
}

func calculatePart1(segments *[]LineSegment2D) int {
	var pointCounts = make(map[Point2D]int)
	for _, segment := range *segments {
		if segment.isHorizontal() || segment.isVertical() {
			for _, point := range segment.points() {
				pointCounts[point]++
			}
		}
	}
	var count = 0
	for _, n := range pointCounts {
		if n > 1 {
			count++
		}
	}
	return count
}

func calculatePart2(segments *[]LineSegment2D) int {
	var pointCounts = make(map[Point2D]int)
	for _, segment := range *segments {
		for _, point := range segment.points() {
			pointCounts[point]++
		}
	}
	var count = 0
	for _, n := range pointCounts {
		if n > 1 {
			count++
		}
	}
	return count
}

func parseSegments(lines *[]string) []LineSegment2D {
	var segments []LineSegment2D
	for _, line := range *lines {
		segments = append(segments, parseSegment(line))
	}
	return segments
}

func parseSegment(line string) LineSegment2D {
	var coords = util.StringsToInts(strings.Fields(strings.ReplaceAll(strings.ReplaceAll(line, " -> ", " "), ",", " ")))
	return LineSegment2D{a: Point2D{x: coords[0], y: coords[1]}, b: Point2D{x: coords[2], y: coords[3]}}
}

type Point2D struct {
	x, y int
}

type LineSegment2D struct {
	a, b Point2D
}

func max(left int, right int) int {
	if left >= right {
		return left
	}
	return right
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func d(beg int, end int) int {
	if beg == end {
		return 0
	}
	if beg > end {
		return -1
	}
	return 1
}

func (segment LineSegment2D) points() []Point2D {
	var n = max(abs(segment.a.x-segment.b.x), abs(segment.a.y-segment.b.y)) + 1
	var dx = d(segment.a.x, segment.b.x)
	var dy = d(segment.a.y, segment.b.y)

	var points []Point2D

	for i := 0; i < n; i++ {
		points = append(points, Point2D{x: segment.a.x + i*dx, y: segment.a.y + i*dy})
	}
	return points
}

func (segment LineSegment2D) isVertical() bool {
	return segment.a.x == segment.b.x
}

func (segment LineSegment2D) isHorizontal() bool {
	return segment.a.y == segment.b.y
}
