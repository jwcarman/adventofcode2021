package main

import (
	"adventofcode2021/input"
	"fmt"
	"sort"
)

var pairs = map[int32]int32{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
}

func main() {
	lines, _ := input.ReadLines("./input.txt")
	fmt.Println(calculatePart1(lines))
	fmt.Println(calculatePart2(lines))
}

func calculatePart2(lines []string) int {
	var scores []int
	for _, line := range lines {
		stack, err := checkLine(line)
		if err == 0 {
			var score = 0
			for !stack.IsEmpty() {
				score = (score * 5) + autoCorrectScoreOf(pairs[stack.Pop()])
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func calculatePart1(lines []string) int {
	var sum = 0
	for _, line := range lines {
		_, err := checkLine(line)
		sum += err
	}
	return sum
}

func checkLine(line string) (Stack, int) {
	var stack Stack
	for _, char := range line {
		if _, ok := pairs[char]; ok {
			stack.Push(char)
		} else {
			if pairs[stack.Peek()] == char {
				stack.Pop()
			} else {
				return stack, syntaxErrorScoreOf(char)
			}
		}
	}
	return stack, 0
}

func autoCorrectScoreOf(c int32) int {
	switch c {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}

func syntaxErrorScoreOf(c int32) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

type Stack []int32

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() int32 {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(top int32) {
	*s = append(*s, top)
}

func (s *Stack) Pop() int32 {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
