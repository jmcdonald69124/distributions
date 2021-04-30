package dist

import (
	"fmt"
	"testing"
)

func Test_Erlang(t *testing.T) {
	var numbers []float64
	lambda := float64(2)
	k := 1

	d := ErlangDistribution{
		DistributionType: "Erlang",
	}

	for i := 1; i < 100000; i++ {
		n, _ := d.RandVar(k, lambda)
		numbers = append(numbers, n)
		// fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(k, lambda)
	v := d.Variance(k, lambda)

	fmt.Printf("%v Distribution Expected value for lambda %v, k %v = %v", d.DistributionType, lambda, k, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
	fmt.Println("")

	HistPlot(numbers)
}
