package lib

import (
	"testing"
)

func TestMove(t *testing.T) {
	s := &Ship{}
	s.Move(1)
	if s.Position != [3]float64{0, 0, 0} {
		t.Error("Ship with zero velocity moved.")
	}
	s.Velocity[1] = 2.0
	s.Move(3)
	if p := [3]float64{0, 6.0, 0}; p != s.Position {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
	s.Velocity[2] = -10.0
	s.Move(0.1)
	if p := [3]float64{0, 6.2, -1.0}; p != s.Position {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
}

func TestAcceleration(t *testing.T) {
	s := &Ship{}
	s.Acceleration[0], s.Acceleration[1] = 1.0, 0.5
	for i := 0; i < 10; i++ { s.Move(1.0) }
	if p := [3]float64{50, 25, 0}; p != s.Position {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
	// Movement should work regardless of the time delta.
	s.Velocity[0], s.Velocity[1] = 0, 0
	for i := 0; i < 40; i++ { s.Move(0.25) }
	if p := [3]float64{100, 50, 0}; p != s.Position {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
}

func TestSquaredDistance(t *testing.T) {
	p1, p2 := [3]float64{0, 1, 2}, [3]float64{-1, 3, 5}
	if d := SquaredDistance(p1[:], p2[:]); d != 1 + 4 + 9 {
		t.Errorf("Expected distance of %v; got %v.", 14, d)
	}
	s1, s2 := &Ship{}, &Ship{}
	s2.Position[1] = -2
	if d := s1.SquaredDistance(s2); d != 4 {
		t.Errorf("Expected distance of 4.0; got %v.", d)
	}
}
