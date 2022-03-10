package day19

//
//import (
//	"fmt"
//)
//
//func ProcessOld(fileName string) int {
//	scanners := ReadScannersFromFile(fileName)
//
//	for a, scannerA := range scanners {
//		for b, scannerB := range scanners {
//			if a == b {
//				continue
//			}
//
//			allRotationNuageB := AllRotations(scannerB.beacons)
//			for _, rotationNuageB := range allRotationNuageB {
//				similarDistance, delta, _ := SimilarDistance(scannerA.beacons, rotationNuageB)
//				if similarDistance >= 12 {
//
//					nuageBrelativeToA := MoveNuage(rotationNuageB, delta)
//					sumEqualsToA := 0
//					for _, pB := range nuageBrelativeToA {
//						for _, pA := range scannerA.beacons {
//							if pB.Equals(pA) {
//								sumEqualsToA++
//								pB.isBeacon = true
//								pA.isBeacon = true
//							}
//						}
//					}
//
//					if sumEqualsToA >= 12 {
//
//						if _, isExist := scannerB.posRelativeTo[a]; !isExist {
//							scannerB.posRelativeTo[a] = delta
//							scannerB.beaconsRelativeTo[a] = nuageBrelativeToA
//						}
//						//scannerA.posRelativeTo[b] = &Point{x: -delta.x, y: -delta.y, z: -delta.z}
//						//scannerA.realBeacons = append(scannerA.realBeacons, similarNuage...)
//						break
//					}
//
//					//fmt.Println(similarDistance)
//					//fmt.Println(delta)
//					//for _, d := range similarNuage {
//					//	fmt.Println(d)
//					//}
//				}
//			}
//
//		}
//	}
//
//	// Now we have all data needed, try move all scanners relative to 0,0,0 (scanner 0)
//	scannersFromZero := make([]*Scanner, 0)
//	scannersFromZero = append(scannersFromZero, scanners[0])
//	for _, scanner := range scanners[1:] {
//		deltaToZero := FindDeltaToZero(scanner, scanners, &Point{0, 0, 0, false})
//		scanner.posRelativeTo[0] = deltaToZero
//		scanner.beacons = MoveNuage(scanner.beacons, deltaToZero)
//		scannersFromZero = append(scannersFromZero, scanner)
//	}
//
//	fmt.Println("============ RESULTS ============")
//	nb := 0
//	for _, scanner := range scanners {
//		fmt.Println(scanner.name)
//
//		fmt.Println(scanner.posRelativeTo)
//		//fmt.Println(scanner.beaconsRelativeTo)
//		//fmt.Println(scanner.realBeacons)
//		nb += len(scanner.realBeacons)
//	}
//
//	fmt.Println("============ RESULTS FROM ZERO ============")
//	for _, scanner := range scanners {
//		fmt.Println(scanner.name)
//
//		fmt.Println(scanner.posRelativeTo)
//		//fmt.Println(scanner.beaconsRelativeTo)
//		//fmt.Println(scanner.realBeacons)
//	}
//
//	return nb
//}
//
//
