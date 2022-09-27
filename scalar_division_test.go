package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar division is calculated properly.
func TestScalarDivision(t *testing.T) {
	// Coefficients as integer numbers.
	n1 := []int64{2, 3, 4, 5, 6, 7, 8, 9}
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
	m := NewMultivector(c1[0], c1[1], c1[2], c1[3], c1[4], c1[5], c1[6], c1[7])
	// Scalar divisor.
	s := int64(2)

	m = ScalarDivision(m, s)

	if strings.Compare(m.E0.String(), "1/1") != 0 ||
		strings.Compare(m.E1.String(), "3/2") != 0 ||
		strings.Compare(m.E2.String(), "2/1") != 0 ||
		strings.Compare(m.E3.String(), "5/2") != 0 ||
		strings.Compare(m.E12.String(), "3/1") != 0 ||
		strings.Compare(m.E13.String(), "7/2") != 0 ||
		strings.Compare(m.E23.String(), "4/1") != 0 ||
		strings.Compare(m.E123.String(), "9/2") != 0 {
		t.Errorf("wrong results for the rational form of coefficients")
	}
}
