package dist

import (
	"errors"
	"math"
)

type PoissonDistribution struct {
	DistributionType string
}

// Direct method - using queues from
// https://hpaulkeeler.com/simulating-poisson-random-variables-direct-method/
func (d PoissonDistribution) RandVar(λ float64) (float64, error) {

	if λ < 1 || λ > 20 {
		return -1, errors.New("λ must be greater than 1 and less than 20")
	}

	u1 := float64(1)
	produ1 := float64(1)
	eλ := math.Exp(-λ)
	randPoisson := -1

	for randPoisson = -1; float64(produ1) > eλ; randPoisson++ {
		u1 = RandomFloat()
		produ1 = float64(produ1) * u1
	}

	return float64(randPoisson), nil
}

// Expected Value : λ
func (d PoissonDistribution) ExpectedValue(λ float64) float64 {
	return λ
}

// Variance : λ
func (d PoissonDistribution) Variance(λ float64) float64 {
	return λ
}
