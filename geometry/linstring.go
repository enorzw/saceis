package geometry

import (
	"strings"
)

type LineString Points

func NewLineStringByArray(points []Point) LineString {
	length := len(points)
	line := make([]Point, length, length)
	copy(line, points)
	return line
}

func NewLineString(points ...Point) LineString {
	return NewLineStringByArray(points)
}

func (p LineString) Count() int {
	return len(p)
}

func (l LineString) Length() float64 {
	var count int = l.Count()
	var length float64 = 0
	for i := 0; i < count-1; i++ {
		length += l[i].Distance(l[i+1])
	}
	return length
}

func (p LineString) String() string {
	buf := "LINESTRING ("
	count := p.Count()
	for i := 0; i < count; i++ {
		buf += p[i].Coordinate.String()
		buf += ","
	}
	buf = strings.TrimRight(buf, ",")
	buf += ")"
	return buf
}

func (l LineString) Clone() LineString {
	length := len(l)
	nl := make([]Point, length, length)
	copy(nl, l)
	return nl
}

func (l LineString) WktType() string {
	return "LINESTRING"
}

func (l LineString) Envelope() Envelope {
	var min_x, min_y, max_x, max_y float64
	for i := 0; i < len(l); i++ {
		if i == 0 {
			min_x = l[i].X
			min_y = l[i].Y
			max_x = l[i].X
			max_y = l[i].Y
		} else {
			if l[i].X < min_x {
				min_x = l[i].X
			} else if l[i].X > max_x {
				max_x = l[i].X
			}
			if l[i].Y < min_y {
				min_y = l[i].Y
			} else if l[i].Y > max_y {
				max_y = l[i].Y
			}
		}
	}
	return NewEnvelopeF(max_y, min_y, min_x, max_x)
}
