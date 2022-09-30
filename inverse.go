package ga3d

import "math/big"

// Inverse gives the inverse of a multivector.
func (m Multivector) Inverse() Multivector {
	mc := m.copy()
	n := m.numerator()
	d := mc.denominator()

	return n.SDiv(d)
}

// numerator to be defined.
func (m Multivector) numerator() Multivector {
	mc := m.copy()
	mc.AmplitudeSquared()
	mc.Reverse()
	m.CliffordConjugation()

	return m.GeometricProduct(mc)
}

// denominator to be defined.
func (m Multivector) denominator() *big.Int {
	mc := m.copy()
	mc.Rationalize()

	return mc.E0.Num()
}
