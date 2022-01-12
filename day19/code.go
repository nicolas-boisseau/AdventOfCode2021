package day19

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strings"
)

func Process(fileName string) int {
	scanners := ReadScannersFromFile(fileName)

	for a, scannerA := range scanners {
		for b, scannerB := range scanners {
			if a == b {
				continue
			}

			allRotationNuageB := AllRotations(scannerB.beacons)
			for _, rotationNuageB := range allRotationNuageB {
				similarDistance, delta, similarNuage := SimilarDistance(scannerA.beacons, rotationNuageB)
				if similarDistance >= 12 {
					if _, isExist := scannerB.posRelativeTo[a]; !isExist {
						scannerB.posRelativeTo[a] = delta
						scannerB.beaconsRelativeTo[a] = rotationNuageB
					}
					//scannerA.posRelativeTo[b] = &Point{x: -delta.x, y: -delta.y, z: -delta.z}
					scannerA.realBeacons = append(scannerA.realBeacons, similarNuage...)
					break

					//fmt.Println(similarDistance)
					//fmt.Println(delta)
					//for _, d := range similarNuage {
					//	fmt.Println(d)
					//}
				}
			}

		}
	}

	//for

	fmt.Println("============ RESULTS ============")
	nb := 0
	for _, scanner := range scanners {
		fmt.Println(scanner.name)

		fmt.Println(scanner.posRelativeTo)
		//fmt.Println(scanner.beaconsRelativeTo)
		//fmt.Println(scanner.realBeacons)
		nb += len(scanner.realBeacons)
	}

	return nb
}

func ReadScannersFromFile(fileName string) []*Scanner {
	lines := common.ReadLinesFromFile(fileName)

	scanners := make([]*Scanner, 0)
	var currentScanner *Scanner
	for _, line := range lines {
		if strings.Contains(line, "scanner") {
			currentScanner = newScanner(line)
		} else if line != "" {
			reader := strings.NewReader(line)
			var x, y, z float64
			_, err := fmt.Fscanf(reader, "%f,%f,%f", &x, &y, &z)
			if err != nil {
				log.Fatal(err)
			}
			p := &Point{
				x: x,
				y: y,
				z: z,
			}
			currentScanner.beacons = append(currentScanner.beacons, p)
		} else {
			scanners = append(scanners, currentScanner)
		}
	}
	// Don't forget the last scanner !
	scanners = append(scanners, currentScanner)

	return scanners
}
