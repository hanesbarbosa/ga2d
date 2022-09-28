package ga3d

import "math/big"

// Addition sums two multivectors
func Addition(m1, m2 Multivector) Multivector {
	// Multivector to be returned.
	m := NewMultivector(new(big.Rat), new(big.Rat), new(big.Rat), new(big.Rat))

	m.E0.Add(m1.E0, m2.E0)
	m.E1.Add(m1.E1, m2.E1)
	m.E2.Add(m1.E2, m2.E2)
	m.E12.Add(m1.E12, m2.E12)

	return m
}
