package starmap

import (
	"testing"
)

type Foo struct { X, Y, Z float64 }
func (s *Foo) Position() (x, y, z float64) { return s.X, s.Y, s.Z }

type Bar struct { X, Y, Z float64 }
func (s *Bar) Position() (x, y, z float64) { return s.X, s.Y, s.Z }

func TestSquaredDistance(t *testing.T) {
	foo, bar := &Foo{1, 2, -3}, &Bar{0, 3, 1}
	if d := SquaredDistance(foo, bar); d != 18 {
		t.Errorf("s1.Distance(s2) = %v, want 18", d)
	}
}

func TestDistance(t *testing.T) {
	bar, foo := &Bar{3, 2, -3}, &Foo{0, 2, 1}
	if d := Distance(bar, foo); d != 5 {
		t.Errorf("s1.Distance(s2) = %v, want 5", d)
	}
}
