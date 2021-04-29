package dist

import "math"

type TriangularDistribution struct {
	DistributionType string
}

// Closed form inversion
func (d TriangularDistribution) RandVar(min float64, mode float64, max float64) (float64, error) {
	u1 := RandomFloat()

	if u1 < ((mode - min) / (max / min)) {
		return min + math.Sqrt((max-min)*(mode-min)*u1), nil
	}

	return max - math.Sqrt((max-min)*(max-mode)*(1-u1)), nil
}

// Expected Value : min + mode + max / 3
func (d TriangularDistribution) ExpectedValue(min float64, mode float64, max float64) float64 {
	x := (min + mode + max) / 3
	return x
}

// Variance : min^2 + max^2 + mode^2 - (min * max) - (min * mode) - (max * mode) / 18
func (d TriangularDistribution) Variance(min float64, mode float64, max float64) float64 {
	x := ((math.Pow(min, 2) + math.Pow(max, 2) + math.Pow(mode, 2)) - ((min * max) - (min * mode) - (max * mode))) / 18
	return x
}
