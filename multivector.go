package ga3d

import "math/big"

// Multivector organizes the basis elements of a 3d multivector.
type Multivector struct {
	E0   *big.Rat
	E1   *big.Rat
	E2   *big.Rat
	E3   *big.Rat
	E12  *big.Rat
	E13  *big.Rat
	E23  *big.Rat
	E123 *big.Rat
}

// NewMultivector creates a new multivector where all coefficients are rational numbers.
// Since the package already informs the dimension (i.e., 3d) we do not need to reinforce
// in the function name.
func NewMultivector(e0, e1, e2, e3, e12, e13, e23, e123 big.Rat) Multivector {
	return Multivector{
		E0:   &e0,
		E1:   &e1,
		E2:   &e2,
		E3:   &e3,
		E12:  &e12,
		E13:  &e13,
		E23:  &e23,
		E123: &e123,
	}
}
