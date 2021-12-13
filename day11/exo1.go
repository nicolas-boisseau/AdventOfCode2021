package day11

import (
	"AdventOfCode2021/common"
	"fmt"
	"strconv"
)

func (g *Grille) allZero() bool {
	for i, _ := range g.content {
		for j, _ := range g.content[i] {
			if g.content[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
func (g *Grille) containsSup(n int) bool {
	for i, _ := range g.content {
		for j, _ := range g.content[i] {
			if g.content[i][j] > n {
				return true
			}
		}
	}
	return false
}

func (g *Grille) pass1() {
	for i, _ := range g.content {
		for j, _ := range g.content[i] {
			g.content[i][j]++
		}
	}
}

func inc(n int) int {
	return n + 1
}

func (g *Grille) pass2() []Point {
	isInRange := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < g.h && y < g.w
	}
	doIfInRange := func(x, y int, f func(int) int) {
		if isInRange(x, y) {
			g.content[x][y] = f(g.content[x][y])
		}
	}

	exploded := make([]Point, 0)
	for i, _ := range g.content {
		for j, _ := range g.content[i] {
			if g.content[i][j] > 9 {
				doIfInRange(i-1, j-1, inc)
				doIfInRange(i-1, j, inc)
				doIfInRange(i-1, j+1, inc)
				doIfInRange(i, j-1, inc)
				doIfInRange(i, j+1, inc)
				doIfInRange(i+1, j-1, inc)
				doIfInRange(i+1, j, inc)
				doIfInRange(i+1, j+1, inc)

				g.content[i][j] = -99999
				exploded = append(exploded, Point{x: i, y: j})
			}
		}
	}

	return exploded
}

//
//func (g *Grille) pass2() []PointWithValue {
//	lowPoints := make([]PointWithValue, 0)
//	for i, _ := range g.content {
//		for j, curr := range g.content[i] {
//
//			up := math.MaxInt
//			down := math.MaxInt
//			left := math.MaxInt
//			right := math.MaxInt
//			if j > 0 {
//				up = g.content[i][j-1]
//			}
//			if j < len(g.content[i])-1 {
//				down = g.content[i][j+1]
//			}
//			if i > 0 {
//				left = g.content[i-1][j]
//			}
//			if i < len(g.content)-1 {
//				right = g.content[i+1][j]
//			}
//
//			if curr < up && curr < down && curr < left && curr < right {
//				l := PointWithValue{value: curr, pos: Point{x: i, y: j}}
//				lowPoints = append(lowPoints, l)
//			}
//
//		}
//	}
//
//	return lowPoints
//}

const MaxSteps = 1000

func Exo1(fileName string) {
	lines := common.ReadLinesFromFile(fileName)

	g := newGrille(10, 10)
	for i, line := range lines {
		//fmt.Println(line)

		for j, c := range line {
			n, _ := strconv.Atoi(string(c))
			g.content[i][j] = n
		}
	}

	// start
	sumFlash := 0
	for step := 0; step < MaxSteps; step++ {

		g.pass1()
		//fmt.Println("After pass 1:")
		//fmt.Println(g)

		exploded := make([]Point, 0)
		for g.containsSup(9) {
			newExploded := g.pass2()
			for _, e := range newExploded {
				exploded = append(exploded, e)
			}
			//fmt.Println("PASS2")
			//fmt.Println(g)
		}
		for _, p := range exploded {
			g.content[p.x][p.y] = 0
			sumFlash++
		}

		if g.allZero() {
			fmt.Println("All flash !!!!!")
			fmt.Println(g)
			fmt.Println("Step:", step+1)
		}

		//fmt.Println("After pass 2:")
		//fmt.Println(g)

	}

	fmt.Println(g)
	fmt.Println(sumFlash)
}
