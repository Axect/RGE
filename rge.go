package rge

import "math"

type Run interface {
	Run(float64, float64) // Run RGE
	Copy()                // Copy RGE to RGE
	InputFormula(RGE, float64, float64)
}

// Single Running - You can change Numerical Integration method here (Default: Euler)
func (R *RGE) Run(mt, xi float64) {
	var B Beta

	B.InputFormula(*R, mt, xi) // reverse pointer
	B.BetaFunction()

	// Real Running
	R.lH += h * B.BlH
	R.yt += h * B.Byt
	R.g1 += h * B.Bg1
	R.g2 += h * B.Bg2
	R.g3 += h * B.Bg3
	R.t += h
	R.phi = math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	R.G -= h * B.gamma / (1 + B.gamma)
}

// Copy protects RGE value which is in Container
func (R RGE) Copy() RGE {
	var NR RGE

	NR.t = R.t
	NR.lH = R.lH
	NR.yt = R.yt
	NR.g1 = R.g1
	NR.g2 = R.g2
	NR.g3 = R.g3
	NR.phi = R.phi
	NR.G = R.G

	return NR
}

// SolveRGE running RGE for Step
func (C *Container) SolveRGE(mt, xi float64) {
	R := Initialize(mt)
	C[0] = R.Copy()

	for i := range C {
		R.Run(mt, xi)
		C[i] = R.Copy()
	}
}
