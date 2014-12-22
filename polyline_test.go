package saceis

import (
	"fmt"
	"testing"
)

func TestNewPolyline(t *testing.T) {
	line := NewPolyline(New2DPoint(10.1, 20.1), New2DPoint(10.2, 20.2))
	fmt.Println("Polyline:")
	fmt.Println(line.ToString())
}
