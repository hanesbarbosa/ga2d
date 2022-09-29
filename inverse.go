package ga3d

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
func (m Multivector) denominator() int64 {
	// TODO: change from int64 to big.int, since numerators can grow past the size of an int64.
	// TODO: also check throughout the code base to change from ints to big.ints.
	m.Rationalize()

	return m.E0.Num().Int64()
}
