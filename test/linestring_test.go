package test

import (
	"saceis/geometry"
	"testing"
)

func TestNewLineString(t *testing.T) {
	point1 := geometry.NewPoint(116.234, 39.321)
	point2 := geometry.NewPoint(116.234, 40.321)
	point3 := geometry.NewPoint(116.234, 41.321)
	point4 := geometry.NewPoint(116.234, 42.321)
	line := geometry.NewLineString(point1, point2, point3, point4)

	if line.String() != "LINESTRING (116.234 39.321,116.234 40.321,116.234 41.321,116.234 42.321)" {
		t.Errorf("LINESTRING STRING Wrong:%s", line.String())
	}

	if line.Count() != 4 {
		t.Errorf("LINESTRING Count Wrong:%d", line.Count())
	}

	if line.Length() != 3 {
		t.Errorf("LINESTRING Length Wrong:%d", line.Length())
	}
}
