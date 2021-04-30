package dist

import (
	"fmt"
	"testing"
)

func Test_Geometric(t *testing.T) {
	var numbers []float64

	p := .8

	d := GeometricDistribution{
		DistributionType: "Geometric",
	}

	for i := 1; i < 100; i++ {
		n, _ := d.RandVar(p)
		numbers = append(numbers, n)

		fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(p)
	v := d.Variance(p)

	fmt.Printf("%v Distribution Expected value for p %v = %v", d.DistributionType, p, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
	fmt.Println("")

	HistPlot(numbers)

}
