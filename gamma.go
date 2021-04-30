package dist

import (
	"math"
)

// Gamma Distribution

type GammaDistribution struct {
	DistributionType string
}

/*
PsudoCode from https://www.stat.purdue.edu/~xbw/research/jss102013.gamma.pdf

θ=ln(α)andc= α
Compute bs(θ) and B.max = exp{bs(θ)} Compute bw(θ) and B.min = −exp{bw(θ)} k=1
while k ≤ n do
u = Uniform (0, 1)
v = Uniform (0, 1) × (B.max − B.min) + B.min t=v/uandt1 =exp(t/c+θ)
if 2ln(u) ≤ α+c×t−t1 then
Deliver t1 if α ≥ 0.01; Otherwise deliver t/c + θ.
k=k+1 end if
end while

*/

// Based on the algorithm found here https://www.hongliangjie.com/2012/12/19/how-to-generate-gamma-random-variables/
func (d GammaDistribution) RandVar(alpha int, lambda int) (float64, error) {

	norm_m := 0.5
	norm_s := 1.0
	v := 0.0
	x := 0.0
	norm := NormalDistribution{
		DistributionType: "Normal",
	}

	if alpha < 1 {
		gamma := GammaDistribution{
			DistributionType: "Gamma",
		}
		x, _ = gamma.RandVar(alpha+1, lambda)
		unif := RandomFloat()
		return x * math.Pow(unif, (float64(alpha)-float64(1))), nil
	}

	if alpha > 1 {

		d := alpha - 1/3
		c := float64(1) / math.Sqrt(float64(9)*float64(d))
		// f := 1
		flag := true
		for flag {
			z, _ := norm.RandVar(norm_m, norm_s)
			if z[0] > -1/c {
				v = math.Pow((1 + c*z[0]), 3)
				u := RandomFloat()
				flag = math.Log(u) > (math.Pow((0.5*float64(z[0])), 2) + float64(d) - float64(d)*v + float64(d)*math.Log(v))
			}
		}
		x = float64(d) * v / float64(lambda)
	}
	return x, nil
}

// Expected Value : λ
func (d GammaDistribution) ExpectedValue(k int, s int) int {
	return k / s
}

// Variance : λ
func (d GammaDistribution) Variance(k int, s int) float64 {
	return float64(k) / math.Pow(float64(s), 2)
}
