package main

import (
	"adventofcode2021/input"
	"adventofcode2021/util"
	"fmt"
	"strings"
)

const boardSize = 5

func main() {
	lines, _ := input.ReadLines("./input.txt")

	numbers := util.StringsToInts(strings.Split(lines[0], ","))
	boards := readBoards(&lines)

	fmt.Println(calculatePart1(numbers, boards))
	fmt.Println(calculatePart2(numbers, boards))
}

func calculatePart2(numbers []int, boards []BingoBoard) int {
	var winners []BingoBoard
	for _, number := range numbers {
		for ndx := range boards {
			var board = &boards[ndx]
			if board.mark(number) && board.score == -1 {
				board.score = board.sumOfUnmarked() * number
				winners = append(winners, *board)
			}
		}
	}
	return winners[len(winners)-1].score
}

func calculatePart1(numbers []int, boards []BingoBoard) int {
	for _, number := range numbers {
		for _, board := range boards {
			if board.mark(number) {
				return board.sumOfUnmarked() * number
			}
		}
	}
	return -1
}

func readBoards(lines *[]string) []BingoBoard {
	var boards []BingoBoard
	for begin := 2; begin < len(*lines); begin += boardSize + 1 {
		var spaces []BingoBoardSpace
		for ndx := begin; ndx < begin+boardSize; ndx++ {
			for _, number := range util.StringsToInts(strings.Fields((*lines)[ndx])) {
				spaces = append(spaces, BingoBoardSpace{number: number, marked: false})
			}
		}
		boards = append(boards, BingoBoard{spaces: spaces, score: -1})
	}
	return boards
}

type BingoBoardSpace struct {
	number int
	marked bool
}

type BingoBoard struct {
	spaces []BingoBoardSpace
	score  int
}

func (board *BingoBoard) mark(number int) bool {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			space := board.spaceAt(row, col)
			if space.number == number {
				space.marked = true
				return board.checkRowForWin(row) || board.checkColForWin(col)
			}
		}
	}

	return false
}

func (board *BingoBoard) sumOfUnmarked() int {
	var sum = 0
	for _, space := range board.spaces {
		if !space.marked {
			sum += space.number
		}
	}
	return sum
}

func (board *BingoBoard) spaceAt(row int, col int) *BingoBoardSpace {
	return &board.spaces[row*boardSize+col]
}

func (board *BingoBoard) checkColForWin(col int) bool {
	for row := 0; row < boardSize; row++ {
		if !board.spaceAt(row, col).marked {
			return false
		}
	}
	return true
}

func (board *BingoBoard) checkRowForWin(row int) bool {
	for col := 0; col < boardSize; col++ {
		if !board.spaceAt(row, col).marked {
			return false
		}
	}
	return true
}
