package ga3d

// AmplitudeSquared gives the amplitude square of a multivector.
func AmplitudeSquared(m Multivector) Multivector {
	// Multivector to be returned.
	mc := m.Copy()
	ccm := CliffordConjugation(m)
	return GeometricProduct(mc, ccm)
}
