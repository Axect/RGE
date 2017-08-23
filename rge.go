package rge

import "math"

// Single Running
func (R *RGE) Run(mt, xi float64) {
	var B Beta
	B.InputFormula(R, mt, xi)
	B.BetaFunction()

	// Real Running
	R.lH += h * B.BlH
	R.yt += h * B.Byt
	R.g1 += h * B.Bg1
	R.g2 += h * B.Bg2
	R.g3 += h * B.Bg3
	R.t += h
	R.phi = math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	R.G -= h * B.gamma / (1 + gamma)
}
