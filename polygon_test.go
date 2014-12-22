package saceis

import (
	"fmt"
	"testing"
)

func TestNewPolygon(t *testing.T) {

	var psc PointsCollection
	psc = make([]Points, 2)
	psc[0] = []Point{New2DPoint(10.1, 20.1), New2DPoint(10.2, 20.2), New2DPoint(10.2, 20.2)}
	psc[1] = []Point{New2DPoint(10.1, 20.1), New2DPoint(10.2, 20.2), New2DPoint(10.2, 20.2), New2DPoint(10.2, 20.2)}
	fmt.Println(psc)
	gon := NewPolygon(psc)
	fmt.Println(gon.ToString())
}
