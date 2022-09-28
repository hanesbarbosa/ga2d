package ga3d

import (
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func initializeMultivector() Multivector {
	// Numerators.
	c := []int64{40, -29, 29, -25, 33, -29, 32, -28}
	// Rationals to create the multivector.
	var r []*big.Rat
	// Generate rational numbers.
	for i := 0; i < len(c); i++ {
		r = append(r, big.NewRat(c[i], int64(1)))
	}
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123.
	m := NewMultivector(r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7])
	return m
}

// Tests Inverse function
func TestInverse(t *testing.T) {
	m := initializeMultivector()
	m = Inverse(m)
	if strings.Compare(m.E0.RatString(), "70168/2313425") != 0 ||
		strings.Compare(m.E1.RatString(), "55739/2313425") != 0 ||
		strings.Compare(m.E2.RatString(), "-30363/2313425") != 0 ||
		strings.Compare(m.E3.RatString(), "50303/2313425") != 0 ||
		strings.Compare(m.E12.RatString(), "-37879/2313425") != 0 ||
		strings.Compare(m.E13.RatString(), "54491/2313425") != 0 ||
		strings.Compare(m.E23.RatString(), "-34752/2313425") != 0 ||
		strings.Compare(m.E123.RatString(), "-24324/2313425") != 0 {
		t.Errorf("wrong results for inverse")
	}
}

// Tests auxiliary function Numerator.
func TestNumerator(t *testing.T) {
	m := initializeMultivector()
	m = numerator(m)

	// 70168e0 + 55739e1 + -30363e2 + 50303e3 + -37879e12 + 54491e13 + -34752e23 + -24324e123
	if strings.Compare(m.E0.RatString(), "70168") != 0 ||
		strings.Compare(m.E1.RatString(), "55739") != 0 ||
		strings.Compare(m.E2.RatString(), "-30363") != 0 ||
		strings.Compare(m.E3.RatString(), "50303") != 0 ||
		strings.Compare(m.E12.RatString(), "-37879") != 0 ||
		strings.Compare(m.E13.RatString(), "54491") != 0 ||
		strings.Compare(m.E23.RatString(), "-34752") != 0 ||
		strings.Compare(m.E123.RatString(), "-24324") != 0 {
		t.Errorf("wrong results for numerator")
	}
}

// Tests auxiliary function Denominator.
func TestDenominator(t *testing.T) {
	m := initializeMultivector()
	e0 := denominator(m)

	// TODO: refactor of comparison.
	if strings.Compare(strconv.Itoa(int(e0)), "2313425") != 0 {
		t.Errorf("wrong result for coefficient")
	}
}
