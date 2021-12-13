package day2

import (
	"AdventOfCode2021/common"
	"fmt"
	"strconv"
	"strings"
)

func Exo2() {

	lines := common.ReadLinesFromFile("day2/input.txt")

	var hpos int64 = 0
	var depth int64 = 0
	var aim int64 = 0
	for _, line := range lines {
		instructions := strings.Split(line, " ")
		direction := instructions[0]
		distance, _ := strconv.ParseInt(instructions[1], 10, 0)

		if direction == "forward" {
			hpos += distance
			depth += aim * distance
		} else if direction == "down" {
			aim += distance
		} else {
			aim -= distance
		}
	}

	fmt.Println("hpos:", hpos, "depth:", depth)
	fmt.Println("Total:", hpos*depth)
}
