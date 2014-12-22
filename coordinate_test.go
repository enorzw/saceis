package saceis

import (
	"fmt"
	"testing"
)

func TestNewCoordinate(t *testing.T) {
	a := NewCoordinate(Dimension_XY, 116.23423423, 23.239872)
	b := NewCoordinate(Dimension_XYZ, 116.23423423, 23.239872, 20.2323)
	c := NewCoordinate(Dimension_XYM, 116.23423423, 23.239872, 20.2323)
	d := NewCoordinate(Dimension_XYZM, 116.23423423, 23.239872, 20.2323, 10.23238)

	if a.ToString() != "116.23423423 23.239872" || a.Dimension != Dimension_XY {
		t.Errorf("%s,%v", a.ToString(), a.Dimension)
	}

	if b.ToString() != "116.23423423 23.239872 20.2323" || b.Dimension != Dimension_XYZ {
		t.Errorf("%s,%v", b.ToString(), b.Dimension)
	}

	if c.ToString() != "116.23423423 23.239872 20.2323" || c.Dimension != Dimension_XYM {
		t.Errorf("%s,%v", c.ToString(), c.Dimension)
	}

	if d.ToString() != "116.23423423 23.239872 20.2323 10.23238" || d.Dimension != Dimension_XYZM {
		t.Errorf("%s,%v", d.ToString(), d.Dimension)
	}
}

func TestNewCoordinatePanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if err != "Demension[XYZ] is not equal to the count[2] of coors." {
				t.Error("panic info error!")
			}
		}
	}()

	a := NewCoordinate(Dimension_XYZ, 116.23423423, 23.239872)
	fmt.Println(a.ToString())
}
