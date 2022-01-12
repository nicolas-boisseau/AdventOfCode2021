package day22

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Cuboid struct {
	x1, x2, y1, y2, z1, z2 int
	enabled                bool
}

func (c Cuboid) IsInScope() bool {
	inScope := func(i int) bool {
		return i >= -50 && i <= 50
	}

	return inScope(c.x1) && inScope(c.x2) && inScope(c.y1) && inScope(c.y2) && inScope(c.z1) && inScope(c.z2)
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	cuboids := make([]Cuboid, 0)
	for _, line := range lines {

		// on x=10..12,y=10..12,z=10..12
		reader := strings.NewReader(line)
		on := line[0:2] == "on"
		prefix := "on"
		if !on {
			prefix = "off"
		}
		c := Cuboid{}
		var x1, x2, y1, y2, z1, z2 int
		_, err := fmt.Fscanf(reader, prefix+" x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
		if err != nil {
			log.Fatal(err)
		}

		c.x1 = x1
		c.x2 = x2
		c.y1 = y1
		c.y2 = y2
		c.z1 = z1
		c.z2 = z2
		c.enabled = on

		if c.IsInScope() || complex {
			cuboids = append(cuboids, c)
		}
	}

	fmt.Println(cuboids)

	return Compute(cuboids)
}

func Compute(cuboids []Cuboid) int {

	cubes := make(map[string]bool)

	for _, cuboid := range cuboids {
		for x := cuboid.x1; x <= cuboid.x2; x++ {
			for y := cuboid.y1; y <= cuboid.y2; y++ {
				for z := cuboid.z1; z <= cuboid.z2; z++ {
					cubes[strconv.Itoa(x)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z)] = cuboid.enabled
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
