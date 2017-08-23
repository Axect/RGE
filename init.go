package rge

import "math"

// Initialize RGE variables by some constants
func Initialize(mt float64) RGE {
	var R RGE
	R.t = 0
	R.lH = 0.12604 + 0.00206*(MH-125.15) - 0.00004*(mt-173.34)
	R.yt = (0.93690 + 0.00556*(mt-173.34) - 0.00003*(MH-125.15) - 0.00042*(alphasMZ-0.1184)/0.0007)
	R.g1 = 0.35830 + 0.00011*(mt-173.34) - 0.00020*(MW-80.384)/0.014
	R.g2 = 0.64779 + 0.00004*(mt-173.34) + 0.00011*(MW-80.384)/0.014
	R.g3 = 1.1666 + 0.00314*(alphasMZ-0.1184)/0.007 - 0.00046*(mt-173.34)
	R.phi = math.Sqrt(2.) / R.yt * mt * math.Exp(R.t)
	R.G = 1

	return R
}
