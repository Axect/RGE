package RGE

import "math"

// RunCosmo calculate Cosmo from RGE
func (R RGE) RunCosmo(xi float64) Cosmo {
	var C Cosmo
	// Potential Formula
	C.V = (R.lH * math.Pow(R.G, 4) * math.Pow(R.phi, 4) / (4 * math.Pow(1+xi*math.Pow(R.G, 2)*math.Pow(R.phi, 2)/math.Pow(MpR, 2), 2))) / math.Pow(MpR, 4)
	return C
}
