package day22

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCuboid_IsInScope(t *testing.T) {
	c1 := Cuboid{
		x1:      48,
		x2:      50,
		y1:      34,
		y2:      12,
		z1:      7,
		z2:      10,
		enabled: false,
	}

	c2 := Cuboid{
		x1:      48,
		x2:      50,
		y1:      34,
		y2:      12,
		z1:      134,
		z2:      655,
		enabled: false,
	}

	assert.Equal(t, true, c1.IsInScope())
	assert.Equal(t, false, c2.IsInScope())
}

func TestCuboid_Intersect(t *testing.T) {
	c1 := &Cuboid{
		x1:      0,
		x2:      5,
		y1:      4,
		y2:      10,
		z1:      3,
		z2:      6,
		enabled: true,
	}
	c2 := &Cuboid{
		x1:      -4,
		x2:      2,
		y1:      4,
		y2:      6,
		z1:      3,
		z2:      5,
		enabled: false,
	}

	isCollision, commonPoints := c1.Intersect(c2)

	fmt.Println(isCollision)
	fmt.Println(commonPoints)
}

func TestCuboid_Intersect_OnePoint(t *testing.T) {
	c1 := &Cuboid{
		x1:      0,
		x2:      2,
		y1:      4,
		y2:      6,
		z1:      6,
		z2:      8,
		enabled: true,
	}
	c2 := &Cuboid{
		x1:      -2,
		x2:      0,
		y1:      0,
		y2:      4,
		z1:      2,
		z2:      6,
		enabled: false,
	}

	fmt.Println(c1.Intersect(c2))
}

func TestCuboid_Volume(t *testing.T) {
	c1 := Cuboid{
		x1: 10,
		x2: 12,
		y1: 10,
		y2: 12,
		z1: 10,
		z2: 12,
	}

	assert.Equal(t, 27, c1.Volume())
}
