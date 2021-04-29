package dist

import (
	"math"
)

// Erlang Distribution

type ErlangDistribution struct {
	DistributionType string
}

//
func (d ErlangDistribution) RandVar(k int, λ float64) (float64, error) {

	produ1 := float64(0)

	for i := 0; i < k; i++ {
		u1 := RandomFloat()

		produ1 = float64(produ1) + u1
	}
	x := -math.Log(produ1) / λ
	return x, nil
}

// Expected Value : λ
func (d ErlangDistribution) ExpectedValue(k int, λ float64) float64 {
	return float64(k) / λ
}

// Variance : λ
func (d ErlangDistribution) Variance(k int, λ float64) float64 {
	return float64(k) / math.Pow(λ, 2)
}
