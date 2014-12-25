package geometry

type Points []Point

func NewPointsByArray(points []Point) Points {
	length := len(points)
	ps := make([]Point, length, length)
	copy(ps, points)
	return ps
}

func NewPoints(points ...Point) Points {
	return NewPointsByArray(points)
}

func (p Points) Append(points ...Point) Points {
	length := len(points)
	lengthself := len(p)
	ps := make([]Point, length+lengthself, length+lengthself)
	at := copy(ps, p)
	copy(ps[at:], points)
	return ps
}
