package geometry

import (
	"strings"
)

type Ring Points

func NewRingByArray(points []Point) Ring {
	length := len(points)
	ring := make([]Point, length, length)
	copy(ring, points)
	return ring
}

func NewRing(points ...Point) Ring {
	length := len(points)
	ring := make([]Point, length, length)
	copy(ring, points)
	return ring
}

func (r Ring) Count() int {
	return len(r)
}

func (r Ring) IsClose() bool {
	count := len(r)
	if count <= 3 {
		return false
	} else {
		if r[0].Distance(r[count-1]) < 0.000001 {
			return true
		}
	}
	return false
}

func (r Ring) Close() Ring {
	if !r.IsClose() && len(r) >= 3 {
		end := r[0].Clone()
		return append(r, end)
	}
	return nil
}

func (r Ring) Clone() Ring {
	return r.Clone()
}

func (r Ring) String() string {
	buf := "Ring ("
	count := r.Count()
	for i := 0; i < count; i++ {
		buf += r[i].Coordinate.String()
		buf += ","
	}
	buf = strings.TrimRight(buf, ",")
	buf += ")"
	return buf
}
