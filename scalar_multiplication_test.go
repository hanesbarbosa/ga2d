package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar multiplication is calculated properly.
func TestSMul(t *testing.T) {
	// Coefficients as integer numbers.
	n1 := []int64{40, -29, 29, 33}
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
	// Scalar.
	n := new(big.Rat)
	n.SetString("2/1")

	m.SMul(n)

	if strings.Compare(m.E0.RatString(), "80") != 0 ||
		strings.Compare(m.E1.RatString(), "-58") != 0 ||
		strings.Compare(m.E2.RatString(), "58") != 0 ||
		strings.Compare(m.E12.RatString(), "66") != 0 {
		t.Errorf("wrong results for the scalar multiplication calculation")
	}
}
