package dist

import (
	"fmt"
	"testing"
)

func Test_Exponential(t *testing.T) {
	var numbers []float64

	a := 1.0
	d := ExponentialDistribution{
		DistributionType: "Exponential",
	}

	for i := 1; i < 10000; i++ {
		n, _ := d.RandVar(a)
		numbers = append(numbers, n)

		// fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(a)
	v := d.Variance(a)

	fmt.Printf("%v Distribution Expected value for a %v = %v", d.DistributionType, a, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
	fmt.Println("")

	HistPlot(numbers)

}
