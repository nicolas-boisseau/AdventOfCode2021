package day5

import (
	"AdventOfCode2021/common"
	"bytes"
	"fmt"
	"log"
	"strings"
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

func (g *Grille) String() string {
	output := bytes.NewBufferString("")
	for i := range g.content {
		fmt.Fprintln(output, g.content[i])
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

func Exo1() {
	lines := common.ReadLinesFromFile("day5/input.txt")

	segments := make([]Segment, 0)
	for _, line := range lines {
		var x1, y1, x2, y2 int
		reader := strings.NewReader(line)
		_, err := fmt.Fscanf(reader, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatal(err)
		}
		p1 := Point{x: x1, y: y1}
		p2 := Point{x: x2, y: y2}
		segment := Segment{start: p1, end: p2}
		fmt.Println(segment)
		segments = append(segments, segment)
	}

	g := newGrille(1000, 1000)
	for _, s := range segments {
		g.drawSegment(s)
	}
	fmt.Println(g)

	fmt.Println("Overlap = ", g.Overlaps())
}
