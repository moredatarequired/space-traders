package lib

import (
	"testing"
	"math"
)

func TestDotProduct(t *testing.T) {
	u, v := Vector{}, Vector{1, 2, 3}
	if r := u.Dot(&v); r != 0 {
		t.Errorf("Non-zero result %v, expected zero.", r)
	}
	u = Vector{-1, 0, 2}
	if r := u.Dot(&v); r != 5 {
		t.Errorf("Expected dot product 5.0, got %v", r)
	}
	if u.Dot(&v) != v.Dot(&u) { t.Errorf("Dot products do not match.") }
}

func BenchmarkDotProduct(b *testing.B) {
	v, u := &Vector{1.1, 0.3, 15.6}, &Vector{-9.3, 1.3, 7.0}
	for i := 0; i < b.N; i++ {
		v.Dot(u)
	}
}

func TestSquaredLength(t *testing.T) {
	if l := (&Vector{}).SquaredLength(); l != 0 {
		t.Errorf("Zero vector has non-zero squared length %v", l)
	}
	if l, d := (&Vector{-2, 1, 4}).SquaredLength(), 21.0; l != d {
		t.Errorf("Vector has squared length %v, expected %v", l, d)
	}
}

func BenchmarkSquaredLength(b *testing.B) {
	v := Vector{1.1, 0.3, 15.6}
	for i := 0; i < b.N; i++ {
		v.SquaredLength()
	}
}

func TestLength(t *testing.T) {
	if l := (&Vector{}).Length(); l != 0 {
		t.Errorf("Zero vector has non-zero length %v", l)
	}
	if l, d := (&Vector{-2, 3, 6}).Length(), 7.0; l != d {
		t.Errorf("Vector has length %v, expected %v", l, d)
	}
}

func BenchmarkLength(b *testing.B) {
	v := Vector{1.1, 0.3, 15.6}
	for i := 0; i < b.N; i++ {
		v.Length()
	}
}

func TestSquaredDistance(t *testing.T) {
	if l := (&Vector{}).SquaredDistance(&Vector{}); l != 0 {
		t.Errorf("Zero vectors have non-zero squared distance %v", l)
	}
	u, v, d := &Vector{-2, 1, 4}, &Vector{1, 0, 2}, 14.0;
	if l := u.SquaredDistance(v); l != d {
		t.Errorf("Vectors have squared distance %v, expected %v", l, d)
	}
}

func BenchmarkSquaredDistance(b *testing.B) {
	v, u := Vector{1.1, 0.3, 15.6}, &Vector{-9.3, 1.3, 7.0}
	for i := 0; i < b.N; i++ {
		v.SquaredDistance(u)
	}
}

func TestDistance(t *testing.T) {
	if l := (&Vector{}).Distance(&Vector{}); l != 0 {
		t.Errorf("Zero vectors have non-zero distance %v", l)
	}
	u, v, d := &Vector{-2, 4, 3}, &Vector{0, 1, -3}, 7.0
	if l := u.Distance(v); l != d {
		t.Errorf("Vectors have distance %v, expected %v", l, d)
	}
}

func BenchmarkDistance(b *testing.B) {
	v, u := Vector{1.1, 0.3, 15.6}, &Vector{-9.3, 1.3, 7.0}
	for i := 0; i < b.N; i++ {
		v.Distance(u)
	}
}

func TestMultiply(t *testing.T) {
	v, s := Vector{-2, 4, 3}, Vector{-1, 2, 1.5}
	if r := v.Times(0.5); *r != s {
		t.Errorf("Expected scaled vector %v, got %v", s, r)
	}
	if r := v.Times(0); *r != (Vector{}) {
		t.Errorf("Expected zero vector, got %v", r)
	}
	if r := v.Times(1); *r != v {
		t.Errorf("Expected (same) vector %v, got %v", v, r)
	}
}

func BenchmarkMultiply(b *testing.B) {
	v := &Vector{1.1, 0.3, 15.6}
	for i := 0; i < b.N; i++ {
		v.Times(3.2)
	}
}

func BenchmarkMultiplyInPlace(b *testing.B) {
	v := &Vector{1.1, 0.3, 15.6}
	for i := 0; i < b.N; i++ {
		v.TimesInPlace(1.00001)
	}
}

func TestScaleTo(t *testing.T) {
	v := &Vector{1.5, 0.3, 15.6}
	s := v.ScaleTo(5)
	if l := s.Length(); l != 5 {
		t.Errorf("Expected vector to scale to 5; got %v (%v)", s, l)
	}
	if v.Z != 15.6 {
		t.Error("Receiver changed during operation.")
	}
	s = (&Vector{5, 5, 5}).ScaleTo(math.Sqrt(3))
	if !s.Equal(&Vector{1, 1, 1}) {
		t.Errorf("Expected vector (1, 1, 1), got %v", s)
	}
}

func BenchmarkScaleTo(b *testing.B) {
	v := &Vector{1.1, 0.3, 15.6}
	for i := 0; i < b.N; i++ {
		v.ScaleTo(3.2)
	}
}

func BenchmarkScaleToInPlace(b *testing.B) {
	v := &Vector{1.1, 0.3, 15.6}
	for i := 0; i < b.N; i++ {
		v.ScaleToInPlace(3.2)
	}
}

func TestUnit(t *testing.T) {
	v := &Vector{1.5, 0.3, 15.6}
	s := v.Unit()
	if l := s.Length(); l != 1 {
		t.Errorf("Expected vector to scale to 1; got %v (%v)", s, l)
	}
	if v.Z != 15.6 {
		t.Error("Receiver changed during operation.")
	}
}

func TestAdd(t *testing.T) {
	u, v, expect := Vector{-2, 4, 3}, Vector{1, -2, 0}, Vector{-1, 2, 3}
	if r := u.Plus(&v); *r != expect {
		t.Errorf("Expected vector sum %v, got %v", expect, r)
	}
	if u.Y != 4 || v.Y != -2 {
		t.Error("Receiver changed during operation.")
	}
	if r := u.Plus(&Vector{}); *r != u {
		t.Errorf("Expected (same) vector %v, got %v", u, r)
	}
}

func BenchmarkAdd(b *testing.B) {
	v, u := &Vector{1.1, 0.3, 15.6}, &Vector{0.7, 19.0, 4.4}
	for i := 0; i < b.N; i++ {
		u.Plus(v)
	}
}

func BenchmarkAddInPlace(b *testing.B) {
	v, u := &Vector{1.1, 0.3, 15.6}, &Vector{0.7, 0.00019, 0.004}
	for i := 0; i < b.N; i++ {
		v.PlusInPlace(u)
	}
}

func TestSubtract(t *testing.T) {
	u, v, expect := Vector{-2, 4, 3}, Vector{1, -2, 0}, Vector{-3, 6, 3}
	if r := u.Minus(&v); *r != expect {
		t.Errorf("Expected vector difference %v, got %v", expect, r)
	}
	if u.Y != 4 || v.Y != -2 {
		t.Error("Receiver changed during operation.")
	}
	if r := u.Minus(&Vector{}); *r != u {
		t.Errorf("Expected (same) vector %v, got %v", u, r)
	}
}

func TestCross(t *testing.T) {
	u, v, expect := Vector{-2, 4, 3}, Vector{1, -2, 0}, Vector{6, 3, 0}
	if r := u.Cross(&v); *r != expect {
		t.Errorf("Expected vector cross product %v, got %v", expect, r)
	}
	if u.Y != 4 || v.Y != -2 {
		t.Error("Receiver changed during operation.")
	}
	if r := u.Cross(&u); *r != (Vector{}) {
		t.Errorf("Expected zero vector, got %v", r)
	}
	if r := u.Cross(&Vector{}); *r != (Vector{}) {
		t.Errorf("Expected zero vector, got %v", r)
	}
}

func BenchmarkCrossProduct(b *testing.B) {
	v, u := &Vector{1.1, 0.3, 15.6}, &Vector{0.7, 19.0, 4.4}
	for i := 0; i < b.N; i++ {
		u.Cross(v)
	}
}

func BenchmarkCrossProductInPlace(b *testing.B) {
	v, u := &Vector{1.1, 0.3, 15.6}, &Vector{0.7, 1.9, 0.4}
	for i := 0; i < b.N; i++ {
		v.CrossInPlace(u)
	}
}
