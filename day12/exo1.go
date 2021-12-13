package day12

import (
	"AdventOfCode2021/common"
	"fmt"
	"strings"
)

func Exo1(fileName string) {
	lines := common.ReadLinesFromFile(fileName)

	g := newGraph()
	for _, line := range lines {
		values := strings.Split(line, "-")

		g.AddEdge(values[0], values[1])
	}

	//fmt.Println(g)

	g.printAllPaths("start", "end")
	fmt.Println(g.sum)
}
