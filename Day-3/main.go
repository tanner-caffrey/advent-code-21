package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/tanner-caffrey/aocutil"
)

func PartOne(input string) int {
	lines := strings.Split(input, "\n")
	bitLength := len(lines[0])
	bits := make([]int, bitLength)

	for _, line := range lines {
		for index, bit := range line {
			n, _ := strconv.Atoi(string(bit))
			if n == 0 {
				bits[index] -= 1
			} else {
				bits[index] += 1
			}
		}
	}

	gama := 0
	// Turn int slice into bits; negative numbers are 0's positive are 1's
	for index, bit := range bits {
		if bit > 0 {
			gama = gama | (1 << (bitLength - (index + 1))) // bitshift to correct digit
		}
	}

	upsilon := gama ^ (int(math.Pow(2, float64(bitLength))) - 1) // XOR with 1's to get least common bits

	return gama * upsilon
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please include cookie.")
	}
	cookie := os.Args[1]
	input, err := aocutil.GetInputFromDay(3, cookie)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part One Solution: ", PartOne(input))
}

//011000011101
//100111100010
