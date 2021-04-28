package dist

import (
	"errors"
	"math"
)

type NormalDistribution struct {
	DistributionType string
}

// Module 7 Slide 4 of 16  / Box-Muller Method

// z0 = sqrt(−2 ln(U1) * cos(2 * PI * U2))
// z1 = sqrt(−2 ln(U1) * sin(2 * PI * U2))
// Paramters
// s Standard deviation σ > 0
// m Mean
func (d NormalDistribution) RandVar(m float64, s float64) ([]float64, error) {

	if s <= 0 {
		return []float64{-1}, errors.New("s must be greater than 0")
	}

	u1 := RandomFloat()
	u2 := RandomFloat()

	lu1 := float64(s) * math.Sqrt(-2*math.Log(u1))

	z0 := (lu1 * math.Cos((2*math.Pi)*u2)) + m
	z1 := (lu1 * math.Sin((2*math.Pi)*u2)) + m

	return []float64{z0, z1}, nil
}

// Expected Value : m
func (d NormalDistribution) ExpectedValue(m float64) float64 {

	return m
}

// Variance : σ^2
func (d NormalDistribution) Variance(s float64) float64 {

	return math.Pow(s, 2)
}
