package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests the addition of two multivectors.
func TestAddition(t *testing.T) {
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
	m1 := NewMultivector(c1[0], c1[1], c1[2], c1[3], c1[4], c1[5], c1[6], c1[7])
	m2 := NewMultivector(c2[0], c2[1], c2[2], c2[3], c2[4], c2[5], c2[6], c2[7])
	// Calculate addition.
	ms := Addition(m1, m2)

	if strings.Compare(ms.E0.RatString(), "9") != 0 ||
		strings.Compare(ms.E1.RatString(), "9") != 0 ||
		strings.Compare(ms.E2.RatString(), "9") != 0 ||
		strings.Compare(ms.E3.RatString(), "9") != 0 ||
		strings.Compare(ms.E12.RatString(), "9") != 0 ||
		strings.Compare(ms.E13.RatString(), "9") != 0 ||
		strings.Compare(ms.E23.RatString(), "9") != 0 ||
		strings.Compare(ms.E123.RatString(), "9") != 0 {
		t.Errorf("error calculating addition")
	}
}
