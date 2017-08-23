package rge

// Container contains coupling constamts
type Container [Step]float64

// Bay is array of Container
type Bay []Container

// Gamma for Convenience
type Gamma func(float64, float64) float64
