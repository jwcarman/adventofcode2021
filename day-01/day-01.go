package main

import (
	"adventofcode2021/input"
	"fmt"
	"strconv"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")

	var ints []int

	for _,line := range lines {
		i,_ := strconv.Atoi(line)
		ints = append(ints, i)
	}

	var count = 0

	for ndx,curr := range ints {
		if ndx > 0 {
			prev := ints[ndx-1]
			if curr > prev {
				count = count + 1
			}
		}
	}
	fmt.Println(count)

}


