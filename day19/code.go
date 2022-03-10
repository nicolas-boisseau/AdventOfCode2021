package day19

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"strconv"
)

func Process(fileName string) int {
	scanners := ReadScannersFromFile(fileName)

	for a, scannerA := range scanners {
		for b, scannerB := range scanners {
			if a == b {
				continue
			}

			scannerB.AlignWith(scannerA, a)
		}
	}

	for k, scanner := range scanners {
		if k == 0 {
			continue
		}

		count := linq.From(scanner.beacons).
			WhereT(func(p *Point) bool {
				return p.isBeacon
			}).
			Count()
		fmt.Println(count, "in scanner", k)

		if _, isAlignedWith0 := scanner.deltaRelativeTo[0]; !isAlignedWith0 {
			path, pathFound := FindPathToZero(scanner, scanners, []int{})
			if pathFound {
				// first step from current scanner
				newScannerBeacons := scanner.beaconsRelativeTo[path[0]]
				newScannerBeacons = MoveNuage(newScannerBeacons, scanner.deltaRelativeTo[path[0]])
				delta := scanner.deltaRelativeTo[path[0]]
				// Next, move / turn between scanners until zero is reached
				currentStep := path[0]
				for _, nextStep := range path[1:] {

					newScannerBeacons = MoveNuage(newScannerBeacons, scanners[currentStep].deltaRelativeTo[nextStep])
					newScannerBeacons = AllRotations(newScannerBeacons)[scanners[currentStep].rotationIndexRelativeTo[nextStep]]
					//if stepBeacons, isExist := scanner.beaconsRelativeTo[nextStep]; isExist {
					//	newScannerBeacons = stepBeacons
					delta = delta.Add(scanners[currentStep].deltaRelativeTo[nextStep])
					//}

					currentStep = nextStep
				}

				//for i, _ := range newScannerBeacons {
				//	newScannerBeacons[i] = newScannerBeacons[i].Add(delta)
				//}

				count := linq.From(newScannerBeacons).
					WhereT(func(p *Point) bool {
						return p.isBeacon
					}).
					Count()
				fmt.Println("After rotation & move :", count, "in scanner", k)

				scanner.beaconsRelativeTo[0] = newScannerBeacons
				scanner.deltaRelativeTo[0] = delta
			}
		}
	}

	// Print all
	fmt.Println("============ RESULTS ============")
	realBeacons := make(map[string]bool)
	for k, scanner := range scanners {
		fmt.Println("***", scanner.name, "***")
		//fmt.Println(scanner.beacons)
		//fmt.Println(scanner.deltaRelativeTo)
		//if k > 0 {
		//	path, pathFound := FindPathToZero(scanner, scanners, []int{})
		//	fmt.Println("Path to ZERO found=", pathFound, "path=", path)
		//}
		//fmt.Println(scanner.beaconsRelativeTo)
		if k == 0 {
			fmt.Println("Real beacons from 0")
			for _, b := range scanner.beacons {
				if b.isBeacon {
					fmt.Println(b.x, b.y, b.z)
					key := strconv.Itoa(int(b.x)) + "_" + strconv.Itoa(int(b.y)) + "_" + strconv.Itoa(int(b.z))
					realBeacons[key] = true
				}
			}
		} else if _, isZeroNeighbor := scanner.beaconsRelativeTo[0]; isZeroNeighbor {
			fmt.Println("Real beacons from ", k)
			for _, b := range scanner.beaconsRelativeTo[0] {
				if b.isBeacon {
					fmt.Println(b.x, b.y, b.z)
					//p0 := b.Add(scanner.deltaRelativeTo[0])
					key := strconv.Itoa(int(b.x)) + "_" + strconv.Itoa(int(b.y)) + "_" + strconv.Itoa(int(b.z))
					realBeacons[key] = true
				}
			}
		}
	}

	fmt.Println("======== REAL BEACONS ========")
	for k, _ := range realBeacons {
		fmt.Println(k)
	}

	return len(realBeacons)
}

func FindPathToZero(scanner *Scanner, otherScanners []*Scanner, previous []int) ([]int, bool) {
	path, zeroReached := previous, false
	for k, _ := range scanner.deltaRelativeTo {
		if k == 0 {
			path = append(path, 0)
			zeroReached = true
			break
		} else {
			path = append(path, k)
			subKPath, subZeroFound := FindPathToZero(otherScanners[k], otherScanners, path)

			if subZeroFound {
				path = subKPath
				zeroReached = subZeroFound
				break
			}
		}
	}

	return path, zeroReached
}
