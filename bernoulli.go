package dist

import (
	"errors"
)

type BernoulliDistribution struct {
	DistributionType string
}

// Generate random numbers that fit the
// Bernoulli distribution given a probability

func (d BernoulliDistribution) RandVar(p float64) (int, error) {

	if p < 0 || p > 1 {
		return -1, errors.New("p must be a value between 0 and 1")
	}

	u1 := RandomFloat()

	if u1 < p {
		return 1, nil
	}
	return 0, nil
}

// Expected Value : p
func (d BernoulliDistribution) ExpectedValue(p float64) float64 {

	return p
}

// Variance: p(1 − p)
func (d BernoulliDistribution) Variance(p float64) float64 {

	return p * (1 - p)
}
