package day6

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"log"
	"strconv"
	"strings"
)

//type fishSet struct {
//	counts map[int]int
//}

func Exo2() {
	lines := common.ReadLinesFromFile("day6/input.txt")
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
	fishesMap := make(map[int]int)
	nombersOfFishOfN := func(fishesArray []int, n int) int {
		return From(fishesArray).WhereT(func(a int) bool { return a == n }).Count()
	}
	for i := 0; i <= 8; i++ {
		fishesMap[i] = nombersOfFishOfN(fishes, i)
	}

	fmt.Println("init map :", fishesMap)

	for i := 1; i <= 256; i++ {
		newFishesMap := make(map[int]int)

		zeros := 0
		for k, v := range fishesMap {
			if k == 0 {
				newFishesMap[8] = v
				zeros = v
			} else {
				newFishesMap[k-1] = v
			}
		}
		if zeros > 0 {
			newFishesMap[6] += zeros
		}

		fishesMap = newFishesMap
		//fmt.Println("After", i, "days (", len(fishes), ")")
		totalFishes := 0
		for _, v := range fishesMap {
			totalFishes += v
		}
		//fmt.Println("After", i, "days (", totalFishes, ") :", fishesMap)
		fmt.Println("After", i, "days (", totalFishes, ")")
	}
}
