package ga3d

import (
	"math/big"
)

// GeometricProduct yields the geometric product of two multivectors.
func (m Multivector) GeometricProduct(m2 Multivector) Multivector {
	// Clone source multivector to preserve values through operations.
	m1 := m.copy()
	// Intermediary variable to store operands before final computation.
	c := [4]*big.Rat{
		big.NewRat(0, 1),
		big.NewRat(0, 1),
		big.NewRat(0, 1),
		big.NewRat(0, 1),
	}

	// 1st Coefficient
	// m.e0 = (m1.e0 * m2.e0) + (m1.e1 * m2.e1) + (m1.e2 * m2.e2) - (m1.e12 * m2.e12)

	// (m1.e0 * m2.e0)
	c[0].Mul(m1.E0, m2.E0)
	// (m1.e1 * m2.e1)
	c[1].Mul(m1.E1, m2.E1)
	// (m1.e2 * m2.e2)
	c[2].Mul(m1.E2, m2.E2)
	// (m1.e12 * m2.e12)
	c[3].Mul(m1.E12, m2.E12)
	// m.e0 = c[0] + c[1] + c[2] - c[3]
	m.E0.Add(c[0], c[1])
	m.E0.Add(m.E0, c[2])
	m.E0.Sub(m.E0, c[3])

	// 2nd Coefficient
	// m.e1 = (m1.e0 * m2.e1) + (m1.e1 * m2.e0) - (m1.e2 * m2.e12) + (m1.e12 * m2.e2)

	// (m1.e0 * m2.e1)
	c[0].Mul(m1.E0, m2.E1)
	// (m1.e1 * m2.e0)
	c[1].Mul(m1.E1, m2.E0)
	// (m1.e2 * m2.e12)
	c[2].Mul(m1.E2, m2.E12)
	// (m1.e12 * m2.e2)
	c[3].Mul(m1.E12, m2.E2)
	// m.e1 = c[0] + c[1] - c[2] + c[3]
	m.E1.Add(c[0], c[1])
	m.E1.Sub(m.E1, c[2])
	m.E1.Add(m.E1, c[3])

	// 3rd Coefficient
	// m.e2 = (m1.e0 * m2.e2) + (m1.e1 * m2.e12) + (m1.e2 * m2.e0) - (m1.e12 * m2.e1)

	// (m1.e0 * m2.e2)
	c[0].Mul(m1.E0, m2.E2)
	// (m1.e1 * m2.e12)
	c[1].Mul(m1.E1, m2.E12)
	// (m1.e2 * m2.e0)
	c[2].Mul(m1.E2, m2.E0)
	// (m1.e12 * m2.e1)
	c[3].Mul(m1.E12, m2.E1)
	// m.e2 = c[0] + c[1] + c[2] - c[3]
	m.E2.Add(c[0], c[1])
	m.E2.Add(m.E2, c[2])
	m.E2.Sub(m.E2, c[3])

	// 4th Coefficient
	// m.e12 = (m1.e0 * m2.e12) + (m1.e1 * m2.e2) - (m1.e2 * m2.e1) + (m1.e12 * m2.e0)

	// (m1.e0 * m2.e12)
	c[0].Mul(m1.E0, m2.E12)
	// (m1.e1 * m2.e2)
	c[1].Mul(m1.E1, m2.E2)
	// (m1.e2 * m2.e1)
	c[2].Mul(m1.E2, m2.E1)
	// (m1.e12 * m2.e0)
	c[3].Mul(m1.E12, m2.E0)
	// m.e12 = c[0] + c[1] - c[2] + c[3]
	m.E12.Add(c[0], c[1])
	m.E12.Sub(m.E12, c[2])
	m.E12.Add(m.E12, c[3])

	return m
}
