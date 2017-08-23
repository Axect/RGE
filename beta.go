package rge

import "math"

//-------------------------
// Declare Beta Functions
//-------------------------

// Beta Function: You can declare all of beta function list here
type Beta struct {
	// 1-loop order
	B1lH   float64
	B1yt   float64
	B1g1   float64
	B1g2   float64
	B1g3   float64
	gamma1 float64

	// 2-loop order
	B2lH   float64
	B2yt   float64
	B2g1   float64
	B2g2   float64
	B2g3   float64
	gamma2 float64

	// Total
	BlH   float64
	Byt   float64
	Bg1   float64
	Bg2   float64
	Bg3   float64
	gamma float64
}

// InputFormula to Beta functions
func (B *Beta) InputFormula(R RGE, mt, xi float64) {
	// Necessary Variables
	hg := math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	sh := (1. + xi*math.Pow(hg, 2)/math.Pow(MpR, 2)) / (1. + (1.+6.*xi)*xi*math.Pow(hg, 2)/math.Pow(MpR, 2))

	// 1-loop Beta Function
	B.B1lH = 6.*(1.+3.*math.Pow(sh, 2))*math.Pow(R.lH, 2) + 12.*R.lH*math.Pow(R.yt, 2) - 6.*math.Pow(R.yt, 4) - 3.*R.lH*(3.*math.Pow(R.g2, 2)+math.Pow(R.g1, 2)) + 3./8*(2.*math.Pow(R.g2, 4)+math.Pow((math.Pow(R.g1, 2)+math.Pow(R.g2, 2)), 2))
	B.B1yt = R.yt * ((23./6+2./3*sh)*math.Pow(R.yt, 2) - (8.*math.Pow(R.g3, 2) + 9./4*math.Pow(R.g2, 2) + 17./12*math.Pow(R.g1, 2)))
	B.B1g1 = (81. + sh) / 12 * math.Pow(R.g1, 3)
	B.B1g2 = (sh - 39.) / 12 * math.Pow(R.g2, 3)
	B.B1g3 = -7. * math.Pow(R.g3, 3)
	B.gamma1 = -1. / (16. * math.Pow(math.Pi, 2)) * (9./4*math.Pow(R.g2, 2) + 3./4*math.Pow(R.g1, 2) - 3.*math.Pow(R.yt, 2))

	// 2-loop Beta function
	B.B2lH = 1./48*((912.+3.*sh)*math.Pow(R.g2, 6)-(290.-sh)*math.Pow(R.g1, 2)*math.Pow(R.g2, 4)-(560.-sh)*math.Pow(R.g1, 4)*math.Pow(R.g2, 2)-(380.-sh)*math.Pow(R.g1, 6)) + (38.-8*sh)*math.Pow(R.yt, 6) - math.Pow(R.yt, 4)*(8./3*math.Pow(R.g1, 2)+32.*math.Pow(R.g3, 2)+(12.-117.*sh+108.*math.Pow(sh, 2))*R.lH) + R.lH*(-1./8*(181.+54.*sh-162.*math.Pow(sh, 2))*math.Pow(R.g2, 4)+1./4*(3.-18.*sh+54.*math.Pow(sh, 2))*math.Pow(R.g1, 2)*math.Pow(R.g2, 2)+1./24*(90.+377.*sh+162.*math.Pow(sh, 2))*math.Pow(R.g1, 4)+(27.+54.*sh+27.*math.Pow(sh, 2))*math.Pow(R.g2, 2)*R.lH+(9.+18.*sh+9*math.Pow(sh, 2))*math.Pow(R.g1, 2)*R.lH-(48.+288.*sh-324.*math.Pow(sh, 2)+624.*math.Pow(sh, 3)-324.*math.Pow(sh, 4))*math.Pow(R.lH, 2)) + math.Pow(R.yt, 2)*(-9./4*math.Pow(R.g2, 4)+21./2*math.Pow(R.g1, 2)*math.Pow(R.g2, 2)-19./4*math.Pow(R.g1, 4)+R.lH*(45./2*math.Pow(R.g2, 2)+85./6*math.Pow(R.g1, 2)+80.*math.Pow(R.g3, 2)-(36.+108.*math.Pow(sh, 2))*R.lH))
	B.B2yt = R.yt * (-23./4*math.Pow(R.g2, 4) - 3./4*math.Pow(R.g1, 2)*math.Pow(R.g2, 2) + 1187./216*math.Pow(R.g1, 4) + 9.*math.Pow(R.g2, 2)*math.Pow(R.g3, 2) + 19./9*math.Pow(R.g1, 2)*math.Pow(R.g3, 2) - 108.*math.Pow(R.g3, 4) + (225./16*math.Pow(R.g2, 2)+131./16*math.Pow(R.g1, 2)+36.*math.Pow(R.g3, 2))*sh*math.Pow(R.yt, 2) + 6.*(-2.*math.Pow(sh, 2)*math.Pow(R.yt, 4)-2.*math.Pow(sh, 3)*math.Pow(R.yt, 2)*R.lH+math.Pow(sh, 2)*math.Pow(R.lH, 2)))
	B.B2g1 = 199./18*math.Pow(R.g1, 5) + 9./2*math.Pow(R.g1, 3)*math.Pow(R.g2, 2) + 44./3*math.Pow(R.g1, 3)*math.Pow(R.g3, 2) - 17./6*sh*math.Pow(R.g1, 3)*math.Pow(R.yt, 2)
	B.B2g2 = 3./2*math.Pow(R.g1, 2)*math.Pow(R.g2, 3) + 35./6*math.Pow(R.g2, 5) + 12.*math.Pow(R.g2, 3)*math.Pow(R.g3, 2) - 3./2*sh*math.Pow(R.g2, 3)*math.Pow(R.yt, 2)
	B.B2g3 = 11./6*math.Pow(R.g1, 2)*math.Pow(R.g3, 3) + 9./2*math.Pow(R.g2, 2)*math.Pow(R.g3, 3) - 26.*math.Pow(R.g3, 5) - 2.*sh*math.Pow(R.g3, 3)*math.Pow(R.yt, 2)
	B.gamma2 = -1. / (math.Pow(16*math.Pow(math.Pi, 2), 2)) * (271./32*math.Pow(R.g2, 4) - 9./16*math.Pow(R.g1, 2)*math.Pow(R.g2, 2) - 431./96*sh*math.Pow(R.g1, 4) - 5./2*(9./4*math.Pow(R.g2, 2)+17./12*math.Pow(R.g1, 2)+8*math.Pow(R.g3, 2))*math.Pow(R.yt, 2) + 27./4*sh*math.Pow(R.yt, 4) - 6*math.Pow(sh, 3)*math.Pow(R.lH, 2))
}

// Calculate Total Beta Function
func (B *Beta) BetaFunction() {
	B.gamma = 1./(16*math.Pow(math.Pi, 2))*B.gamma1 + 1./math.Pow(16*math.Pow(math.Pi, 2), 2)*B.gamma2
	g := MakeBeta(B.gamma)

	B.BlH = g(B.B1lH, B.B2lH)
	B.Byt = g(B.B1yt, B.B2yt)
	B.Bg1 = g(B.B1g1, B.B2g1)
	B.Bg2 = g(B.B1g2, B.B2g2)
	B.Bg3 = g(B.B1g3, B.B2g3)
}

// MakeBeta : Input gamma -> Output Beta function
func MakeBeta(g float64) Gamma {
	return func(f1, f2 float64) float64 {
		temp := 1./(16*math.Pow(math.Pi, 2))*f1 + 1./math.Pow((16*math.Pow(math.Pi, 2)), 2)*f2
		return temp / (1 + g)
	}
}
