package day10

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/golang-collections/collections/stack"
)

func Parse2(line string) (bool, rune, string) {
	s := stack.New()

	m := make(map[string]string)
	m[")"] = "("
	m["]"] = "["
	m["}"] = "{"
	m[">"] = "<"
	m["("] = ")"
	m["["] = "]"
	m["{"] = "}"
	m["<"] = ">"

	output := ""
	for _, c := range line {
		output += string(c)
		//fmt.Println(output)
		//if s.Peek() != nil {
		//	closing := s.Peek().(string)
		//	//fmt.Println("current Peek=", closing)
		//	//fmt.Println(m)
		//	//fmt.Println(closing, " =? ", m[string(c)])
		//
		//}

		if s.Peek() != nil && s.Peek().(string) == m[string(c)] {
			s.Pop()
		} else if string(c) == "(" || string(c) == "[" || string(c) == "{" || string(c) == "<" {
			s.Push(string(c))
		} else {
			return false, c, ""
		}
	}

	toComplete := ""
	for s.Len() > 0 {
		toComplete += m[s.Pop().(string)]
	}

	return true, 'E', toComplete
}

func Exo2() {
	lines := common.ReadLinesFromFile("day10/input.txt")

	noErrorLines := make([]string, 0)
	for _, line := range lines {
		success, c, _ := Parse2(line)
		if !success {
			errorStr := string(c)
			fmt.Println("Error :", errorStr, "not expected.")
		} else {
			noErrorLines = append(noErrorLines, line)
		}
	}

	fmt.Println("Remaining lines:")
	sums := make([]int, 0)
	for _, line := range noErrorLines {
		fmt.Println(line)
		_, _, toComplete := Parse2(line)

		fmt.Println("ToComplete = ", toComplete)

		subSum := 0
		for _, c := range toComplete {
			subSum *= 5
			switch string(c) {
			case ")":
				subSum += 1
			case "]":
				subSum += 2
			case "}":
				subSum += 3
			case ">":
				subSum += 4
			}
		}

		fmt.Println("SubScore = ", subSum)

		sums = append(sums, subSum)
	}

	sortedSums := make([]int, 0)
	From(sums).OrderByT(func(a int) int { return a }).ToSlice(&sortedSums)

	fmt.Println("score=", sortedSums[len(sortedSums)/2])
}
