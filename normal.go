package dist

import (
	"math"
	"math/rand"
)

// Module 7 Slide 4 of 16  / Box-Muller Method
// z = sqrt(−2 ln(U1) cos(2 * PI * U2))

func calcNorm() float64 {
	u1 := 0 + rand.Float64()*(1-0)
	u2 := 0 + rand.Float64()*(1-0)
	lu1 := math.Log(u1)
	cu2 := math.Cos(2 * math.Pi * u2)
	z := math.Sqrt(-2 * lu1 * cu2)

	return z
}

func Normal() float64 {

	var z float64
	for {
		z = calcNorm()
		if !math.IsNaN(z) {
			break
		}
	}

	return z
}
