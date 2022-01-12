package day20

import (
	"AdventOfCode2021/common"
	"bytes"
	"fmt"
	"strconv"
)

const Bonus = 100

func Process(fileName string, enhancePassCount int) int {
	lines := common.ReadLinesFromFile(fileName)

	rulesStr := lines[0]
	rules := make([]int, 0)
	for _, c := range rulesStr {
		toInt := 0
		if string(c) == "#" {
			toInt = 1
		}
		rules = append(rules, toInt)
	}

	inputSize := len(lines[2])
	size := inputSize + Bonus*2

	g := newGrille(size, size)
	for i := 0; i < inputSize; i++ {
		for j := 0; j < inputSize; j++ {
			if string(lines[2+i][j]) == "#" {
				g.content[i+Bonus][j+Bonus] = 1
			}
		}
	}

	for i := 0; i < enhancePassCount; i++ {
		g = g.Enhance(rules)
	}

	result := g.NumberOFLightPixels()
	fmt.Println(result)

	return result
}

func (g *Grille) Enhance(rules []int) *Grille {
	outputG := newGrille(g.h, g.w)

	if rules[0] == 1 && rules[511] == 0 {
		outputG.defaultValue = 1 - g.defaultValue
	}

	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			//if g.content[i][j] == 1 {
			//	g.content[i][j] = 1
			//}
			outputG.content[i][j] = rules[g.BinaryNumberOf(i, j)]
		}
	}

	//fmt.Println(outputG)
	//fmt.Println("Light pixels: ", outputG.NumberOFLightPixels())

	return outputG
}

func (g *Grille) NumberOFLightPixels() int {
	sum := 0
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			sum += g.content[i][j]
		}
	}
	return sum
}

func (g *Grille) BinaryNumberOf(i, j int) int {
	getValueAt := func(i, j int) int {
		if i < 0 || i >= g.h || j < 0 || j >= g.w {
			return g.defaultValue
		}
		return g.content[i][j]
	}

	buff := bytes.NewBufferString("")
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i-1, j-1)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i-1, j)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i-1, j+1)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i, j-1)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i, j)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i, j+1)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i+1, j-1)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i+1, j)))
	fmt.Fprint(buff, strconv.Itoa(getValueAt(i+1, j+1)))

	binaryStr := buff.String()
	intVal, _ := strconv.ParseInt(binaryStr, 2, 32)

	return int(intVal)
}
