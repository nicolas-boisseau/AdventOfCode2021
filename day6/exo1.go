package day6

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Exo1() {
	lines := common.ReadLinesFromFile("day6/sample.txt")
	initFishes := strings.Split(lines[0], ",")

	fishes := make([]int, 0)
	for _, fish := range initFishes {
		n, err := strconv.Atoi(fish)
		if err != nil {
			log.Fatal(err)
		}
		fishes = append(fishes, n)
	}

	fmt.Println("Initial fishes:", fishes)

	for i := 1; i <= 80; i++ {
		newFishes := make([]int, 0)
		babies := 0
		for _, fish := range fishes {
			if fish == 0 {
				fish = 6
				babies++
			} else {
				fish--
			}
			newFishes = append(newFishes, fish)
		}
		for i := 0; i < babies; i++ {
			newFishes = append(newFishes, 8)
		}

		fishes = newFishes
		fmt.Println("After", i, "days (", len(fishes), ")")
		//fmt.Println("After", i, "days (", len(fishes), ") :", fishes)
	}
}
