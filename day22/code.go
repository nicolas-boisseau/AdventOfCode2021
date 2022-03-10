package day22

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Process(fileName string, complex bool) int64 {
	lines := common.ReadLinesFromFile(fileName)

	cuboids := make([]*Cuboid, 0)
	for _, line := range lines {

		// on x=10..12,y=10..12,z=10..12
		reader := strings.NewReader(line)
		on := line[0:2] == "on"
		prefix := "on"
		if !on {
			prefix = "off"
		}
		c := &Cuboid{
			//specificSubPoints: make(map[string]bool),
		}
		var x1, x2, y1, y2, z1, z2 int64
		_, err := fmt.Fscanf(reader, prefix+" x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
		if err != nil {
			log.Fatal(err)
		}

		// ensure that x1 is always < x2
		if x1 < x2 {
			c.x1 = x1
			c.x2 = x2
		} else {
			c.x1 = x2
			c.x2 = x1
		}
		if y1 < y2 {
			c.y1 = y1
			c.y2 = y2
		} else {
			c.y1 = y2
			c.y2 = y1
		}
		if z1 < z2 {
			c.z1 = z1
			c.z2 = z2
		} else {
			c.z1 = z2
			c.z2 = z1
		}
		c.enabled = on

		if c.IsInScope() || complex {
			cuboids = append(cuboids, c)
		}
	}

	//fmt.Println(cuboids)

	if complex {
		return OptimizedCompute(cuboids)
	} else {
		return int64(Compute(cuboids))
	}
}

func OptimizedCompute(cuboids []*Cuboid) int64 {
	processedCuboids := make([]*Cuboid, 0)
	processedCuboids = append(processedCuboids, cuboids[0])

	for _, cuboid := range cuboids[1:] {
		fmt.Println("Processing ", cuboid)
		for _, processedCuboid := range processedCuboids {
			processedCuboid.Intersect(cuboid)
		}
		processedCuboids = append(processedCuboids, cuboid)
	}

	fmt.Println("PROCESSED:")
	var sum int64 = 0
	for _, processedCuboid := range processedCuboids {
		fmt.Println(processedCuboid)
		sum += processedCuboid.OnCount()
	}

	return sum
}

func Compute(cuboids []*Cuboid) int {

	cubes := make(map[string]bool)

	for _, cuboid := range cuboids {
		for x := cuboid.x1; x <= cuboid.x2; x++ {
			for y := cuboid.y1; y <= cuboid.y2; y++ {
				for z := cuboid.z1; z <= cuboid.z2; z++ {
					cubes[strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+"_"+strconv.Itoa(int(z))] = cuboid.enabled
				}
			}
		}
	}

	sum := 0
	for _, v := range cubes {
		if v {
			sum++
		}
	}

	return sum
}
