package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the clifford conjugation is calculated properly.
func TestCliffordConjugation(t *testing.T) {
	// Coefficients as integer numbers.
	n1 := []int64{2, 3, 4, 6}
	// Coefficients as rational numbers.
	var c1 []*big.Rat
	// Denominator.
	d := int64(1)
	// Generate rational coefficients.
	for i := 0; i < len(n1); i++ {
		// Coefficients for multivector.
		c1 = append(c1, big.NewRat(n1[i], d))
	}
	// Generate multivectors.
	m := NewMultivector(c1[0], c1[1], c1[2], c1[3])

	m.CliffordConjugation()

	if strings.Compare(m.E0.String(), "2/1") != 0 ||
		strings.Compare(m.E1.String(), "-3/1") != 0 ||
		strings.Compare(m.E2.String(), "-4/1") != 0 ||
		strings.Compare(m.E12.String(), "-6/1") != 0 {
		t.Errorf("wrong results for the clifford conjugation calculation")
	}
}
