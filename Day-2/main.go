package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tanner-caffrey/aocutil"
)

type Submarine struct {
	horizontal, depth, aim int
}

func (i *Submarine) Rise(n int) {
	i.aim -= n
}

func (i *Submarine) Sink(n int) {
	i.aim += n
}

func (i *Submarine) Straight(n int) {
	i.horizontal += n
	i.depth += i.aim * n
}

func Navigate(input string) Submarine {
	sub := Submarine{0, 0, 0}

	directions := strings.Split(input[:len(input)-1], "\n")
	for _, direction := range directions {
		direction := strings.Split(direction, " ")
		unit, _ := strconv.Atoi(direction[1])
		dir := direction[0]
		if dir == "up" {
			sub.Rise(unit)
		} else if dir == "down" {
			sub.Sink(unit)
		} else if dir == "forward" {
			sub.Straight(unit)
		}
	}

	return sub
}

func PartOne(sub Submarine) int {
	return sub.aim * sub.horizontal
}

func PartTwo(sub Submarine) int {
	return sub.depth * sub.horizontal
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please include session cookie.")
		os.Exit(1)
	}
	cookie := os.Args[1]

	input, err := aocutil.GetInputFromDay(2, cookie)
	if err != nil {
		fmt.Println("[X] ERROR: ", err)
	}

	sub := Navigate(input)

	fmt.Println("Part One Solution: ", PartOne(sub))

	fmt.Println("Part Two Solution: ", PartTwo(sub))
}
