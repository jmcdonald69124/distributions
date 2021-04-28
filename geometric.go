package dist

import (
	"errors"
	"math"
)

// Generate random numbers that fit the
// geometric distribution

// Parameters: p = Probability of success

type GeometricDistribution struct {
	DistributionType string
}

// Inverse transformation
func (d GeometricDistribution) RandVar(p float64) (float64, error) {

	if p < 0 || p > 1 {
		return -1, errors.New("p must be a value between 0 and 1")
	}

	u1 := RandomFloat()

	x := math.Log(u1) / math.Log(1-p)
	return x, nil
}

//  Expected Value: 1 / p
func (d GeometricDistribution) ExpectedValue(p float64) float64 {

	return 1 / p
}

//  Variance: 1 - p / p^2
func (d GeometricDistribution) Variance(p float64) float64 {

	return 1 - p/math.Pow(p, 2)
}
