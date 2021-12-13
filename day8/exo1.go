package day8

import (
	"AdventOfCode2021/common"
	"fmt"
	"strings"
)

func Exo1() {
	lines := common.ReadLinesFromFile("day8/input.txt")

	sums := make(map[int]int, 0)
	for _, line := range lines {
		//leftEntries := make(map[string]int, 0)
		//rightEntries := make(map[string]int, 0)
		entries := strings.Split(line, "|")
		//leftStrArray := strings.Split(entries[0], " ")
		rightStrArray := strings.Split(entries[1], " ")

		//for _, leftEntry := range leftStrArray {
		//	//leftEntries[leftEntry] = len(leftEntry)
		//	sums[len(leftEntry)]++
		//}

		for _, rightEntry := range rightStrArray {
			//rightEntries[rightEntry] = len(rightEntry)
			sums[len(rightEntry)]++
		}

		//sums[1] += From(rightEntries).WhereT(func(k string, v int) bool { v == 2 }).SelectT(func (k string, v int) int { return kv.Value}).SumInts()

	}

	fmt.Println(sums[4] + sums[3] + sums[2] + sums[7])
}
