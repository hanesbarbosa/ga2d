package ga3d

import "math/big"

// Multivector organizes the basis elements of a 3d multivector.
type Multivector struct {
	E0  *big.Rat
	E1  *big.Rat
	E2  *big.Rat
	E12 *big.Rat
}

// New creates a new multivector where all coefficients are rational numbers.
// Since the package already informs the dimension (i.e., 3d) we do not need to reinforce
// in the function name.
func NewMultivector(e0, e1, e2, e12 *big.Rat) Multivector {
	return Multivector{
		E0:  e0,
		E1:  e1,
		E2:  e2,
		E12: e12,
	}
}

// Copy is a function that copies the coefficient values into a new multivector
// to avoid errors by pointer reference.
func (m Multivector) copy() Multivector {
	// Multivector copy (cloned).
	mc := Multivector{
		E0:  new(big.Rat),
		E1:  new(big.Rat),
		E2:  new(big.Rat),
		E12: new(big.Rat),
	}

	mc.E0.SetString(m.E0.RatString())
	mc.E1.SetString(m.E1.RatString())
	mc.E2.SetString(m.E2.RatString())
	mc.E12.SetString(m.E12.RatString())

	return mc
}
