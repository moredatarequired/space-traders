package main

import (
	"fmt"
	"math"
)

type Ship struct {
	Position []float64
	Velocity []float64
}

var dT float64 = 0.1

func (s *Ship) Accelerate(dv []float64) {
	for i, v := range dv {
		s.Velocity[i] += v * dT
	}
}

func (s *Ship) Move() {
	for i, v := range s.Velocity {
		s.Position[i] += v * dT
	}
}

func (s *Ship) String() string {
	return fmt.Sprintf("at %v moving at %v km/s", s.Position, s.Velocity)
}

func Norm(k float64, fs ...float64) float64 {
	total := 0.0
	for _, f := range fs {
		total += math.Pow(math.Abs(f), k)
	}
	return math.Pow(total, 1 / k)
}

// dv is the maximum total Delta-V than be expended.
func (s *Ship) Approach(t *Ship, dv float64) []float64 {
	diff := []float64{0, 0 , 0}
	for i, p := range s.Position {
		c := p - t.Position[i]
		v := s.Velocity[i] - t.Velocity[i]
		diff[i] = 2*dv*c - v*v
	}
	acc := []float64{0, 0, 0}
	k := Norm(1, diff...)
	for i, d := range diff {
		acc[i] = -d * dv / k
	}
	return acc
}

func main() {
	hero := &Ship{[]float64{1, 4, 9}, []float64{1, 1, 1}}
	foe := &Ship{[]float64{0, 0, 0}, []float64{0, 0, 0}}
	for i := 0; i < 1000; i++ {
		acc := hero.Approach(foe, 5)
		hero.Accelerate(acc)
		hero.Move()
		if 0 == i % 100 {
			fmt.Println(hero)
		}
	}
}
