package day9

import (
	"AdventOfCode2021/common"
	"fmt"
	"math"
	"strconv"
)

func (g *Grille) findLowPoints() []int {
	lowPoints := make([]int, 0)
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
				lowPoints = append(lowPoints, curr)
			}

		}
	}

	return lowPoints
}

func Exo1() {
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

	lowPoints := g.findLowPoints()

	fmt.Println(g)
	fmt.Println(lowPoints)

	sum := 0
	for _, n := range lowPoints {
		sum += 1 + n
	}
	fmt.Println(sum)
}
