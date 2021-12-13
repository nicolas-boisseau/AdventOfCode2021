package day11

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
}

func newGrille(h int, w int) *Grille {
	g := Grille{h: 10, w: 10}
	g.content = make([][]int, h)
	for i := range g.content {
		g.content[i] = make([]int, w)
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
			if g.content[i][j] == 0 {
				fmt.Fprint(output, colorGreen, g.content[i][j], colorReset)

			} else {
				fmt.Fprint(output, g.content[i][j])
			}
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func (g *Grille) drawSegment(s Segment) {
	if s.start.x == s.end.x {
		inc := 1
		if s.start.y > s.end.y {
			inc = -1
		}
		for y := s.start.y; y != s.end.y+inc; y += inc {
			g.content[y][s.start.x] += 1
		}
	} else if s.start.y == s.end.y {
		inc := 1
		if s.start.x > s.end.x {
			inc = -1
		}
		for x := s.start.x; x != s.end.x+inc; x += inc {
			g.content[s.start.y][x] += 1
		}
	} else {
	}
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
