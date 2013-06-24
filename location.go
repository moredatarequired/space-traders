package starmap

import (
	"math"
)

type Positioner interface {
	Position() (X, Y, Z float64)
}

func SquaredDistance(p1, p2 Positioner) float64 {
	x1, y1, z1 := p1.Position()
	x2, y2, z2 := p2.Position()
	xd, yd, zd := x2 - x1, y2 - y1, z2 - z1
	return xd*xd + yd*yd + zd*zd
}

func Distance(p1, p2 Positioner) float64 {
	return math.Pow(SquaredDistance(p1, p2), 0.5)
}
