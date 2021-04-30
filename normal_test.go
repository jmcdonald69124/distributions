package dist

import (
	"fmt"
	"testing"
)

func Test_Normal(t *testing.T) {
	var numbers []float64
	m := 0.5
	s := 1.0
	d := NormalDistribution{
		DistributionType: "Normal",
	}

	for i := 1; i < 100000; i++ {
		n, _ := d.RandVar(m, s)
		numbers = append(numbers, n[0])
		numbers = append(numbers, n[1])

		// fmt.Println(n[0])
		// fmt.Println(n[1])

	}

	n := ArrayMean(numbers)
	ev := d.ExpectedValue(m)
	v := d.Variance(s)

	fmt.Printf("%v Distribution Expected value for m %v, s %v = %v", d.DistributionType, m, s, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", n)
	fmt.Println("")

}
