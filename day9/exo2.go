package day9

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"math"
	"strconv"
)

type lowPoint struct {
	value int
	pos   Point
}

func (g *Grille) findLowPoints2() []lowPoint {
	lowPoints := make([]lowPoint, 0)
	for i, _ := range g.content {
		for j, curr := range g.content[i] {

			up := math.MaxInt
			down := math.MaxInt
			left := math.MaxInt
			right := math.MaxInt
			if j > 0 {
				up = g.content[i][j-1]
			}
			if j < len(g.content[i])-1 {
				down = g.content[i][j+1]
			}
			if i > 0 {
				left = g.content[i-1][j]
			}
			if i < len(g.content)-1 {
				right = g.content[i+1][j]
			}

			if curr < up && curr < down && curr < left && curr < right {
				l := lowPoint{value: curr, pos: Point{x: i, y: j}}
				lowPoints = append(lowPoints, l)
			}

		}
	}

	return lowPoints
}

func (g *Grille) findBasins(p lowPoint) []lowPoint {
	output := make([]lowPoint, 0)
	g.content[p.pos.x][p.pos.y] = 9
	if p.value >= 9 {
		return output
	}
	output = append(output, p)

	i := p.pos.x
	j := p.pos.y

	addNeighbor := func(p lowPoint, g *Grille, i int, j int, output []lowPoint) []lowPoint {

		neighbor := g.content[i][j]
		if math.Abs(float64(p.value-neighbor)) >= 1 {
			toAddNeighbors := g.findBasins(lowPoint{neighbor, Point{i, j}})
			for _, toAdd := range toAddNeighbors {
				output = append(output, toAdd)
			}
		}
		return output
	}

	if j > 0 {
		output = addNeighbor(p, g, i, j-1, output)
	}
	if j < len(g.content[i])-1 {
		output = addNeighbor(p, g, i, j+1, output)
	}
	if i > 0 {
		output = addNeighbor(p, g, i-1, j, output)
	}
	if i < len(g.content)-1 {
		output = addNeighbor(p, g, i+1, j, output)
	}

	return output
}

func Exo2() {
	lines := common.ReadLinesFromFile("day9/input.txt")

	h := len(lines)
	w := len(lines[0])

	g := newGrille(h, w)

	for i, line := range lines {
		for j, char := range line {
			n, _ := strconv.Atoi(string(char))
			g.content[i][j] = n
		}
	}

	lowPoints := g.findLowPoints2()

	//basins := make([][]int, 0)
	sizes := make([]int, 0)
	for _, lowP := range lowPoints {
		basin := g.findBasins(lowP)
		fmt.Println("len  =", len(basin))

		//basinDistinct := make([]lowPoint, 0)
		//From(basin).DistinctByT(func(ll lowPoint) string { return string(ll.pos.x) + "," + string(ll.pos.y) }).ToSlice(&basinDistinct)
		//fmt.Println("len distinct =", len(basinDistinct))
		sizes = append(sizes, len(basin))
		//for _, l := range basinDistinct {
		//	fmt.Print(l.value)
		//	fmt.Print(",")
		//}
		//fmt.Println()
	}

	sortedSizes := make([]int, 0)
	From(sizes).OrderByDescendingT(func(a int) int { return a }).ToSlice(&sortedSizes)
	fmt.Println(sortedSizes)
	fmt.Println(sortedSizes[0] * sortedSizes[1] * sortedSizes[2])

	//fmt.Println(g)
	//fmt.Println(lowPoints)

	//sum := 0
	//for _, n := range lowPoints {
	//	sum += 1 + n
	//}
	//fmt.Println(sum)
}
