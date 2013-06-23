package starmap

import (
	"testing"
)

func TestDistance(t *testing.T) {
	star1, star2 := &Star{X:3, Y:2, Z:-3}, &Star{X:0, Y:2, Z:1}
	if d := star1.Distance(star2); d != 5 {
		t.Errorf("s1.Distance(s2) = %v, want 5", d)
	}
}

func TestSquaredDistance(t *testing.T) {
	star1, star2 := &Star{X:1, Y:2, Z:-3}, &Star{X:0, Y:3, Z:1}
	if d := star1.squaredDistance(star2); d != 18 {
		t.Errorf("s1.Distance(s2) = %v, want 18", d)
	}
}
