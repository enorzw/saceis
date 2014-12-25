package geometry

import (
	"strings"
)

type Polygon struct {
	OuterRing  Ring
	InnerRings []Ring
}

func NewPolygonByArray(points []Point) Polygon {
	polygon := Polygon{}
	ring := NewRingByArray(points)
	polygon.OuterRing = ring.Close()
	return polygon
}

func NewPolygon(points ...Point) Polygon {
	return NewPolygonByArray(points)
}

func NewPolygonByRing(outer Ring) Polygon {
	polygon := Polygon{}
	polygon.OuterRing = outer.Close()
	return polygon
}

func NewPolygonWithHole(outer Ring, inner []Ring) Polygon {
	polygon := Polygon{}
	polygon.OuterRing = outer.Close()
	polygon.InnerRings = inner
	return polygon
}

func (p Polygon) String() string {
	buf := "POLYGON ("
	if p.OuterRing != nil {
		buf += "("
		for _, point := range p.OuterRing {
			buf += point.Coordinate.String()
			buf += ","
		}
		buf = strings.TrimRight(buf, ",")
		buf += ")"
	}
	buf += ")"
	return buf
}

func (p Polygon) Envelope() Envelope {
	l := p.OuterRing
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
