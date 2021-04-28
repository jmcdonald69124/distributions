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

	for i := 1; i < 100; i++ {
		n, _ := d.RandVar(a)
		numbers = append(numbers, n)

		fmt.Println(n)
	}

}
