package ga3d

import (
	"math/big"
	"strings"
	"testing"
)

func initializeMultivector() Multivector {
	// Numerators.
	c := []int64{40, -29, 29, 33}
	// Rationals to create the multivector.
	var r []*big.Rat
	// Generate rational numbers.
	for i := 0; i < len(c); i++ {
		r = append(r, big.NewRat(c[i], int64(1)))
	}
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123.
	m := NewMultivector(r[0], r[1], r[2], r[3])
	return m
}

// Tests Inverse function
func TestInverse(t *testing.T) {
	m := initializeMultivector()
	m.Inverse()

	if strings.Compare(m.E0.RatString(), "40/1007") != 0 ||
		strings.Compare(m.E1.RatString(), "29/1007") != 0 ||
		strings.Compare(m.E2.RatString(), "-29/1007") != 0 ||
		strings.Compare(m.E12.RatString(), "-33/1007") != 0 {
		t.Errorf("wrong results for inverse")
	}
}

// Tests auxiliary function Numerator.
func TestNumerator(t *testing.T) {
	m := initializeMultivector()
	n := m.numerator()

	if strings.Compare(n.E0.RatString(), "40280") != 0 ||
		strings.Compare(n.E1.RatString(), "29203") != 0 ||
		strings.Compare(n.E2.RatString(), "-29203") != 0 ||
		strings.Compare(n.E12.RatString(), "-33231") != 0 {
		t.Errorf("wrong results for numerator")
	}
}

// Tests auxiliary function Denominator.
func TestDenominator(t *testing.T) {
	m := initializeMultivector()
	d := m.denominator()

	// TODO: refactor of comparison.
	if strings.Compare(d.String(), "1014049") != 0 {
		t.Errorf("wrong result for coefficient")
	}
}
