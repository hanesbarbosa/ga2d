package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the geometric product between two multivectors are calculated properly.
func TestGeometricProduct(t *testing.T) {
	// Coefficients as integer numbers.
	n1 := []int64{40, -29, 29, -25}
	n2 := []int64{113, -84, 89, -65}
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
	m := NewMultivector(c1[0], c1[1], c1[2], c1[3])
	m2 := NewMultivector(c2[0], c2[1], c2[2], c2[3])

	m.GeometricProduct(m2)

	if strings.Compare(m.E0.RatString(), "7912") != 0 ||
		strings.Compare(m.E1.RatString(), "-6977") != 0 ||
		strings.Compare(m.E2.RatString(), "6622") != 0 ||
		strings.Compare(m.E12.RatString(), "-5570") != 0 {
		t.Errorf("error calculating geometric")
	}
}
