package geometry

type Envelope struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

func NewEnvelopeF(top, bottom, left, right float64) Envelope {
	ne := Envelope{}
	ne.Top = top
	ne.Left = left
	ne.Bottom = bottom
	ne.Right = right
	return ne
}

func NewEnvelope(topLeft Point, bottomRight Point) Envelope {
	ne := Envelope{}
	ne.Top = topLeft.Y
	ne.Left = topLeft.X
	ne.Bottom = bottomRight.Y
	ne.Right = bottomRight.X
	return ne
}

func (e *Envelope) TopLeftPoint() Point {
	return NewPoint(e.Left, e.Top)
}

func (e *Envelope) BottomLeftPoint() Point {
	return NewPoint(e.Left, e.Bottom)
}

func (e *Envelope) TopRightPoint() Point {
	return NewPoint(e.Right, e.Top)
}

func (e *Envelope) BottomRightPoint() Point {
	return NewPoint(e.Right, e.Bottom)
}
