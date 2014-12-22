package saceis

type Point struct {
	*Coordinate
}

type IPoint interface {
}

func New2DPoint(x, y float64) *Point {
	var p = new(Point)
	p.Coordinate = NewCoordinate(Dimension_XY, x, y)
	return p
}

func New3DPoint(x, y, z float64) *Point {
	var p = new(Point)
	p.Coordinate = NewCoordinate(Dimension_XYZ, x, y, z)
	return p
}

func (p *Point) Distance(q *Point) float64 {
	return p.Coordinate.Distance(q.Coordinate)
}

func (p *Point) ToString() string {
	buf := "POINT " + p.Coordinate.ToString()
	return buf
}
