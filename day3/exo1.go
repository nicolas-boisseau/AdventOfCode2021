package day3

import (
	"AdventOfCode2021/common"
	"fmt"
	"strconv"
)

func Exo1() {
	lines := common.ReadLinesFromFile("day3/input.txt")

	ones := make([]int, len(lines[0]))
	zeros := make([]int, len(lines[0]))
	for _, line := range lines {
		for j, c := range line {
			if c == '1' {
				ones[j]++
			} else {
				zeros[j]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i, _ := range ones {
		if ones[i] > zeros[i] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaNumber, _ := strconv.ParseInt(gamma, 2, 0)
	epsilonNumber, _ := strconv.ParseInt(epsilon, 2, 0)
	fmt.Println(gammaNumber)
	fmt.Println(epsilonNumber)
	fmt.Println("Power consumption:", gammaNumber*epsilonNumber)
}
