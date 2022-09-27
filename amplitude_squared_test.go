package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar multiplication is calculated properly.
func TestAmplitudeSquared(t *testing.T) {
	// Numerators.
	n := []int64{40, -29, 29, -25, 33, -29, 32, -28}
	// Rational coefficients to create the multivector.
	var c []*big.Rat
	// Generate rational numbers.
	for i := 0; i < len(n); i++ {
		c = append(c, big.NewRat(n[i], int64(1)))
	}
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123.
	m := NewMultivector(c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7])
	m = AmplitudeSquared(m)

	if strings.Compare(m.E0.RatString(), "1463") != 0 ||
		strings.Compare(m.E1.RatString(), "0") != 0 ||
		strings.Compare(m.E2.RatString(), "0") != 0 ||
		strings.Compare(m.E3.RatString(), "0") != 0 ||
		strings.Compare(m.E12.RatString(), "0") != 0 ||
		strings.Compare(m.E13.RatString(), "0") != 0 ||
		strings.Compare(m.E23.RatString(), "0") != 0 ||
		strings.Compare(m.E123.RatString(), "-416") != 0 {
		t.Errorf("wrong results for the amplitude squared calculation")
	}
}
