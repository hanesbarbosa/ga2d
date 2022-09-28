package ga3d

import "math/big"

// Reverse of a multivector.
func Reverse(m Multivector) Multivector {
	// pf = positive factor
	pf := big.NewRat(1, 1)
	// nf = negative factor
	nf := big.NewRat(-1, 1)

	// 1, 1, 1, -1
	m.E0.Mul(pf, m.E0)
	m.E1.Mul(pf, m.E1)
	m.E2.Mul(pf, m.E2)
	m.E12.Mul(nf, m.E12)

	return m
}
