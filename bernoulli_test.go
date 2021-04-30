package dist

import (
	"fmt"
	"testing"
)

// Generate random numbers that fit the
// Bernoulli distribution

func Test_Bernoulli(t *testing.T) {
	var numbers []float64
	var p = .25
	d := BernoulliDistribution{
		DistributionType: "Bernoulli",
	}

	for i := 1; i < 100000; i++ {
		n, _ := d.RandVar(p)
		numbers = append(numbers, float64(n))

		// fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(p)
	v := d.Variance(p)

	fmt.Printf("Expected value for p %v = %v", p, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)

	HistPlot(numbers)
}
