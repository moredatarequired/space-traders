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
		acc[i] = -0.9 * d * dv / k
	}
	return acc
}

func NewPIDController(p, i, d float64) (func (v, t float64) float64) {
	integral := 0.0
	last_e := 0.0
	return func (v, t float64) float64 {
		e := v - t
		integral += e * dT
		de := (e - last_e) / dT
		last_e = e
		return p * e + i * integral + d * de
	}
}

func NewMotionController(s *Ship, dv, p, i, d float64) (func (t []float64) []float64) {
	var controller [](func (v, t float64) float64) = nil
	for _ = range s.Position {
		controller = append(controller, NewPIDController(p, i, d))
	}
	return func (t []float64) []float64 {
		var output []float64
		for k, v := range t {
			desired := controller[k](s.Position[k], v)
			output = append(output, math.Max(-dv, math.Min(dv, desired)))
		}
		return output
	}
}

func main() {
	hero := &Ship{[]float64{1, 4, 9}, []float64{1, 1, 1}}
	foe := &Ship{[]float64{0, 0, 0}, []float64{0, 5, 0}}
	c := NewMotionController(hero, 5, -1, 0, -1)
	for i := 0; i < 250; i++ {
		acc := c(foe.Position)
		if 0 == i % 10 {
			fmt.Printf("at %v/%v, on vector %v\n", hero.Position, foe.Position, acc)
		}
		hero.Accelerate(acc)
		hero.Move()
		foe.Move()
	}
}
