package day19

import (
	"bytes"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/golang-collections/collections/set"
	"math"
)

type Point struct {
	x, y, z float64
}

type Scanner struct {
	name              string
	posRelativeTo     map[int]*Point
	beaconsRelativeTo map[int][]*Point
	beacons           []*Point
	realBeacons       []*Point
}

func newPoint(x, y, z float64) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
	}
}

func (p *Point) Equals(p2 *Point) bool {
	return p.x == p2.x && p.y == p2.y && p.z == p2.z
}

func (p *Point) Distance(p2 *Point) (float64, *Point) {
	distance := math.Sqrt(
		math.Pow(p2.x-p.x, 2) +
			math.Pow(p2.y-p.y, 2) +
			math.Pow(p2.z-p.z, 2))
	delta := &Point{
		x: p2.x - p.x,
		y: p2.y - p.y,
		z: p2.z - p.z,
	}

	return distance, delta
}

func FindSimilaritiesByRotation(nuage1, nuage2 []*Point) (int, *Point, []*Point, int, []*Point) {
	allRotationNuage := AllRotations(nuage1)
	for i, rotationNuage := range allRotationNuage {
		similarDistance, delta, similarNuage := SimilarDistance(nuage2, rotationNuage)
		if similarDistance >= 12 {
			return similarDistance, delta, similarNuage, i, rotationNuage
		}
	}

	return -1, nil, nil, -1, nil
}

func SimilarDistance(nuage1 []*Point, nuage2 []*Point) (int, *Point, []*Point) {

	distances := make(map[float64]int)
	deltas := make(map[float64]*Point)
	pMax := make(map[float64][]*Point)
	for _, p1 := range nuage1 {
		for _, p2 := range nuage2 {
			distance, delta := p2.Distance(p1)
			distances[distance]++
			deltas[distance] = delta

			_, isExist := pMax[distance]
			if !isExist {
				pMax[distance] = make([]*Point, 0)
			}
			pMax[distance] = append(pMax[distance], p1)
		}
	}

	max := From(distances).OrderByDescendingT(func(kv KeyValue) int { return kv.Value.(int) }).First().(KeyValue)

	return max.Value.(int), deltas[max.Key.(float64)], pMax[max.Key.(float64)]
}

func newScanner(name string) *Scanner {
	return &Scanner{
		name:              name,
		posRelativeTo:     make(map[int]*Point), // position unknown
		beaconsRelativeTo: make(map[int][]*Point),
		beacons:           make([]*Point, 0),
		realBeacons:       make([]*Point, 0),
	}
}

func (p *Point) String() string {
	buff := bytes.NewBufferString("")
	fmt.Fprintf(buff, "x:%0.0f, y:%0.0f, z:%0.0f", p.x, p.y, p.z)
	return buff.String()
}

// https://www.mathworks.com/matlabcentral/answers/123763-how-to-rotate-entire-3d-data-with-x-y-z-values-along-a-particular-axis-say-x-axis
// https://keisan.casio.com/exec/system/15362817755710
func (p *Point) Rotate(degre float64, axe string) *Point {

	angle := degre * (math.Pi / 180.0)

	if axe == "z" {
		return &Point{
			x: math.Round(p.x*math.Cos(angle) - p.y*math.Sin(angle)),
			y: math.Round(p.x*math.Sin(angle) + p.y*math.Cos(angle)),
			z: p.z,
		}
	} else if axe == "x" {
		return &Point{
			x: p.x,
			y: math.Round(p.y*math.Cos(angle) - p.z*math.Sin(angle)),
			z: math.Round(p.y*math.Sin(angle) + p.z*math.Cos(angle)),
		}
	} else {
		return &Point{
			x: math.Round(p.z*math.Sin(angle) + p.x*math.Cos(angle)),
			y: p.y,
			z: math.Round(p.z*math.Cos(angle) - p.x*math.Sin(angle)),
		}
	}
}

func (p *Point) ReverseDirection(axe string) *Point {
	if axe == "z" {
		return &Point{
			x: p.x,
			y: p.y,
			z: -p.z,
		}
	} else if axe == "x" {
		return &Point{
			x: -p.x,
			y: p.y,
			z: p.z,
		}
	} else {
		return &Point{
			x: p.x,
			y: -p.y,
			z: p.z,
		}
	}
}

func AllRotations(nuage []*Point) [][]*Point {
	allRotationsNuages := make([][]*Point, 24)
	for i := 0; i < 24; i++ {
		allRotationsNuages[i] = make([]*Point, 0)
	}
	for _, p := range nuage {
		allRotationsP := p.AllRotations()
		for i, _ := range allRotationsNuages {
			allRotationsNuages[i] = append(allRotationsNuages[i], allRotationsP[i])
		}
	}

	return allRotationsNuages
}

func (pOrigin *Point) AllRotations() []*Point {

	allPositions := make([]*Point, 0)

	hashSet := set.New()
	record := func(p *Point) {
		str := p.String()
		//fmt.Println(str)
		if !hashSet.Has(str) {
			hashSet.Insert(str)
			allPositions = append(allPositions, p)
		}
	}

	record(pOrigin)
	pX := pOrigin.Rotate(90, "x")
	record(pX)
	p := pX.Rotate(90, "y")
	record(p)
	p = pX.Rotate(180, "y")
	record(p)
	p = pX.Rotate(270, "y")
	record(p)
	p = pX.Rotate(90, "z")
	record(p)
	p = pX.Rotate(180, "z")
	record(p)
	p = pX.Rotate(270, "z")
	record(p)

	pX = pOrigin.Rotate(180, "x")
	record(pX)
	p = pX.Rotate(90, "y")
	record(p)
	p = pX.Rotate(180, "y")
	record(p)
	p = pX.Rotate(270, "y")
	record(p)
	p = pX.Rotate(90, "z")
	record(p)
	p = pX.Rotate(180, "z")
	record(p)
	p = pX.Rotate(270, "z")
	record(p)

	pX = pOrigin.Rotate(270, "x")
	record(pX)
	p = pX.Rotate(90, "y")
	record(p)
	p = pX.Rotate(180, "y")
	record(p)
	p = pX.Rotate(270, "y")
	record(p)
	p = pX.Rotate(90, "z")
	record(p)
	p = pX.Rotate(180, "z")
	record(p)
	p = pX.Rotate(270, "z")
	record(p)

	pY := pOrigin.Rotate(90, "y")
	record(pY)
	p = pY.Rotate(90, "x")
	record(p)
	p = pY.Rotate(180, "x")
	record(p)
	p = pY.Rotate(270, "x")
	record(p)
	p = pY.Rotate(90, "z")
	record(p)
	p = pY.Rotate(180, "z")
	record(p)
	p = pY.Rotate(270, "z")
	record(p)

	pY = pOrigin.Rotate(180, "y")
	record(pY)
	p = pY.Rotate(90, "x")
	record(p)
	p = pY.Rotate(180, "x")
	record(p)
	p = pY.Rotate(270, "x")
	record(p)
	p = pY.Rotate(90, "z")
	record(p)
	p = pY.Rotate(180, "z")
	record(p)
	p = pY.Rotate(270, "z")
	record(p)

	pY = pOrigin.Rotate(270, "y")
	record(pY)
	p = pY.Rotate(90, "x")
	record(p)
	p = pY.Rotate(180, "x")
	record(p)
	p = pY.Rotate(270, "x")
	record(p)
	p = pY.Rotate(90, "z")
	record(p)
	p = pY.Rotate(180, "z")
	record(p)
	p = pY.Rotate(270, "z")
	record(p)

	pZ := pOrigin.Rotate(90, "z")
	record(pZ)
	p = pZ.Rotate(90, "x")
	record(p)
	p = pZ.Rotate(180, "x")
	record(p)
	p = pZ.Rotate(270, "x")
	record(p)
	p = pZ.Rotate(90, "y")
	record(p)
	p = pZ.Rotate(180, "y")
	record(p)
	p = pZ.Rotate(270, "y")
	record(p)

	pZ = pOrigin.Rotate(180, "z")
	record(pZ)
	p = pZ.Rotate(90, "x")
	record(p)
	p = pZ.Rotate(180, "x")
	record(p)
	p = pZ.Rotate(270, "x")
	record(p)
	p = pZ.Rotate(90, "y")
	record(p)
	p = pZ.Rotate(180, "y")
	record(p)
	p = pZ.Rotate(270, "y")
	record(p)

	pZ = pOrigin.Rotate(270, "z")
	record(pZ)
	p = pZ.Rotate(90, "x")
	record(p)
	p = pZ.Rotate(180, "x")
	record(p)
	p = pZ.Rotate(270, "x")
	record(p)
	p = pZ.Rotate(90, "y")
	record(p)
	p = pZ.Rotate(180, "y")
	record(p)
	p = pZ.Rotate(270, "y")
	record(p)

	return allPositions
}

func (s *Scanner) Overlap(s2 *Scanner) ([]*Point, bool, *Point) {
	maxMatches := 0
	var maxMatchBeacons []*Point
	var deltaPos *Point
	for x := -500.0; x < 500; x++ {
		for y := -500.0; y < 500; y++ {
			for z := -500.0; z < 500; z++ {
				matches := 0
				matchBeacons := make([]*Point, 0)
				for _, p2 := range s2.beacons {
					for _, p1 := range s.beacons {
						if p1.x == p2.x+x && p1.y == p2.y+y && p1.z == p2.z+z {
							matches++
							matchBeacons = append(matchBeacons, p2)
						}
					}
				}

				if matches > maxMatches {
					maxMatches = matches
					maxMatchBeacons = matchBeacons
					deltaPos = &Point{
						x: x,
						y: y,
						z: z,
					}
				}
			}
		}
	}

	fmt.Println("Max matches:", maxMatches, deltaPos)

	return maxMatchBeacons, maxMatches > 0, deltaPos
}
