package saceis

import (
	"strings"
)

type Points []Point

type Polyline struct {
	Points
}

func NewPolyline(points ...Point) Polyline {
	line := Polyline{}
	line.Points = make([]Point, len(points))
	copy(line.Points, points)
	return line
}

func (p Polyline) Count() int {
	return len(p.Points)
}

func (p Polyline) ToString() string {
	buf := "Polyline ("
	count := p.Count()
	for i := 0; i < count; i++ {
		buf += p.Points[i].Coordinate.ToString()
		buf += ","
	}
	buf = strings.TrimRight(buf, ",")
	buf += ")"
	return buf
}
