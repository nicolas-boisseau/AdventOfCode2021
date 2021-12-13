package day10

import (
	"AdventOfCode2021/common"
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func Parse(line string) (bool, rune) {
	s := stack.New()

	m := make(map[string]string)
	m[")"] = "("
	m["]"] = "["
	m["}"] = "{"
	m[">"] = "<"

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
			return false, c
		}
	}

	return true, 'E'
}

func Exo1(fileName string) {
	lines := common.ReadLinesFromFile(fileName)

	sum := 0
	for _, line := range lines {
		//fmt.Println(line)

		success, c := Parse(line)
		if !success {
			errorStr := string(c)
			fmt.Println("Error :", errorStr, "not expected.")
			switch errorStr {
			case ")":
				sum += 3
			case "]":
				sum += 57
			case "}":
				sum += 1197
			case ">":
				sum += 25137
			}
		}

	}
	fmt.Println("score=", sum)
}
