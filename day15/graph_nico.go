package day15

import (
	"bytes"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/golang-collections/collections/set"
	"math"
)

type Graph struct {
	nodes      *set.Set
	edges      []*Edge
	edges2     map[string]map[string]int
	sum        int
	distByNode map[string]int
}

type Edge struct {
	start string
	end   string
	cost  int
}

func newEdge(start, end string, cost int) *Edge {
	e := Edge{
		start: start,
		end:   end,
		cost:  cost,
	}
	return &e
}

func newGraph() *Graph {
	g := Graph{
		nodes:      set.New(),
		edges:      make([]*Edge, 0),
		edges2:     make(map[string]map[string]int),
		distByNode: make(map[string]int),
	}

	return &g
}

func (g *Graph) AddEdge(n1, n2 string, cost int) {
	//fmt.Println("Adding edge...")
	if !g.nodes.Has(n1) {
		g.nodes.Insert(n1)
	}
	if !g.nodes.Has(n2) {
		g.nodes.Insert(n2)
	}

	if g.edges2[n1] == nil {
		g.edges2[n1] = make(map[string]int)
	}
	if g.edges2[n2] == nil {
		g.edges2[n2] = make(map[string]int)
	}
	if _, isPresent := g.edges2[n1][n2]; !isPresent {
		g.edges2[n1][n2] = cost
	}
	if _, isPresent := g.edges2[n2][n1]; !isPresent {
		g.edges2[n2][n1] = cost
	}

	//fmt.Println("Edge added.")
	//g.edges = append(g.edges, newEdge(n1, n2, cost))
	//g.edges = append(g.edges, newEdge(n2, n1, cost))
	//g.edges[n1] = n2
	//g.edges[n2] = n1
}

func (g *Graph) String() string {
	buffer := bytes.NewBufferString("")
	fmt.Fprintln(buffer, "NODES:", g.nodes)
	fmt.Fprintln(buffer, "EDGES:")
	//visited := make([]string, 0)
	for n1, sub := range g.edges2 {
		for n2, cost := range sub {
			fmt.Fprintln(buffer, n1, "->", n2, "[", cost, "]")
		}
	}
	return buffer.String()
}

func (g *Graph) dijkstra(startNode, endNode string) int {

	visited := set.New()

	g.nodes.Do(func(item interface{}) {
		g.distByNode[item.(string)] = math.MaxInt
	})

	previous := make(map[string]string)
	g.distByNode[startNode] = 0

	nodes := make([]string, 0)
	g.nodes.Do(func(item interface{}) {
		nodes = append(nodes, item.(string))
	})

	for visited.Len() < g.nodes.Len() {

		if visited.Len()%100 == 0 {
			fmt.Println(visited.Len(), "/", g.nodes.Len(), "(", float64(visited.Len())/float64(g.nodes.Len())*100.0, "%)")
		}

		a := From(nodes).OrderByT(func(s string) int { return g.distByNode[s] }).WhereT(func(n string) bool { return !visited.Has(n) }).First().(string)

		visited.Insert(a)

		edgesFromA := g.edges2[a]
		for e, cost := range edgesFromA {
			var ee interface{}
			ee = e
			if !visited.Has(ee) {
				if g.distByNode[e] > (g.distByNode[a] + cost) {
					g.distByNode[e] = g.distByNode[a] + cost
					previous[e] = a
				}
			}
		}
	}

	fmt.Println(visited.Len(), "/", g.nodes.Len(), "(", visited.Len()/g.nodes.Len()*100, "%)")

	fmt.Println(g.distByNode)

	return g.distByNode[endNode]
}

/*

for visited.Len() < len(g.nodes) {

		fmt.Println(visited.Len(), "/", len(g.nodes), "(", visited.Len() / len(g.nodes) * 100, "%)")

		a := From(g.nodes).WhereT(func (n string) bool { return !visited.Has(n) }).First().(string)
		visited.Insert(a)

		edges_from_a := g.edges2[a]
		for e, cost := range edges_from_a {
			if !visited.Has(e) {
				if g.distByNode[e] > (g.distByNode[a] + cost) {
					g.distByNode[e] = g.distByNode[a] + cost
					previous[e] = a
				}
			}
		}
	}

*/
