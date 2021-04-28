package dist

import (
	"fmt"
	"testing"
)

func Test_Triangular(t *testing.T) {
	var numbers []float64
	min := float64(0)
	mode := .5
	max := float64(1)

	d := TriangularDistribution{
		DistributionType: "Triangular",
	}

	for i := 1; i < 100; i++ {
		n, _ := d.RandVar(min, mode, max)
		numbers = append(numbers, n)
		// fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(min, mode, max)

	fmt.Printf("Expected value for a %v, b %v = %v", a, b, ev)
	fmt.Println("--")
	fmt.Printf("Actual mean of array values %v", m)
}
