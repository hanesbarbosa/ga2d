package ga3d

import "math/big"

// ScalarDivision divides a multivector by a scalar.
func ScalarDivision(m Multivector, n int64) Multivector {
	d := big.NewRat(int64(1), n)
	m.E0.Mul(m.E0, d)
	m.E1.Mul(m.E1, d)
	m.E2.Mul(m.E2, d)
	m.E12.Mul(m.E12, d)

	return m
}
