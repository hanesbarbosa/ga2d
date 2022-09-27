package ga3d

import (
	"math/big"
	"testing"
)

func TestNewMultivector(t *testing.T) {
	// Numerators.
	c := []int64{12, 23, 34, 45, 56, 67, 78, 89}
	// Rationals to create the multivector.
	var r []*big.Rat
	// Generate rational numbers.
	for i := 0; i < len(c); i++ {
		r = append(r, big.NewRat(c[i], int64(1)))
	}
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123.
	m := NewMultivector(r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7])
	// Check generated multivector.
	if m.E0.RatString() != "12" ||
		m.E1.RatString() != "23" ||
		m.E2.RatString() != "34" ||
		m.E3.RatString() != "45" ||
		m.E12.RatString() != "56" ||
		m.E13.RatString() != "67" ||
		m.E23.RatString() != "78" ||
		m.E123.RatString() != "89" {
		t.Errorf("wrong values for new multivector")
	}
}
