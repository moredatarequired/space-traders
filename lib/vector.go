// Implements operations on three dimensional vectors consisting of float64.

package lib

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

// Assign the values in u to v.
func (v *Vector) Set(u *Vector) {
	v.X, v.Y, v.Z = u.X, u.Y, u.Z
}

func (v *Vector) Dot(u *Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v *Vector) SquaredLength() float64 {
	return v.Dot(v)
}

func (v *Vector) SquaredDistance(u *Vector) float64 {
	x, y, z := u.X - v.X, u.Y - v.Y, u.Z - v.Z
	return x*x + y*y + z*z
}

// The 2-norm, or Euclidean length.
func (v *Vector) Len() float64 {
	return math.Pow(v.SquaredLength(), 0.5)
}

func (v *Vector) Distance(u *Vector) float64 {
	return math.Pow(v.SquaredDistance(u), 0.5)
}

// Produce the vector c*v, where c is a scalar.
func (v *Vector) Times(c float64) *Vector {
	return &Vector{v.X * c, v.Y * c, v.Z * c}
}

// Produce a vector parallel to v with length l
func (v *Vector) ScaleTo(l float64) *Vector {
	s := l / v.Len()
	return v.Times(s)
}

// Produce the unit vector along v.
func (v *Vector) Unit() *Vector {
	return v.ScaleTo(1)
}

// Produce the vector (v + u).
func (v *Vector) Plus(u *Vector) *Vector {
	return &Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// Produce the vector (v - u)
func (v *Vector) Minus(u *Vector) *Vector {
	return &Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// Produce the cross product (v * u)
func (v *Vector) Cross(u *Vector) *Vector {
	return &Vector{v.Y*u.Z - v.Z*u.Y, v.Z*u.X - v.X*u.Z, v.X*u.Y - v.Y*u.X}
}
