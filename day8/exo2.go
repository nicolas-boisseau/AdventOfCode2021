package day8

import (
	"AdventOfCode2021/common"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func findCommonChars(a string, b string) ([]rune, int) {
	commonChars := make([]rune, 0)
	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				commonChars = append(commonChars, c1)
			}
		}
	}
	return commonChars, len(commonChars)
}

func pass1(str string) int {
	switch {
	case len(str) == 2:
		return 1
	case len(str) == 7:
		return 8
	case len(str) == 3:
		return 7
	case len(str) == 4:
		return 4
	}
	return -1
}

func pass2(str string, digitStrings map[int]string) int {
	// déjà 1, 4, 7, 8

	// 9 => 6, dont 2 commun avec 1
	// 3 => 5 dont 2 commun avec 1, 3 commun avec 7
	// 6 => 6 mais 1 seul commun avec 1
	// 0 => 6 dont 2 commun avec 1

	_, n4 := findCommonChars(str, digitStrings[4])
	_, n1 := findCommonChars(str, digitStrings[1])
	_, n7 := findCommonChars(str, digitStrings[7])

	if len(str) == 6 && n4 == 3 && n1 == 2 {
		return 0
	} else if len(str) == 6 && n4 == 4 {
		return 9
	} else if len(str) == 6 && n1 == 1 {
		return 6
	} else if len(str) == 5 && n1 == 2 && n7 == 3 {
		return 3
	}

	return -1
}

func pass3(str string, digitStrings map[int]string) int {

	// 5 => 5, dont 3 commun avec 4, 2 commun avec 1 mais pas les mêmes
	// 2 => 5 mais 1 seul commun avec 1 != de celui commun avec le 5, (le même que 6)

	common_6_1, _ := findCommonChars(digitStrings[6], digitStrings[1])
	common_with_1, _ := findCommonChars(str, digitStrings[1])

	if len(str) == 5 && common_6_1[0] == common_with_1[0] {
		return 5
	} else {
		return 2
	}

	return -1
}

func Exo2() {
	lines := common.ReadLinesFromFile("day8/input.txt")

	//sums := make(map[int]int, 0)
	sum := 0
	for _, line := range lines {
		leftEntries := make(map[string]int, 0)
		digitStrings := make(map[int]string, 0)

		//rightEntries := make(map[string]int, 0)
		entries := strings.Split(line, "|")
		leftStrArray := strings.Split(entries[0], " ")
		rightStrArray := strings.Split(entries[1], " ")

		// sort strings
		for i := range leftStrArray {
			leftStrArray[i] = SortString(leftStrArray[i])
		}
		for i := range rightStrArray {
			rightStrArray[i] = SortString(rightStrArray[i])
		}

		// Pass 1
		for _, leftEntry := range leftStrArray {
			leftEntries[leftEntry] = pass1(leftEntry)
			if leftEntries[leftEntry] != -1 {
				digitStrings[leftEntries[leftEntry]] = leftEntry
			}
		}

		fmt.Println("PASS 1 :", leftEntries)

		// Pass 2
		for _, leftEntry := range leftStrArray {
			// si pas déjà trouvé dans pass 1
			if leftEntries[leftEntry] == -1 {
				leftEntries[leftEntry] = pass2(leftEntry, digitStrings)
				digitStrings[leftEntries[leftEntry]] = leftEntry
			}
		}

		fmt.Println("PASS 2 :", leftEntries)

		// Pass 3
		for _, leftEntry := range leftStrArray {
			// si pas déjà trouvé dans pass 1
			if leftEntries[leftEntry] == -1 {
				leftEntries[leftEntry] = pass3(leftEntry, digitStrings)
			}
		}

		fmt.Println("PASS 3 :", leftEntries)

		//for _, rightEntry := range rightStrArray {
		//	//rightEntries[rightEntry] = len(rightEntry)
		//	//sums[len(rightEntry)]++
		//}

		//fmt.Println(leftEntries[rightStrArray[1]], leftEntries[rightStrArray[2]], leftEntries[rightStrArray[3]], leftEntries[rightStrArray[4]])
		buf := bytes.NewBufferString("")
		fmt.Fprintf(buf, "%d%d%d%d", leftEntries[rightStrArray[1]], leftEntries[rightStrArray[2]], leftEntries[rightStrArray[3]], leftEntries[rightStrArray[4]])
		//fmt.Println(buf.String())
		n, _ := strconv.Atoi(buf.String())
		//fmt.Println(buf.String())
		sum += n
		//sums[1] += From(rightEntries).WhereT(func(k string, v int) bool { v == 2 }).SelectT(func (k string, v int) int { return kv.Value}).SumInts()

		//break
	}
	fmt.Println(sum)
	//fmt.Println(sums[4]+sums[3]+sums[2]+sums[7])
}
