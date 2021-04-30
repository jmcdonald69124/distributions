package dist

import (
	"fmt"
	"testing"
)

func Test_NegativeBinomialDistribution(t *testing.T) {
	var numbers []float64
	p := .25
	n := 4
	d := NegativeBinomialDistribution{
		DistributionType: "NegativeBinomial",
	}

	for i := 1; i < 10000000; i++ {
		n, _ := d.RandVar(p, n)
		numbers = append(numbers, n)

		//fmt.Println(n)

	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(p, n)
	v := d.Variance(p, n)

	fmt.Printf("Expected value for p %v, n %v = %v", p, n, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
	fmt.Println("")
}
