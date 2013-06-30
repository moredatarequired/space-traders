package lib

type Ship struct {
	Position Vector
	Velocity Vector
	Acceleration Vector
	c Controller
}

// Move the ship over t seconds.
func (s *Ship) Move(t float64) {
	p, v, a := &s.Position, &s.Velocity, &s.Acceleration
	p.AddWithScaleInPlace(v, t)  // p += t*v
	p.AddWithScaleInPlace(a, 0.5*t*t)
	v.AddWithScaleInPlace(a, t)
}

func (s1 *Ship) SquaredDistance(s2 *Ship) float64 {
	return s1.Position.SquaredDistance(&s2.Position)
}

func (s1 *Ship) Distance(s2 *Ship) float64 {
	return s1.Position.Distance(&s2.Position)
}

type Controller interface {
	Redirect()
}

// Adopt the maximum acceleration away from the given point.
func (s *Ship) Flee(p *Vector, a float64) {
	// Since acceleration is overwritten, use it as "scratch" space.
	s.Acceleration = s.Position
	s.Acceleration.MinusInPlace(p)
	if s.Acceleration.IsZero() {
		s.Acceleration.X = a
	} else {
		s.Acceleration.ScaleToInPlace(a)
	}
}

type FleeController struct {
	s *Ship
	p *Vector
	a float64
}

func (f *FleeController) Redirect() {
	f.s.Flee(f.p, f.a)
}

// Return the perpendicular to u that lies nearest v.
func PerpendicularNearest(u, v *Vector) *Vector {
	return u.Cross(v.Cross(u))
}

// Return the perpendicular to u that lies nearest v.
func PerpendicularNearestInPlace(u, v *Vector) {
	v.CrossInPlace(u)
	u.CrossInPlace(v)
}

// Accelerate tangentially away from the target.
func (s *Ship) SpiralAway(t *Ship, a float64) {
	// Since acceleration is overwritten, use it as "scratch" space.
	s.Acceleration = s.Velocity
	s.Acceleration.MinusInPlace(&t.Velocity)
	p := t.Position
	p.MinusInPlace(&s.Position)
	s.Acceleration.RejectInPlace(&p)
	s.Acceleration.ScaleToInPlace(a)
	if s.Acceleration.IsZero() {
		s.Acceleration.X = a
	}
}

// "Orbits" the target. Orbit is circular when a = v^2 / r.
func (s *Ship) Circle(p *Vector, a float64) {
	s.Acceleration = *p
	s.Acceleration.MinusInPlace(&s.Position)
	s.Acceleration.ScaleToInPlace(a)
}
