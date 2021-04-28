package dist

import (
	"errors"
	"math"
)

type WeibullDistribution struct {
	DistributionType string
}

// Generate random numbers that fit the
// weibull distribution

// Paramaters : a = Scale parameter a > 0
//				b = Shape parameter b > 0
func (d WeibullDistribution) RandVar(a float64, b float64) (float64, error) {
	u1 := RandomFloat()

	if a <= 0 {
		return -1, errors.New("a must be greater than 0")
	}

	if b <= 0 {
		return -1, errors.New("b must be greater than 0")
	}

	// b < 1 ⇒ failure rate increasing with time
	// b > 1 ⇒ failure rate decreases with time
	// b = 1 ⇒ failure rate is constant

	// Inverse Transform
	x := a * math.Pow(math.Log(u1), (1/b))
	return x, nil
}

// Expected Value : a / b Γ (1/b)
func (d WeibullDistribution) ExpectedValue(a float64, b float64) float64 {
	x := (a / b) * math.Gamma(1/b)
	return math.Ceil(x*100) / 100
}

// Variance : a^2 / b^2 [2 * b * Γ(2/b) - {Γ(1/b)}^2]
func (d WeibullDistribution) Variance(a float64, b float64) float64 {
	x := (math.Pow(a, 2) / math.Pow(b, 2)) * ((2 * b * math.Gamma(2/b)) - math.Pow(math.Gamma(1/b), 2))
	return math.Ceil(x*100) / 100
}
