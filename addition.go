package ga3d

// Add sums two multivectors
func (m Multivector) Add(m2 Multivector) Multivector {
	m.E0.Add(m.E0, m2.E0)
	m.E1.Add(m.E1, m2.E1)
	m.E2.Add(m.E2, m2.E2)
	m.E12.Add(m.E12, m2.E12)

	return m
}
