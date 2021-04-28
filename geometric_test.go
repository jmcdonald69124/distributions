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
	// Todo mean and std dev needs to be calculated

}
