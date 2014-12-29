package test

import (
	"fmt"
	"saceis/geometry"
	"testing"
)

func TestCoordinate(t *testing.T) {
	fmt.Println("-------------Coordinate Test------------")

	a := geometry.NewCoordinate(geometry.Dimension_XY, 119.232143, 49.123124)
	fmt.Println(a)

	aa := geometry.NewCoordinate(geometry.Dimension_XY, 116.232143, 45.123124)
	fmt.Println(a.Distance(aa))

	b := geometry.NewCoordinate(geometry.Dimension_XYZ, 119.232143, 49.123124, 20.2323425)
	fmt.Println(b)

	bb := geometry.NewCoordinate(geometry.Dimension_XYZ, 116.232143, 45.123124, 29.2323425)
	fmt.Println(b.Distance(bb))

	c := geometry.NewCoordinate(geometry.Dimension_XYM, 119.232143, 49.123124, 2.23)
	fmt.Println(c)

	d := geometry.NewCoordinate(geometry.Dimension_XYZM, 119.232143, 49.123124, 20.2323425, 2.23)
	fmt.Println(d)

	fmt.Println("-------------Coordinate Test------------")
}

func TestGeometry(t *testing.T) {
	p := geometry.NewPoint(10.123, 20.321)
	fmt.Println(p)
}
