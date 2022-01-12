package day19

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample0(t *testing.T) {
	result := Process("sample0.txt")

	assert.Equal(t, 3, result)
}

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt")

	assert.Equal(t, 79, result)
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
