package lib

type Ship struct {
	Position Vector
	Velocity Vector
	Acceleration Vector
}

// Move the ship over t seconds.
func (s *Ship) Move(t float64) {
	p, v, a := &s.Position, &s.Velocity, &s.Acceleration
	// Who needs function overloading? (me)
	s.Position = *p.Plus(v.Times(t).Plus(a.Times(t*t/2)))
	s.Velocity = *v.Plus(a.Times(t))
}

func (s1 *Ship) SquaredDistance(s2 *Ship) float64 {
	return s1.Position.SquaredDistance(&s2.Position)
}

func (s1 *Ship) Distance(s2 *Ship) float64 {
	return s1.Position.Distance(&s2.Position)
}

// Adopt the maximum acceleration away from the given point.
func (s *Ship) Flee(p *Vector, a float64) {
	dir := s.Position.Minus(p)
	s.Acceleration = *dir.ScaleTo(a)
}

// Return the perpendicular to u that lies nearest v.
func PerpendicularNearest(u, v *Vector) *Vector {
	return u.Cross(v.Cross(u))
}

// Set the acceleration of s in order to circle t.
func (s *Ship) CircleTarget(t *Ship, a float64) {
	v := s.Velocity.Minus(&t.Velocity)
	p := s.Position.Minus(&t.Position)
	s.Acceleration = *PerpendicularNearest(p, v).ScaleTo(a)
}
