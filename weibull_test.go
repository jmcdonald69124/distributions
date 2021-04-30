package dist

import (
	"fmt"
	"testing"
)

func Test_Weibull(t *testing.T) {
	var numbers []float64
	a := 1.5
	b := float64(1)

	d := WeibullDistribution{
		DistributionType: "Weibull",
	}

	for i := 1; i < 100000; i++ {
		n, _ := d.RandVar(a, b)
		numbers = append(numbers, n)
		// fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(a, b)
	v := d.Variance(a, b)

	fmt.Printf("Expected value for a %v, b %v = %v", a, b, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
}
