package day13

import (
	"bytes"
	"fmt"
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
	g := Grille{h: h, w: w}
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
	fmt.Fprintln(output, "GRILLE[W=", g.w, ",H=", g.h, "]")
	for i := range g.content {
		for j := range g.content[i] {
			if g.content[i][j] == 1 {
				fmt.Fprint(output, colorGreen, "#", colorReset)

			} else {
				fmt.Fprint(output, " ")
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
			if g.content[i][j] >= 1 {
				overlapSum++
			}
		}
	}
	return overlapSum
}

func (g *Grille) Fold(axe string, pos int) *Grille {
	var newG *Grille
	if axe == "y" {
		newG = newGrille(g.h/2, g.w)

		for i := 0; i < pos; i++ {
			for j := 0; j < g.w; j++ {
				a := g.content[i][j]
				b := g.content[g.h-i-1][j]
				if a == 1 || b == 1 {
					newG.content[i][j] = 1
				}
			}
		}
	} else {
		newG = newGrille(g.h, g.w/2)

		for i := 0; i < g.h; i++ {
			for j := 0; j < pos; j++ {
				a := g.content[i][j]
				b := g.content[i][g.w-j-1]
				if a == 1 || b == 1 {
					newG.content[i][j] = 1
				}
			}
		}
	}

	return newG
}
