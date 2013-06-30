package lib

import (
	"testing"
)

func TestMove(t *testing.T) {
	s, p := &Ship{}, Vector{}
	s.Move(1)
	if s.Position != p {
		t.Error("Ship with zero velocity moved to %v.", p)
	}
	s.Velocity.Y = 2.0
	p.Y = 6.0
	s.Move(3)
	if s.Position != p {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
	s.Velocity.Z = -10.0
	p.Y, p.Z = 6.2, -1.0
	s.Move(0.1)
	if s.Position != p {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
}

func BenchmarkMove(b *testing.B) {
	s := &Ship{Acceleration:Vector{0.1, 0.2, 0.3}}
	for i := 0; i < b.N; i++ {
		s.Move(0.01)
	}
}

// func TestAcceleration(t *testing.T) {
//	s := &Ship{}
//	s.Acceleration[0], s.Acceleration[1] = 1.0, 0.5
//	for i := 0; i < 10; i++ { s.Move(1.0) }
//	if p := [3]float64{50, 25, 0}; p != s.Position {
//		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
//	}
//	// Movement should work regardless of the time delta.
//	s.Velocity[0], s.Velocity[1] = 0, 0
//	for i := 0; i < 40; i++ { s.Move(0.25) }
//	if p := [3]float64{100, 50, 0}; p != s.Position {
//		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
//	}
// }

// func TestSquaredDistance(t *testing.T) {
//	p1, p2 := [3]float64{0, 1, 2}, [3]float64{-1, 3, 5}
//	if d := SquaredDistance(p1[:], p2[:]); d != 1 + 4 + 9 {
//		t.Errorf("Expected distance of %v; got %v.", 14, d)
//	}
//	s1, s2 := &Ship{}, &Ship{}
//	s2.Position[1] = -2
//	if d := s1.SquaredDistance(s2); d != 4 {
//		t.Errorf("Expected distance of 4.0; got %v.", d)
//	}
// }
