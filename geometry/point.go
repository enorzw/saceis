package geometry

type Point struct {
	Coordinate
}

func NewPoint(x, y float64) Point {
	var p = Point{}
	p.Coordinate = NewCoordinate(Dimension_XY, x, y)
	return p
}

func NewPointZ(x, y, z float64) Point {
	var p = Point{}
	p.Coordinate = NewCoordinate(Dimension_XYZ, x, y, z)
	return p
}

func NewPointM(x, y, m float64) Point {
	var p = Point{}
	p.Coordinate = NewCoordinate(Dimension_XYM, x, y, m)
	return p
}

func NewPointZM(x, y, z, m float64) Point {
	var p = Point{}
	p.Coordinate = NewCoordinate(Dimension_XYZM, x, y, z, m)
	return p
}

func (p Point) Distance(q Point) float64 {
	return p.Coordinate.Distance(q.Coordinate)
}

func (p Point) String() string {
	buf := "POINT (" + p.Coordinate.String() + ")"
	return buf
}

func (p Point) Clone() Point {
	var np = Point{}
	np.Coordinate = Coordinate{}
	np.Dimension = p.Dimension
	np.X = p.X
	np.Y = p.Y
	np.Z = p.Z
	np.M = p.M
	return np
}

func (p Point) WktType() string {
	return "POINT"
}

func (p Point) Envelope() Envelope {
	return NewEnvelope(p, p)
}
