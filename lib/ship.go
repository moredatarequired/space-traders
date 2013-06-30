package lib

type Ship struct {
	Position Vector
	Velocity Vector
	Acceleration Vector
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

// Adopt the maximum acceleration away from the given point.
func (s *Ship) Flee(p *Vector, a float64) {
	// TODO: modify acceleration vector in place.
	dir := s.Position.Minus(p)
	s.Acceleration = *dir.ScaleTo(a)
}

// Return the perpendicular to u that lies nearest v.
func PerpendicularNearest(u, v *Vector) *Vector {
	return u.Cross(v.Cross(u))
}

// Set the acceleration of s in order to circle t.
func (s *Ship) CircleTarget(t *Ship, a float64) {
	// TODO: modify acceleration vector in place.
	v := s.Velocity.Minus(&t.Velocity)
	p := s.Position.Minus(&t.Position)
	s.Acceleration = *PerpendicularNearest(p, v).ScaleTo(a)
}
