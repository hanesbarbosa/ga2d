package ga3d

// Rationalize is a function that rationalizes a multivector.
func (m Multivector) Rationalize() Multivector {
	m.AmplitudeSquared()
	mc := m.copy()
	m.Reverse()

	return m.GeometricProduct(mc)
}
