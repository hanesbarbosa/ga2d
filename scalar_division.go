package ga3d

import "math/big"

// SDiv divides a multivector by a scalar.
func (m Multivector) SDiv(d *big.Int) Multivector {
	f := new(big.Rat).SetFrac(big.NewInt(1), d)
	m.E0.Mul(m.E0, f)
	m.E1.Mul(m.E1, f)
	m.E2.Mul(m.E2, f)
	m.E12.Mul(m.E12, f)

	return m
}
