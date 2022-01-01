package day15

import (
	"bytes"
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}
type Segment struct {
	start Point
	end   Point
}

type Grille struct {
	h       int
	w       int
	content [][]int
	visited [][]bool
	sumPath int
}

func newGrille(h int, w int) *Grille {
	g := Grille{h: h, w: w}
	g.content = make([][]int, h)
	for i := range g.content {
		g.content[i] = make([]int, w)
	}
	g.visited = make([][]bool, h)
	for i := range g.visited {
		g.visited[i] = make([]bool, w)
	}
	return &g
}

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

func (g *Grille) String() string {
	output := bytes.NewBufferString("")
	for i := range g.content {
		for j := range g.content[i] {
			if g.visited[i][j] {
				fmt.Fprint(output, colorGreen, g.content[i][j], colorReset)

			} else {
				fmt.Fprint(output, g.content[i][j])
			}
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func (g *Grille) Overlaps() int {
	overlapSum := 0
	for i, _ := range g.content {
		for j, _ := range g.content[i] {
			if g.content[i][j] > 1 {
				overlapSum++
			}
		}
	}
	return overlapSum
}

type PointWithValue struct {
	value int
	pos   Point
}

func (g *Grille) findLowPoints2() []PointWithValue {
	lowPoints := make([]PointWithValue, 0)
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
				l := PointWithValue{value: curr, pos: Point{x: i, y: j}}
				lowPoints = append(lowPoints, l)
			}

		}
	}

	return lowPoints
}

func (g *Grille) Expand(scale int) *Grille {
	scaledG := newGrille(g.h*scale, g.w*scale)

	for i := 0; i < scale; i++ {
		for j := 0; j < scale; j++ {

			for sub_i := 0; sub_i < g.h; sub_i++ {
				for sub_j := 0; sub_j < g.w; sub_j++ {

					scaledValue := 0
					if i == 0 && j == 0 {
						scaledValue = g.content[sub_i][sub_j]
					} else if i > 0 && j == 0 {
						scaledValue = scaledG.content[sub_i+(i-1)*g.h][sub_j+j*g.w] + 1
					} else if j > 0 && i == 0 {
						scaledValue = scaledG.content[sub_i+i*g.h][sub_j+(j-1)*g.w] + 1
					} else {
						scaledValue = scaledG.content[sub_i+i*g.h][sub_j+(j-1)*g.w] + 1
					}

					if scaledValue > 9 {
						scaledValue = 1
					}

					scaledG.content[sub_i+i*g.h][sub_j+j*g.w] = scaledValue
				}
			}

			//fmt.Println(scaledG)
		}
	}

	return scaledG
}
