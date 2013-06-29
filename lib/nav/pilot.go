package lib

import (
	"github.com/moredatarequired/optbench"
	"fmt"
	"math"
	"math/rand"
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

func Distance(a, b []float64) float64 {
	var diff []float64
	for i, v := range a {
		diff = append(diff, v - b[i])
	}
	return Norm(2, diff...)
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
		integral = integral * 0.99 + e * dT
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
			output = append(output, controller[k](s.Position[k], v))
		}
		n := Norm(1, output...)
		if n <= dv {
			return output
		}
		var acc []float64
		for _, o := range output {
			acc = append(acc, o * dv / n)
		}
		return acc
	}
}

func (s *Ship) RunFrom(t []float64, dv float64) {
	var vector []float64
	for i, p := range s.Position {
		vector = append(vector, p - t[i])
	}
	n := Norm(1, vector...)
	var acc []float64
	for _, v := range vector {
		acc = append(acc, -v * dv / n)
	}
	s.Accelerate(acc)
	s.Move()
}

func CrossProduct(u, v []float64) []float64 {
	return []float64{u[1]*v[2] - u[2]*v[1], u[2]*v[0] - u[0]*v[2], u[0]*v[1] - u[1]*v[0]}
}

func PerpendicularNearest(u, v []float64) []float64 {
	return CrossProduct(u, CrossProduct(v, u))
}

func (s *Ship) RunAround( t []float64, dv float64) {
	var vector []float64
	for i, p := range s.Position {
		vector = append(vector, p - t[i])
	}
	vector = PerpendicularNearest(vector, s.Velocity)
	n := Norm(1, vector...)
	var acc []float64
	for _, v := range vector {
		acc = append(acc, v * dv / n)
	}
	s.Accelerate(acc)
	s.Move()
}

func RandomShip(b, c float64) *Ship {
	return &Ship{
		[]float64{rand.Float64() * b, rand.Float64() * b, rand.Float64() * b},
		[]float64{rand.Float64() * c, rand.Float64() * c, rand.Float64() * c}}
}

func FlightGame(p, i, d float64) float64 {
	points := 0.0
	targets, time := 2, 40.0
	for k := 0; k < targets; k++ {
		hero := &Ship{[]float64{0, 0, 0}, []float64{0, 0, 0}}
		foe := &Ship{[]float64{10, 10, 10}, []float64{0, 0, 0}}
		c := NewMotionController(hero, 1.05, p, i, d)
		for ticks := 0.0; ticks * dT < time; ticks += 1 {
			switch k % 3 {
			case 0: foe.Move()
			case 1: foe.RunFrom(hero.Position, 1)
			case 2: foe.RunAround(hero.Position, 1)
			}
			acc := c(foe.Position)
			hero.Accelerate(acc)
			hero.Move()
			if d := Distance(hero.Position, foe.Position); d < 1 {
				points += dT
			} else {
				points += dT / d
			}
		}
		//fmt.Printf("Enemy %v shot down at time %v (position %v)\n", k, ticks, foe.Position)
	}
	return points / (float64(targets) * time)
}

func Rand(a, b float64) float64 {
	return rand.Float64() * (b - a) + a
}

func main() {
	pop := optbench.NewPopulation(2, 100)
	fn := func (xs []float64) float64 {
		return -FlightGame(-xs[0], 0.0, -xs[1])
	}

	for k := 0; k < 10000; k++ {
		best := pop.Evaluate(fn)
		member := pop.Members[len(pop.Members) - 1]
		ratio := member.Genes[0] / member.Genes[1]
		fmt.Printf("Gen %v got to %.5f/%.5f with %v(%v)\n", k, -best,
			-pop.Fittest, member.Genes, ratio)
		optbench.Epoch(pop)
	}
	// The result of this experiment was that an ideal PID controller is
	// p, i, d := 0.2, 0.0, 0.67 (to 2 significant figures).
}
