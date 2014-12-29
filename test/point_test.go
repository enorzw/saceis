package test

import (
	_ "fmt"
	"saceis/geometry"
	"testing"
)

func TestNewPoint(t *testing.T) {
	point := geometry.NewPoint(116.234, 39.321)
	if point.String() != "POINT (116.234 39.321)" {
		t.Errorf("Point String Wrong:%s", point.String())
	}
}

func TestNewPointZ(t *testing.T) {
	point := geometry.NewPointZ(116.234, 39.321, 20.235)
	if point.String() != "POINT Z(116.234 39.321 20.235)" {
		t.Errorf("PointZ String Wrong:%s", point.String())
	}
}

func TestNewPointM(t *testing.T) {
	point := geometry.NewPointM(116.234, 39.321, 20.235)
	if point.String() != "POINT M(116.234 39.321 20.235)" {
		t.Errorf("PointM String Wrong:%s", point.String())
	}
}

func TestNewPointZM(t *testing.T) {
	point := geometry.NewPointZM(116.234, 39.321, 10.875, 20.235)
	if point.String() != "POINT ZM(116.234 39.321 10.875 20.235)" {
		t.Errorf("PointZM String Wrong:%s", point.String())
	}
}
