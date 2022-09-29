package ga3d

// AmplitudeSquared gives the amplitude square of a multivector.
func (m Multivector) AmplitudeSquared() Multivector {
	mc := m.copy()
	mc.CliffordConjugation()
	return m.GeometricProduct(mc)
}
