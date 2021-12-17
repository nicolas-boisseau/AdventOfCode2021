package day17

import (
	"AdventOfCode2021/common"
	"fmt"
	"strings"
)

func Process(fileName string) (int, int) {
	lines := common.ReadLinesFromFile(fileName)

	var x1, x2, y1, y2 int
	reader := strings.NewReader(lines[0])
	fmt.Fscanf(reader, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	fmt.Println(x1, x2, y1, y2)

	maxY := 0
	reachedShots := 0
	for initialVx := -1000; initialVx <= 1000; initialVx++ {
		for initialVy := -1000; initialVy <= 1000; initialVy++ {

			// Launch
			maxY, reachedShots = TryLaunch(initialVx, initialVy, x1, x2, y1, y2, maxY, reachedShots)
		}
	}

	//maxY, reachedShots = TryLaunch(6, 0, x1, x2, y1, y2, maxY, reachedShots)

	fmt.Println(maxY)
	fmt.Println(reachedShots)

	return maxY, reachedShots
}

func TryLaunch(initialVx int, initialVy int, x1 int, x2 int, y1 int, y2 int, maxY int, reachedShots int) (int, int) {
	//fmt.Println("Try launch with v=", initialVx, initialVy)
	vX := initialVx
	vY := initialVy
	missed := false
	reached := false
	x := 0
	y := 0
	maxYForLaunch := 0
	for !missed && !reached {
		x += vX
		y += vY

		if y > maxYForLaunch {
			maxYForLaunch = y
		}

		//fmt.Println("pos=", x, y)
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			//fmt.Println("Target reached. v was", initialVx, ",", initialVy)

			//fmt.Printf("%d,%d", initialVx, initialVy)
			//fmt.Println()
			if maxYForLaunch > maxY {
				maxY = maxYForLaunch
			}

			reachedShots++
			reached = true
		}

		if x > x2 || y < y1 {
			missed = true
		}

		if vX > 0 {
			vX--
		} else if vX < 0 {
			vX++
		}
		vY--
	}
	return maxY, reachedShots
}
