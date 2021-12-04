package main

import (
	"adventofcode2021/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := input.ReadLines("./input.txt")
	fmt.Println(calculatePart1(lines))
	fmt.Println(calculatePart2(lines))
}

func calculatePart1(lines []string) int {
	sub := Submarine{}
	moveSubmarine(lines, &sub)
	return sub.horizontalPosition * sub.depth
}

func calculatePart2(lines []string) int {
	sub := AttackSubmarine{}
	moveSubmarine(lines, &sub)
	return sub.horizontalPosition * sub.depth
}

func moveSubmarine(lines []string, controller SubmarineController) {
	for _, line := range lines {
		fields := strings.Fields(line)
		param, _ := strconv.Atoi(fields[1])

		switch fields[0] {
		case "up":
			controller.up(param)
		case "down":
			controller.down(param)
		case "forward":
			controller.forward(param)
		}
	}
}

type SubmarineController interface {
	up(param int)
	down(param int)
	forward(param int)
}

func (sub *Submarine) up(param int) {
	sub.depth -= param
}

func (sub *Submarine) down(param int) {
	sub.depth += param
}

func (sub *Submarine) forward(param int) {
	sub.horizontalPosition += param
}

func (sub *AttackSubmarine) up(param int) {
	sub.aim -= param
}

func (sub *AttackSubmarine) down(param int) {
	sub.aim += param
}

func (sub *AttackSubmarine) forward(param int) {
	sub.horizontalPosition += param
	sub.depth += sub.aim * param
}

type Submarine struct {
	horizontalPosition int
	depth              int
}

type AttackSubmarine struct {
	Submarine
	aim int
}
