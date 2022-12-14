package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the reverse is calculated properly.
func TestReverse(t *testing.T) {
	// Numerators.
	n := []int64{40, -29, 29, 33}
	// Rational coefficients to create the multivector.
	var c []*big.Rat
	// Generate rational numbers.
	for i := 0; i < len(n); i++ {
		c = append(c, big.NewRat(n[i], int64(1)))
	}
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123.
	m := NewMultivector(c[0], c[1], c[2], c[3])

	m.Reverse()

	// 40e0 + 29e1 + -29e2 + 25e3 + 33e12 + -29e13 + 32e23 + 28e123
	if strings.Compare(m.E0.RatString(), "40") != 0 ||
		strings.Compare(m.E1.RatString(), "-29") != 0 ||
		strings.Compare(m.E2.RatString(), "29") != 0 ||
		strings.Compare(m.E12.RatString(), "-33") != 0 {
		t.Errorf("wrong results for the reverse calculation")
	}
}
