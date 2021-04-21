package dist

import (
	"math"
)

// Generate random numbers that fit the
// weibull distribution
func Exp() float64 {
	u1 := RandomFloat()

	// Lambda === 1.5
	// Inverse Transform
	x := math.Log(1-u1) / 1.5

	return x
}
