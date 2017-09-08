package RGE

// Container contains coupling constants
type Container [Step]RGE

// CosmoContainer contains cosmological variables
type CosmoContainer [Step]Cosmo

// Gamma for Convenience
type Gamma func(float64, float64) float64
