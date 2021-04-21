package dist

import (
	"math"
	"math/rand"
)

// To get a number between 0 and 1
func RandomFloat() float64 {
	for {
		x := rand.Float64()
		if x >= 0 && x <= 1 {
			return x
		}
	}

}

// Module 7 Slide 4 of 16  / Box-Muller Method

// z0 = sqrt(−2 ln(U1) * cos(2 * PI * U2))
// z1 = sqrt(−2 ln(U1) * sin(2 * PI * U2))

// standard normal distribution (mean = 0, stddev = 1)
func calcNorm() (float64, float64) {
	u1 := RandomFloat()
	u2 := RandomFloat()
	mean := 0
	stdev := 1

	lu1 := float64(stdev) * math.Sqrt(-2*math.Log(u1))

	z0 := (lu1 * math.Cos((2*math.Pi)*u2)) + float64(mean)
	z1 := (lu1 * math.Sin((2*math.Pi)*u2)) + float64(mean)
	return z0, z1
}

// standard normal distribution (mean = 0, stddev = 1)
func Mean(norm []float64) float64 {
	total := float64(0)
	for _, number := range norm {
		total = total + number
	}

	return total / float64(len(norm))

}

func NormStdDev([]float64) int32 {

	return 1
}

func Normal() (float64, float64) {
	var z0 float64
	var z1 float64
	for {
		z0, z1 = calcNorm()
		if !math.IsNaN(z0) || !math.IsNaN(z1) {
			break
		}
	}
	return z0, z1
}
