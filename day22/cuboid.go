package day22

import (
	//"fmt"
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"math"
)

type Cuboid struct {
	//p1, p2 Point
	x1, x2, y1, y2, z1, z2 int64
	enabled                bool
	subCuboids             []*Cuboid
}

type Point struct {
	x, y, z int64
}

func (c *Cuboid) String() string {
	buff := bytes.NewBufferString("")
	if c.enabled {
		fmt.Fprintf(buff, "on (x%d)", c.OnCount())
	} else {
		fmt.Fprintf(buff, "off (x%d)", c.OnCount())
	}
	fmt.Fprintf(buff, " x=%d..%d, y=%d..%d, z=%d..%d", c.x1, c.x2, c.y1, c.y2, c.z1, c.z2)

	if c.subCuboids != nil {
		fmt.Fprint(buff, " subCuboids= [")
		for _, subC := range c.subCuboids {
			fmt.Fprint(buff, subC.String(), ",")
		}
		fmt.Fprint(buff, " ]")
	} else {
		fmt.Fprint(buff, " sub=[]")
	}

	return buff.String()
}

func (c *Cuboid) IsInScope() bool {
	inScope := func(i int64) bool {
		return i >= -50 && i <= 50
	}

	//return inScope(c.p1.x) && inScope(c.p2.x) && inScope(c.p1.y) && inScope(c.p2.y) && inScope(c.p1.z) && inScope(c.p2.z)
	return inScope(c.x1) && inScope(c.x2) && inScope(c.y1) && inScope(c.y2) && inScope(c.z1) && inScope(c.z2)
}

func (c *Cuboid) AllCorners() []Point {
	corners := make([]Point, 0)
	corners = append(corners, Point{x: c.x1, y: c.y1, z: c.z1})
	corners = append(corners, Point{x: c.x2, y: c.y1, z: c.z1})
	corners = append(corners, Point{x: c.x1, y: c.y2, z: c.z1})
	corners = append(corners, Point{x: c.x1, y: c.y1, z: c.z2})
	corners = append(corners, Point{x: c.x2, y: c.y2, z: c.z1})
	corners = append(corners, Point{x: c.x2, y: c.y1, z: c.z2})
	corners = append(corners, Point{x: c.x1, y: c.y2, z: c.z2})
	corners = append(corners, Point{x: c.x2, y: c.y2, z: c.z2})

	return corners
}

func (p Point) IsInside(c *Cuboid) bool {
	isInside := p.x >= c.x1 && p.x <= c.x2 && p.y >= c.y1 && p.y <= c.y2 && p.z >= c.z1 && p.z <= c.z2

	return isInside
}

func (c *Cuboid) Volume() int64 {
	return (int64(math.Abs(float64(c.x2-c.x1))) + 1) * (int64(math.Abs(float64(c.y2-c.y1))) + 1) * (int64(math.Abs(float64(c.z2-c.z1))) + 1)
}

func (c *Cuboid) OnCount() int64 {
	var subVolumeOffCount int64 = 0
	for _, subCuboid := range c.subCuboids {
		subVolumeOffCount += subCuboid.OffCount()
	}
	var subVolumeOnCount int64 = 0
	for _, subCuboid := range c.subCuboids {
		subVolumeOnCount += subCuboid.OnCount()
	}
	if c.enabled {
		volume := c.Volume()

		return volume - subVolumeOffCount + subVolumeOnCount
	} else {

		return subVolumeOnCount - subVolumeOffCount
	}
}

func (c *Cuboid) OffCount() int64 {
	var subVolumeOnCount int64 = 0
	for _, subCuboid := range c.subCuboids {
		subVolumeOnCount += subCuboid.OnCount()
	}
	var subVolumeOffCount int64 = 0
	for _, subCuboid := range c.subCuboids {
		subVolumeOffCount += subCuboid.OffCount()
	}
	if !c.enabled {
		volume := c.Volume()

		return volume - subVolumeOnCount + subVolumeOffCount
	} else {

		return subVolumeOnCount - subVolumeOffCount
	}
}

//func (p Point) IsAlreadyEnabledIn(c Cuboid) bool {
//	key := strconv.Itoa(int(p.x)) + "_" + strconv.Itoa(int(p.y)) + "_" + strconv.Itoa(int(p.z))
//	value, isPresent := c.specificSubPoints[key]
//	if isPresent {
//		return value
//	} else {
//		return c.enabled
//	}
//}

func (c *Cuboid) Intersect(c2 *Cuboid) (bool, map[string]bool) {

	fmt.Println("INTERSECT :")
	fmt.Println(c, "and", c2)

	commonPoints := make(map[string]bool)

	collision := false
	cCorners := c.AllCorners()
	c2Corners := c2.AllCorners()
	c2CornersInsideC := make([]Point, 0)
	cCornersInsideC := make([]Point, 0)

	for _, pc := range cCorners {
		if pc.IsInside(c2) {
			//fmt.Println("Collision ! c dans c2 !")
			//fmt.Println(pc)
			collision = collision || true
			cCornersInsideC = append(cCornersInsideC, pc)
		}
	}

	for _, pc2 := range c2Corners {
		if pc2.IsInside(c) {
			//fmt.Println("Collision ! c2 dans c !")
			//fmt.Println(pc2)
			collision = collision || true
			c2CornersInsideC = append(c2CornersInsideC, pc2)
		}
	}

	if collision {

		shouldEnableSubCuboid := !c.enabled && c2.enabled
		shouldEnableSubCuboid2 := !(!c.enabled && c2.enabled)
		//if pp2.IsAlreadyEnabledIn(c) && c2.enabled {
		//	commonPoints[key] = false
		//} else if !pp2.IsAlreadyEnabledIn(c) && c2.enabled {
		//	commonPoints[key] = false
		//} else {
		//	commonPoints[key] = c2.enabled
		//}

		corner := c2CornersInsideC[0]
		// try find a way through the opposite corner for x, y, z but stop when out of C
		var oppositeCorner Point
		oppositeCorners := linq.From(c2Corners).
			WhereT(func(p Point) bool { return p.x != corner.x && p.y != corner.y && p.z != corner.z })
		if oppositeCorners.Any() {
			oppositeCorner = oppositeCorners.First().(Point)
		} else {
			// Special case with a 1 pixel cuboid
			subCuboid := &Cuboid{
				x1:      corner.x,
				x2:      corner.x,
				y1:      corner.y,
				y2:      corner.y,
				z1:      corner.z,
				z2:      corner.z,
				enabled: shouldEnableSubCuboid,
			}
			subCuboid2 := &Cuboid{
				x1:      corner.x,
				x2:      corner.x,
				y1:      corner.y,
				y2:      corner.y,
				z1:      corner.z,
				z2:      corner.z,
				enabled: shouldEnableSubCuboid2,
			}

			c.subCuboids = append(c.subCuboids, subCuboid)
			c2.subCuboids = append(c.subCuboids, subCuboid2)
			return collision, commonPoints
		}

		var limitX, limitY, limitZ int64
		var incX, incY, incZ int64 = 1, 1, 1
		if corner.x > oppositeCorner.x {
			incX = -1
		}
		if corner.y > oppositeCorner.y {
			incY = -1
		}
		if corner.z > oppositeCorner.z {
			incZ = -1
		}
		for x := corner.x; (incX > 0 && x <= oppositeCorner.x) || (incX < 0 && x >= oppositeCorner.x); x += incX {
			if (Point{x: x, y: corner.y, z: corner.z}.IsInside(c)) {
				limitX = x
			} else {
				break
			}
		}

		for y := corner.y; (incY > 0 && y <= oppositeCorner.y) || (incY < 0 && y >= oppositeCorner.y); y += incY {
			if (Point{x: corner.x, y: y, z: corner.z}.IsInside(c)) {
				limitY = y
			} else {
				break
			}
		}
		for z := corner.z; (incZ > 0 && z <= oppositeCorner.z) || (incZ < 0 && z >= oppositeCorner.z); z += incZ {
			if (Point{x: corner.x, y: corner.y, z: z}.IsInside(c)) {
				limitZ = z
			} else {
				break
			}
		}

		if c.subCuboids == nil {
			c.subCuboids = make([]*Cuboid, 0)
		}

		subCuboid := &Cuboid{
			x1:      corner.x,
			x2:      limitX,
			y1:      corner.y,
			y2:      limitY,
			z1:      corner.z,
			z2:      limitZ,
			enabled: shouldEnableSubCuboid,
		}
		subCuboid2 := &Cuboid{
			x1:      corner.x,
			x2:      limitX,
			y1:      corner.y,
			y2:      limitY,
			z1:      corner.z,
			z2:      limitZ,
			enabled: shouldEnableSubCuboid2,
		}

		c.subCuboids = append(c.subCuboids, subCuboid)

		c2.subCuboids = append(c.subCuboids, subCuboid2)

		//for _, p2 := range c2CornersInsideC {
		//
		//	var limitX, limitY, limitZ int64 = c2.x2, c2.y2, c2.z2
		//	var incX, incY, incZ int64 = 1, 1, 1
		//	if p2.x == c2.x2 {
		//		limitX = c2.x1
		//		incX = -1
		//	}
		//	if p2.y == c2.y2 {
		//		limitY = c2.y1
		//		incY = -1
		//	}
		//	if p2.z == c2.z2 {
		//		limitZ = c2.z1
		//		incZ = -1
		//	}
		//	for x := p2.x; !((incX > 0 && x > limitX) || (incX < 0 && x < limitX)); x += incX {
		//		for y := p2.y; !((incY > 0 && y > limitY) || (incY < 0 && y < limitY)); y += incY {
		//			for z := p2.z; !((incZ > 0 && z > limitZ) || (incZ < 0 && z < limitZ)); z += incZ {
		//				//fmt.Println(x, y, z)
		//				pp2 := Point{
		//					x: x,
		//					y: y,
		//					z: z,
		//				}
		//
		//				if pp2.IsInside(c) {
		//					key := strconv.Itoa(int(x)) + "_" + strconv.Itoa(int(y)) + "_" + strconv.Itoa(int(z))
		//
		//					if pp2.IsAlreadyEnabledIn(c) && c2.enabled {
		//						commonPoints[key] = false
		//					} else if !pp2.IsAlreadyEnabledIn(c) && c2.enabled {
		//						commonPoints[key] = false
		//					} else {
		//						commonPoints[key] = c2.enabled
		//					}
		//				}
		//			}
		//		}
		//	}
		//}
	}

	//for k, v := range commonPoints {
	//	c.specificSubPoints[k] = v
	//}

	return collision, commonPoints
}
