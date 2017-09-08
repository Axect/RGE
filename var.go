package RGE

// =============================================================================
//  Declare your rge variables
// =============================================================================

// RGE contains Gauge constants & field variables
type RGE struct {
	t   float64
	lH  float64
	yt  float64
	g1  float64
	g2  float64
	g3  float64
	phi float64
	G   float64
}

// Cosmo contains potential & cosmological parameters
type Cosmo struct {
	V   float64
	eps float64
	eta float64
	A   float64
}
