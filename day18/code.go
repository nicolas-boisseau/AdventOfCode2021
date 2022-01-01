package day18

import (
	"AdventOfCode2021/common"
)

func Process(fileName string) int {
	lines := common.ReadLinesFromFile(fileName)

	n := ReadNode(lines[0])
	for _, line := range lines[1:] {
		n2 := ReadNode(line)
		n = Add(n, n2)
	}

	return int(n.Magnitude())
}

func Process2(fileName string) int {
	lines := common.ReadLinesFromFile(fileName)

	var maxMagnitude int64 = 0

	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j {
				continue
			}
			n1 := ReadNode(line1)
			n2 := ReadNode(line2)
			addition := Add(n1, n2)
			magnitude := addition.Magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
			n1 = ReadNode(line1)
			n2 = ReadNode(line2)
			addition = Add(n2, n1)
			magnitude = addition.Magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}

	return int(maxMagnitude)
}
