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
	// Since acceleration is overwritten, use it as "scratch" space.
	s.Acceleration = s.Position
	s.Acceleration.MinusInPlace(p)
	if s.Acceleration.IsZero() {
		s.Acceleration.X = a
	} else {
		s.Acceleration.ScaleToInPlace(a)
	}
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

// Set the acceleration of s in order to circle t.
func (s *Ship) CircleTarget(t *Ship, a float64) {
	v := s.Velocity
	v.MinusInPlace(&t.Velocity)
	// Since acceleration is overwritten, use it as "scratch" space.
	s.Acceleration = t.Position
	s.Acceleration.MinusInPlace(&s.Position)
	PerpendicularNearestInPlace(&s.Acceleration, &v)
	s.Acceleration.ScaleToInPlace(a)
	if s.Acceleration.IsZero() {
		s.Acceleration.X = a
	}
}
