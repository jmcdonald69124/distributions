package dist

import (
	"math"
)

// Generate random numbers that fit the
// weibull distribution
func Weibull() float64 {
	u1 := RandomFloat()

	// Lambda === 1.5
	// Inverse Transform
	x := 1.5 * math.Pow(-math.Log(1-u1), (1/1.5))

	return x
}
