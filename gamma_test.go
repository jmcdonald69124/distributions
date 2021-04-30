package dist

import (
	"fmt"
	"testing"
)

func Test_Gamma(t *testing.T) {
	var numbers []float64
	k := 5
	s := 1

	d := GammaDistribution{
		DistributionType: "Gamma",
	}

	for i := 1; i < 100; i++ {
		n, _ := d.RandVar(k, s)
		numbers = append(numbers, n)
		fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(k, s)
	v := d.Variance(k, s)

	fmt.Printf("%v Distribution Expected value for s %v, k %v = %v", d.DistributionType, s, k, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
	fmt.Println("")
}
