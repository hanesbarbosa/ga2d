package ga3d

import "math/big"

// ScalarDivision divides a multivector by a scalar.
func ScalarDivision(m Multivector, n int64) Multivector {
	d := big.NewRat(int64(1), n)
	m.E0.Mul(m.E0, d)
	m.E1.Mul(m.E1, d)
	m.E2.Mul(m.E2, d)
	m.E3.Mul(m.E3, d)
	m.E12.Mul(m.E12, d)
	m.E13.Mul(m.E13, d)
	m.E23.Mul(m.E23, d)
	m.E123.Mul(m.E123, d)

	return m
}
