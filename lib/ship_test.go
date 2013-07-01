package lib

import (
	"testing"
	"math"
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

func TestAcceleration(t *testing.T) {
	s := &Ship{}
	s.Acceleration = Vector{1, 0.5, 0}
	for i := 0; i < 10; i++ { s.Move(1.0) }
	if p := (Vector{50, 25, 0}); p != s.Position {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
	// Movement should work regardless of the time delta.
	s.Velocity = Vector{}
	for i := 0; i < 40; i++ { s.Move(0.25) }
	if p := (Vector{100, 50, 0}); p != s.Position {
		t.Errorf("Ship moved to %v; expected %v.", s.Position, p)
	}
}

func TestFlee(t *testing.T) {
	p, s := &Ship{Position:Vector{1, 1, 1}}, &Ship{Position:Vector{X:1}}
	s.Flee(&p.Position, math.Sqrt(2))
	acc := &Vector{0, -1, -1}
	if !s.Acceleration.Equals(acc) {
		t.Errorf("Ship should flee at %v; instead going at %v.",
			acc, s.Acceleration)
	}
}

func BenchmarkFlee(b *testing.B) {
	p, s := &Vector{X:1}, &Ship{}
	for i := 0; i < b.N; i++ {
		s.Flee(p, 5)
	}
}

func TestSpiralAway(t *testing.T) {
	p, s := &Ship{}, &Ship{Position:Vector{X:1}, Velocity:Vector{Y:1}}
	s.SpiralAway(p, 1.5)
	acc := &Vector{0, 1.5, 0}
	if !s.Acceleration.Equals(acc) {
		t.Errorf("Ship should circle at %v; instead going at %v.",
			acc, s.Acceleration)
	}
	s.Position.Y = 1
	s.SpiralAway(p, math.Sqrt(2))
	acc = &Vector{-1, 1, 0}
	if !s.Acceleration.Equals(acc) {
		t.Errorf("Ship should circle at %v; instead going at %v.",
			acc, s.Acceleration)
	}
}

func BenchmarkSpiralAway(b *testing.B) {
	p, s := &Ship{}, &Ship{Position:Vector{X:1}, Velocity:Vector{Y:1}}
	for i := 0; i < b.N; i++ {
		s.SpiralAway(p, 5)
	}
}

func TestCircle(t *testing.T) {
	p, s := &Ship{Position:Vector{1, 1, 1}}, &Ship{Position:Vector{X:1}}
	s.Circle(&p.Position, math.Sqrt(2))
	acc := &Vector{0, 1, 1}
	if !s.Acceleration.Equals(acc) {
		t.Errorf("Ship should flee at %v; instead going at %v.",
			acc, s.Acceleration)
	}
}

func BenchmarkCircle(b *testing.B) {
	p, s := &Vector{X:1}, &Ship{}
	for i := 0; i < b.N; i++ {
		s.Circle(p, 5)
	}
}

func TestMaintainDistance(t *testing.T) {
	f := &Ship{}
	for _, p := range []float64{1, 10, 100} {
		for _, d := range []float64{1, 10, 100} {
			s := &Ship{Position:Vector{X:p}, Velocity:Vector{Y:3}}
			for i := 0; i < 10000; i++ {
				s.MaintainDistance(f, 10, d)
				s.Move(0.1)
			}
			if math.Abs(s.Distance(f) / d - 1) > 0.25 {
				t.Errorf("Ship %v is out of bounds %v", s, d)
			}
		}
	}
}
