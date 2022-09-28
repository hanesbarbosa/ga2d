package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests the addition of two multivectors.
func TestAdd(t *testing.T) {
	// Coefficients as integer numbers.
	n1 := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	n2 := []int64{8, 7, 6, 5, 4, 3, 2, 1}
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
	m1 := NewMultivector(c1[0], c1[1], c1[2], c1[3])
	m2 := NewMultivector(c2[0], c2[1], c2[2], c2[3])
	// Calculate addition.
	m1.Add(m2)

	if strings.Compare(m1.E0.RatString(), "9") != 0 ||
		strings.Compare(m1.E1.RatString(), "9") != 0 ||
		strings.Compare(m1.E2.RatString(), "9") != 0 ||
		strings.Compare(m1.E12.RatString(), "9") != 0 {
		t.Errorf("error calculating addition")
	}
}
