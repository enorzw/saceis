package saceis

import (
	_ "fmt"
	"strings"
)

type PointsCollection []Points

type Polygon struct {
	PointsCollection
}

func NewPolygon(pc PointsCollection) Polygon {
	gon := Polygon{}
	count := len(pc)
	gon.PointsCollection = make([]Points, count, 10)
	for i := 0; i < count; i++ {
		pcount := len(pc[i])
		gon.PointsCollection[i] = make([]Point, pcount, 10)
		for j := 0; j < pcount; j++ {
			gon.PointsCollection[i][j] = New2DPoint(pc[i][j].X, pc[i][j].Y)
			// gon.PointsCollection[i][j].X = pc[i][j].X
			// gon.PointsCollection[i][j].Y = pc[i][j].Y
			// gon.PointsCollection[i][j].Z = pc[i][j].Z
			// gon.PointsCollection[i][j].M = pc[i][j].M
			// gon.PointsCollection[i][j].Dimension = pc[i][j].Dimension
		}
	}
	return gon
}

func (p *Polygon) ToString() string {
	buf := "Polygon ("
	count := len(p.PointsCollection)
	for i := 0; i < count; i++ {
		pcount := len(p.PointsCollection[i])
		if pcount > 0 {
			buf += "("
			for j := 0; j < pcount; j++ {
				buf += p.PointsCollection[i][j].Coordinate.ToString()
				buf += ","
			}
		}
		buf = strings.TrimRight(buf, ",")
		buf += "),"
	}
	buf = strings.TrimRight(buf, ",")
	buf += ")"
	return buf
}
