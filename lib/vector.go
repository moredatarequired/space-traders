// Implements operations on three dimensional vectors consisting of float64.

package lib

import (
	"math"
)

// Equality approximation for float64.
func fequal(x, y float64) bool { return math.Abs(x - y) < 1e-15 }

type Vector struct {
	X, Y, Z float64
}

// Approximation for vector equality.
func (v *Vector) Equal(u *Vector) bool {
	return fequal(v.X, u.X) && fequal(v.Y, u.Y) && fequal(v.Z, u.Z)
}

func (v *Vector) Dot(u *Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v *Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// The 2-norm, or Euclidean length.
func (v *Vector) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v *Vector) SquaredDistance(u *Vector) float64 {
	x, y, z := u.X - v.X, u.Y - v.Y, u.Z - v.Z
	return x*x + y*y + z*z
}

func (v *Vector) Distance(u *Vector) float64 {
	return math.Sqrt(v.SquaredDistance(u))
}

// Produce the vector c*v, where c is a scalar.
func (v *Vector) Times(c float64) *Vector {
	return &Vector{v.X * c, v.Y * c, v.Z * c}
}

// Multiply this vector by a scalar.
func (v *Vector) TimesInPlace(c float64) {
	v.X *= c
	v.Y *= c
	v.Z *= c
}

// Produce a vector parallel to v with length l.
func (v *Vector) ScaleTo(l float64) *Vector {
	if *v == (Vector{}) { return &Vector{} }
	s := l / v.Length()
	return v.Times(s)
}

// Resize v to length l.
func (v *Vector) ScaleToInPlace(l float64) {
	if *v != (Vector{}) {
		v.TimesInPlace(l / v.Length())
	}
}

// Produce the unit vector along v.
func (v *Vector) Unit() *Vector {
	return v.ScaleTo(1)
}

// Turn v into the parallel unit vector.
func (v *Vector) UnitInPlace() {
	v.ScaleTo(1)
}

// Produce the vector (v + u).
func (v *Vector) Plus(u *Vector) *Vector {
	return &Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// Set v to the vector sum v + u.
func (v *Vector) PlusInPlace(u *Vector) {
	v.X += u.X
	v.Y += u.Y
	v.Z += u.Z
}

// Produce the vector difference (v - u).
func (v *Vector) Minus(u *Vector) *Vector {
	return &Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// Set v to the vector sum v - u.
func (v *Vector) MinusInPlace(u *Vector) {
	v.X -= u.X
	v.Y -= u.Y
	v.Z -= u.Z
}

// Produce the cross product (v x u).
func (v *Vector) Cross(u *Vector) *Vector {
	return &Vector{v.Y*u.Z - v.Z*u.Y, v.Z*u.X - v.X*u.Z, v.X*u.Y - v.Y*u.X}
}

// Set v to the cross product (v x u).
func (v *Vector) CrossInPlace(u *Vector) {
	v.X, v.Y, v.Z = v.Y*u.Z - v.Z*u.Y, v.Z*u.X - v.X*u.Z, v.X*u.Y - v.Y*u.X
}

// Set v to the vector v + s*u.
func (v *Vector) AddWithScaleInPlace(u *Vector, s float64) {
	v.X += u.X * s
	v.Y += u.Y * s
	v.Z += u.Z * s
}
