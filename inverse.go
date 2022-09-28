package ga3d

// Inverse gives the inverse of a multivector.
func Inverse(m Multivector) Multivector {
	mc := m.copy()
	n := numerator(m)
	d := denominator(mc)
	return ScalarDivision(n, d)
}

// numerator to be defined.
func numerator(m Multivector) Multivector {
	mc := m.copy()
	as := AmplitudeSquared(m)
	asr := Reverse(as)
	cc := CliffordConjugation(mc)

	return GeometricProduct(cc, asr)
}

// denominator to be defined.
func denominator(m Multivector) int64 {
	// TODO: change from int64 to bit.int, since numerators can grow past the size of an int64.
	// TODO: also check throughout the code base to change from ints to big.ints.
	mc := m.copy()
	r := Rationalize(mc)
	return r.E0.Num().Int64()
}
