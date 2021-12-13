package day5

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strings"
)

func (g *Grille) drawSegment2(s Segment) {
	inc_x := 0
	if s.start.x > s.end.x {
		inc_x = -1
	} else if s.start.x < s.end.x {
		inc_x = 1
	}
	inc_y := 0
	if s.start.y > s.end.y {
		inc_y = -1
	} else if s.start.y < s.end.y {
		inc_y = 1
	}

	for x, y := s.start.x, s.start.y; x != s.end.x+inc_x || y != s.end.y+inc_y; x, y = x+inc_x, y+inc_y {
		g.content[y][x] += 1
	}
}

func Exo2() {
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
		g.drawSegment2(s)
	}
	fmt.Println(g)

	fmt.Println("Overlap = ", g.Overlaps())
}
