package dist

import (
	"errors"
	"math"
)

type NegativeBinomialDistribution struct {
	DistributionType string
}

func (d NegativeBinomialDistribution) RandVar(p float64, n int) (float64, error) {

	s := 0.0

	if p < 0 || p > 1 {
		return -1, errors.New("p must be a value between 0 and 1")
	}

	if n <= 0 {
		return -1, errors.New("n must be a positive integer")
	}

	for i := 0; i < n; i++ {
		u1 := RandomFloat()
		x := math.Log(u1) / math.Log(1-p)
		s = s + x
	}

	return float64(s), nil
}

// Expected Value : n(1-p)/p
func (d NegativeBinomialDistribution) ExpectedValue(p float64, n int) float64 {
	x := (float64(n) * (1.0 - p)) / p
	return x
}

// Variance : n(1-p)/p^2
func (d NegativeBinomialDistribution) Variance(p float64, n int) float64 {
	x := float64(n) * (1.0 - p) / math.Pow(p, 2)
	return x
}
