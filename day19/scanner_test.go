package day19

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Point_Rotate(t *testing.T) {

	p := &Point{
		x: 3.0,
		y: 4.0,
	}

	p2 := p.Rotate(90, "z")
	assert.Equal(t, -4.0, p2.x)
	assert.Equal(t, 3.0, p2.y)
}

func Test_Point_Rotate2(t *testing.T) {

	p := &Point{
		x: 0,
		y: 3.0,
		z: 4.0,
	}

	p2 := p.Rotate(90, "x")
	assert.Equal(t, -4.0, p2.y)
	assert.Equal(t, 3.0, p2.z)
}

func Test_All_Position(t *testing.T) {
	pOrigin := &Point{
		x: 42.0,
		y: 23.0,
		z: -5.0,
	}

	allPositions := pOrigin.AllRotations()

	assert.Equal(t, 24, len(allPositions))
}

func Test_Overlap_Scanners(t *testing.T) {
	s1Beacons := []*Point{
		newPoint(0, 2, 0),
		newPoint(4, 1, 0),
		newPoint(3, 3, 0),
	}
	s1 := newScanner("1")
	s1.beacons = s1Beacons

	s2Beacons := []*Point{
		newPoint(-1, -1, 0),
		newPoint(-5, 0, 0),
		newPoint(-2, 1, 0),
	}
	s2 := newScanner("2")
	s2.beacons = s2Beacons

	beacons, match, deltaPos := s1.Overlap(s2)

	fmt.Println(len(beacons), match, deltaPos)
	for _, b := range beacons {
		fmt.Println(b)
	}
}

func Test_SimilarDistances(t *testing.T) {
	s1Beacons := []*Point{
		newPoint(0, 2, 0),
		newPoint(4, 1, 0),
		newPoint(3, 3, 0),
	}
	s2Beacons := []*Point{
		newPoint(-1, -1, 0),
		newPoint(-5, 0, 0),
		newPoint(-2, 1, 0),
	}

	fmt.Println(SimilarDistance(s1Beacons, s2Beacons))
}

func Test_Orientations(t *testing.T) {
	scanners := ReadScannersFromFile("orientation_samples.txt")

	scanner_pos1 := scanners[0]

	for i_other_scanner, scanner_other_pos := range scanners[1:] {

		b1 := scanner_other_pos.beacons[0]
		all_b1_pos := b1.AllRotations()
		for i_rotation, rotated_b1 := range all_b1_pos {
			if scanner_pos1.beacons[0].Equals(rotated_b1) {
				fmt.Println("found!", b1, "=>", rotated_b1)

				fmt.Println("Scanner ", scanners[1:][i_other_scanner].name, i_other_scanner, "become :")
				for _, pos := range scanners[1:][i_other_scanner].beacons {
					rotations := pos.AllRotations()
					fmt.Println(pos, "=>", rotations[i_rotation])
				}

				break
			}
		}

	}
}

func Test_SimilarDistances_Sample(t *testing.T) {
	scanners := ReadScannersFromFile("sample.txt")

	similarDistance, delta, similarNuage, rotationIndex, rotationNuage := FindSimilaritiesByRotation(scanners[0].beacons, scanners[1].beacons)

	fmt.Println(similarDistance)
	fmt.Println("DELTA: ", delta)
	fmt.Println("SIMILAR NUAGE :")
	for _, d := range similarNuage {
		fmt.Println(d)
	}
	fmt.Println("ROTATED:")
	for _, d := range rotationNuage {
		fmt.Println(d)
	}
	fmt.Println(rotationIndex)
}
