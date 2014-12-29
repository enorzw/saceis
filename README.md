#saceis
## Geometry
* Point
* LineString
* Ring
* Polygon
* Envelope

```go
type IGeometry interface {
	String() string
	WktType() string
	Envelope() Envelope
}
```
