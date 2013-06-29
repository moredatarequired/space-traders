package lib

import (
	"math"
)

type Ship struct {
	Position [3]float64  // meters
	Velocity [3]float64  // meters per second
	Acceleration [3]float64  // meters per second^2
}

// Move the ship over t seconds.
func (s *Ship) Move(t float64) {
	for i, a := range s.Acceleration {
		s.Position[i] += s.Velocity[i]*t + a*t*t/2
		s.Velocity[i] += a*t
	}
}

func SquaredDistance(p1, p2 []float64) float64 {
	total := 0.0
	for i, p := range p1 {
		d := p - p2[i]
		total += d*d
	}
	return total
}

func (s1 *Ship) SquaredDistance(s2 *Ship) float64 {
	return SquaredDistance(s1.Position[:], s2.Position[:])
}

func Distance(p1, p2 []float64) float64 {
	return math.Pow(SquaredDistance(p1, p2), 0.5)
}

func (s1 *Ship) Distance(s2 *Ship) float64 {
	return math.Pow(Distance(s1.Position[:], s2.Position[:]), 0.5)
}

func Norm1(xs []float64) float64 {
	total := 0.0
	for _, x := range xs { total += math.Abs(x) }
	return total
}

// Scale a vector by a constant.
func ScaleVector(v []float64, c float64) []float64 {
	out := make([]float64, len(v))
	for i, val := range v { out[i] = val * c }
	return out
}

func VectorDiff(p1, p2 []float64) []float64 {
	v := make([]float64, len(p1))
	for i := range v { v[i] = p2[i] - p1[i] }
	return v
}

func UnitVector(v []float64) []float64 {
	return ScaleVector(v, 1 / Norm1(v))
}

// Adopt the maximum acceleration away from the given point.
func (s *Ship) Flee(p []float64, a float64) {
	dir := VectorDiff(p, s.Position[:])
	for i, v := range ScaleVector(UnitVector(dir), a) {
		s.Acceleration[i] = v
	}
}

func CrossProduct(u, v []float64) []float64 {
	return []float64{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0]}
}

// Return the perpendicular to u that lies nearest v.
func PerpendicularNearest(u, v []float64) []float64 {
	return CrossProduct(u, CrossProduct(v, u))
}

// Set the acceleration of s in order to circle t.
func (s *Ship) CircleTarget(t *Ship, a float64) {
	v := VectorDiff(s.Velocity[:], t.Velocity[:])
	p := VectorDiff(s.Position[:], t.Position[:])
	acc := ScaleVector(UnitVector(PerpendicularNearest(p, v)), a)
	for i, a := range acc {
		s.Acceleration[i] = a
	}
}
