package geometry

type IGeometry interface {
	String() string
	WktType() string
	Envelope() Envelope
}
