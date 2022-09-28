package ga3d

// Rationalize is a function that rationalizes a multivector.
func Rationalize(m Multivector) Multivector {
	asm := AmplitudeSquared(m)
	asmc := asm.Copy()
	asrm := Reverse(asm)
	return GeometricProduct(asmc, asrm)
}
