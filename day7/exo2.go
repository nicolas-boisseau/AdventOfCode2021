package day7

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"log"
	"math"
	"strconv"
	"strings"
)

// 	. "github.com/ahmetalpbalkan/go-linq"

func Exo2() {
	lines := common.ReadLinesFromFile("day7/input.txt")

	crabsPositionsStr := strings.Split(lines[0], ",")
	crabsPositions := make([]int, 0)
	for _, crabPosStr := range crabsPositionsStr {
		pos, err := strconv.Atoi(crabPosStr)
		if err != nil {
			log.Fatal(err)
		}
		crabsPositions = append(crabsPositions, pos)
	}

	// compute cost for each pos
	maxPos := From(crabsPositions).Max().(int)
	costsByPos := make(map[int]int)
	for i := 1; i <= maxPos; i++ {
		costsByPos[i] = 0
	}
	for _, crabPos := range crabsPositions {
		for k, _ := range costsByPos {
			diff := float64(crabPos - k)
			sum := 0
			for i := 1; i <= int(math.Abs(diff)); i++ {
				sum += i
			}
			costsByPos[k] += sum
		}
	}

	fmt.Println(costsByPos)
	costByPosSorted := map[int]int{}
	//From(costsByPos).SortT(func(a KeyValue, b KeyValue) bool {
	//	fmt.Println(a.Value.(int), " < ", b.Value.(int))
	//	return a.Value.(int) < b.Value.(int) }).ToMap(&costByPosSorted)
	fmt.Println(From(costsByPos).OrderByT(func(a KeyValue) int { return a.Value.(int) }).First().(KeyValue).Value)

	//betterPosCost := From(costsByPos).Min()
	fmt.Println(costByPosSorted)
}
