package day15

import (
	"AdventOfCode2021/common"
	"fmt"
	//"github.com/jpierer/astar"
	//. "github.com/ahmetalpbalkan/go-linq"
	"strconv"
)

func Exo1(fileName string, size int, startNode, endNode string, complexMode bool) {
	lines := common.ReadLinesFromFile(fileName)

	g := newGrille(size, size)
	for i, line := range lines {
		//fmt.Println(line)

		for j, c := range line {
			n, _ := strconv.Atoi(string(c))
			g.content[i][j] = n
		}
	}

	if complexMode {
		g = g.Expand(5)
		fmt.Println("Expanded OK")
		//fmt.Println(g)
		//os.Exit(0)
	}

	isInRange := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < g.h && y < g.w
	}

	addEdgeIfInRange := func(fromX, fromY int, toX, toY int, g *Grille, gg *Graph) {
		if isInRange(toX, toY) {
			n1 := strconv.Itoa(fromX) + "_" + strconv.Itoa(fromY)
			n2 := strconv.Itoa(toX) + "_" + strconv.Itoa(toY)
			gg.AddEdge(n1, n2, g.content[toX][toY])
		}
	}

	gg := newGraph()
	for i, _ := range g.content {
		for j, _ := range g.content[i] {

			addEdgeIfInRange(i, j, i+1, j, g, gg)
			addEdgeIfInRange(i, j, i, j+1, g, gg)
			addEdgeIfInRange(i, j, i-1, j, g, gg)
			addEdgeIfInRange(i, j, i, j-1, g, gg)

		}
	}

	fmt.Println("Graph generated OK")

	//fmt.Println(gg)
	d := gg.dijkstra(startNode, endNode) //- g.content[0][0]
	fmt.Println(d)
	//gg.a_etoile("0_0", "9_9")
}
