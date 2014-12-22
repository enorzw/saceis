package saceis

import (
	"fmt"
	"testing"
)

func TestNew2DPoint(t *testing.T) {
	p := New2DPoint(123, 23)
	q := New2DPoint(122, 89)

	fmt.Println(p.Distance(q))
	fmt.Println(p.Coordinate.ToString())
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.Z)
	fmt.Println(p.Dimension)
	fmt.Println(p.ToString())
}
