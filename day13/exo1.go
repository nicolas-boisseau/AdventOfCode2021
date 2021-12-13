package day13

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"strconv"
	"strings"
)

type Fold struct {
	axe   string
	value int
}

func Exo1(fileName string) {
	lines := common.ReadLinesFromFile(fileName)

	points := make([]Point, 0)
	folds := make([]Fold, 0)
	mode := "points"

	for _, line := range lines {
		//fmt.Println(line)
		if line == "" {
			mode = "fold"
		}

		if mode == "points" {
			pts := strings.Split(line, ",")
			x, _ := strconv.Atoi(pts[0])
			y, _ := strconv.Atoi(pts[1])
			p := Point{x: x, y: y}
			points = append(points, p)
		} else if line != "" && mode == "fold" {
			reader := strings.NewReader(line)
			var s string
			fmt.Fscanf(reader, "fold along %s", &s)
			ss := strings.Split(s, "=")
			axe := ss[0]
			value, _ := strconv.Atoi(ss[1])
			f := Fold{axe: axe, value: value}
			folds = append(folds, f)
		}
	}

	fmt.Println(points)
	fmt.Println(folds)

	max_x := From(points).SelectT(func(pp Point) int { return pp.x }).Max().(int)
	max_y := From(points).SelectT(func(pp Point) int { return pp.y }).Max().(int)

	fmt.Println(max_x, "x", max_y)

	g := newGrille(max_y+1, max_x+1)
	for _, p := range points {
		g.content[p.y][p.x] = 1
	}

	//fmt.Println(g)

	foldedG := g
	for i, fold := range folds {
		foldedG = foldedG.Fold(fold.axe, fold.value)
		fmt.Println("Fold", i, ":")
		//fmt.Println(foldedG)
		fmt.Println(foldedG.Overlaps())
	}
	fmt.Println(foldedG)
}
