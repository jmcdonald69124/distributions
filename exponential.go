package dist

import (
	"errors"
	"math"
)

type ExponentialDistribution struct {
	DistributionType string
}

// Generate random numbers that fit the
// exponential distribution

// Parameters: a = Scale parameter and a > 0

//  Inverse transformation
func (d ExponentialDistribution) RandVar(a float64) (float64, error) {

	u1 := RandomFloat()

	if a <= 0 {
		return -1, errors.New("a must be greater than 0")
	}

	x := -a * math.Log(u1)
	return x, nil
}

//  Expected Value: a
func (d ExponentialDistribution) ExpectedValue(a float64) float64 {
	return 1/a
}

//  Variance: a^2
func (d ExponentialDistribution) Variance(a float64) float64 {
	return math.Pow(1/a, 2)
}
