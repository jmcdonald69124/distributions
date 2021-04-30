package dist

import (
	"fmt"
	"testing"
)

func Test_Poisson(t *testing.T) {
	var numbers []float64
	lambda := float64(2)

	d := PoissonDistribution{
		DistributionType: "Poisson",
	}

	for i := 1; i < 100000; i++ {
		n, _ := d.RandVar(lambda)
		numbers = append(numbers, n)
		// fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(lambda)
	v := d.Variance(lambda)

	fmt.Printf("%v Distribution Expected value for lambda %v, = %v", d.DistributionType, lambda, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
	fmt.Println("")

	HistPlot(numbers)
}
