package dist

import (
	"math"
	"math/rand"
)

// Gamma Distribution

type GammaDistribution struct {
	DistributionType string
}

//
func (d GammaDistribution) RandVar(k int, s int) (float64, error) {

	produ1 := float64(0)

	for i := 0; i < k; i++ {
		u1 := rand.NormFloat64()

		produ1 = float64(produ1) + u1
	}
	x := float64(-1/s) * math.Log(produ1)
	return x, nil
}

// Expected Value : λ
func (d GammaDistribution) ExpectedValue(k int, s int) int {
	return k * s
}

// Variance : λ
func (d GammaDistribution) Variance(k int, s int) float64 {
	return float64(k) * math.Pow(float64(s), 2)
}
