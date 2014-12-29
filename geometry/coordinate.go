package geometry

import (
	"fmt"
	"math"
)

//Dimension
const (
	Dimension_None = 0
	Dimension_X    = 1
	Dimension_Y    = 2
	Dimension_XY   = 3
	Dimension_Z    = 4
	Dimension_XYZ  = 7
	Dimension_M    = 8
	Dimension_XYM  = 11
	Dimension_XYZM = 15
)

//坐标：支持二维、三维和M
type Coordinate struct {
	Dimension int
	X         float64
	Y         float64
	Z         float64
	M         float64
}

//创建坐标：参数dimension：Dimesion枚举，coors：坐标序列顺序为x、y、z、m，长度不能超过4
func NewCoordinate(dimension int, coors ...float64) Coordinate {
	coord := Coordinate{}
	length := len(coors)
	switch dimension {
	case Dimension_XY:
		coord.Dimension = Dimension_XY
		if length == 2 {
			coord.X = coors[0]
			coord.Y = coors[1]
		} else {
			errText := fmt.Sprintf("Demension[%v] is not equal to the count[%v] of coors.", "XY", length)
			panic(errText)
		}
	case Dimension_XYZ:
		coord.Dimension = Dimension_XYZ
		if length == 3 {
			coord.X = coors[0]
			coord.Y = coors[1]
			coord.Z = coors[2]
		} else {
			errText := fmt.Sprintf("Demension[%v] is not equal to the count[%v] of coors.", "XYZ", length)
			panic(errText)
		}

	case Dimension_XYM:
		coord.Dimension = Dimension_XYM
		if length == 3 {
			coord.X = coors[0]
			coord.Y = coors[1]
			coord.M = coors[2]
		} else {
			errText := fmt.Sprintf("Demension[%v] is not equal to the count[%v] of coors.", "XYM", length)
			panic(errText)
		}
	case Dimension_XYZM:
		coord.Dimension = Dimension_XYZM
		if length == 4 {
			coord.X = coors[0]
			coord.Y = coors[1]
			coord.Z = coors[2]
			coord.M = coors[3]
		} else {
			errText := fmt.Sprintf("Demension[%v] is not equal to the count[%v] of coors.", "XYZM", length)
			panic(errText)
		}
	default:
		errText := "dimension must be one of Dimension_XY,Dimension_XYZ,Dimension_XYM,Dimension_XYZM"
		panic(errText)
	}
	return coord
}

//将坐标转化为wkt格式的坐标字符串，在各个geometry的wkt转化中都有使用
func (c Coordinate) String() string {
	switch c.Dimension {
	case Dimension_XY:
		return fmt.Sprintf("%v %v", c.X, c.Y)
	case Dimension_XYZ:
		return fmt.Sprintf("%v %v %v", c.X, c.Y, c.Z)
	case Dimension_XYM:
		return fmt.Sprintf("%v %v %v", c.X, c.Y, c.M)
	case Dimension_XYZM:
		return fmt.Sprintf("%v %v %v %v", c.X, c.Y, c.Z, c.M)
	default:
		panic("coordinate demension error!")
	}
}

func (fromC Coordinate) Distance(toC Coordinate) float64 {
	if fromC.Dimension == toC.Dimension {
		switch fromC.Dimension {
		case Dimension_XY:
			fallthrough
		case Dimension_XYM:
			disX := toC.X - fromC.X
			disY := toC.Y - fromC.Y
			return math.Sqrt((disX*disX + disY*disY))
		case Dimension_XYZ:
			fallthrough
		case Dimension_XYZM:
			disX := toC.X - fromC.X
			disY := toC.Y - fromC.Y
			disZ := toC.Z - fromC.Z
			return math.Sqrt((disX*disX + disY*disY + disZ*disZ))
		default:
			panic("coordinate dimesion is wrong!")
		}
	} else {
		panic("Dimension is not Equal!")
	}
}
