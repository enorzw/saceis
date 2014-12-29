//几何图形<br>
//点、线、面的几何图形定义<br>
//作者:于中玮<br>
package geometry

type IGeometry interface {
	String() string
	WktType() string
	Envelope() Envelope
}
