package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the geometric product between two multivectors are calculated properly.
func TestGeometricProduct(t *testing.T) {
	// Coefficients as integer numbers.
	n1 := []int64{40, -29, 29, -25, 33, -29, 32, -28}
	n2 := []int64{113, -84, 89, -65, 131, -107, 117, -93}
	// Coefficients as rational numbers.
	var c1, c2 []*big.Rat
	// Denominator.
	d := int64(1)
	// Generate rational coefficients.
	for i := 0; i < len(n1); i++ {
		// Coefficients for first multivector.
		c1 = append(c1, big.NewRat(n1[i], d))
		// Coefficients for second multivector.
		c2 = append(c2, big.NewRat(n2[i], d))
	}
	// Generate multivectors.
	m1 := NewMultivector(c1[0], c1[1], c1[2], c1[3], c1[4], c1[5], c1[6], c1[7])
	m2 := NewMultivector(c2[0], c2[1], c2[2], c2[3], c2[4], c2[5], c2[6], c2[7])

	m := GeometricProduct(m1, m2)

	if strings.Compare(m.E0.RatString(), "-2612") != 0 ||
		strings.Compare(m.E1.RatString(), "-2037") != 0 ||
		strings.Compare(m.E2.RatString(), "12348") != 0 ||
		strings.Compare(m.E3.RatString(), "2524") != 0 ||
		strings.Compare(m.E12.RatString(), "12938") != 0 ||
		strings.Compare(m.E13.RatString(), "-2914") != 0 ||
		strings.Compare(m.E23.RatString(), "13417") != 0 ||
		strings.Compare(m.E123.RatString(), "-12701") != 0 {
		t.Errorf("error calculating geometric")
	}
}
