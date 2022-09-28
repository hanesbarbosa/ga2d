package ga3d

import "math/big"

// SMul multiplies a multivector by a scalar.
func (m Multivector) SMul(s *big.Rat) Multivector {
	m.E0.Mul(m.E0, s)
	m.E1.Mul(m.E1, s)
	m.E2.Mul(m.E2, s)
	m.E12.Mul(m.E12, s)

	return m
}
