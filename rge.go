package RGE

import (
	"fmt"
	"math"

	"github.com/Axect/csv"
)

// Run just run
type Run interface {
	Run(float64, float64) // Run RGE
	Copy()                // Copy RGE to RGE
	InputFormula(RGE, float64, float64)
}

// Run is Single Running - You can change Numerical Integration method here (Default: Euler)
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

// SolveRGE running RGE for Step
func (C *Container) SolveRGE(mt, xi float64) {
	R := Initialize(mt)
	C[0] = R.Copy()

	for i := range C {
		R.Run(mt, xi)
		C[i] = R.Copy()
	}
}

// Running is main tool
func Running(mt, xi float64) []int {
	var C Container
	C.SolveRGE(mt, xi)

	W := make([][]string, len(C), len(C))

	for i, elem := range C {
		W[i] = Convert([]float64{elem.t, elem.lH, elem.yt, elem.g1, elem.g2, elem.g3, elem.G})
	}
	mtint := int(mt)
	mtfloat := int((mt-float64(mtint))*100 + 0.49)
	xiint := int(xi)
	title := fmt.Sprintf("Data/Gauge_%d_%d_%d.csv", mtint, mtfloat, xiint)
	csv.Write(W, title)
	return []int{mtint, mtfloat, xiint}
}
