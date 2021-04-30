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

	for i := 1; i < 100000; i++ {
		n, _ := d.RandVar(min, mode, max)
		numbers = append(numbers, n)
		//fmt.Println(n)
	}

	m := ArrayMean(numbers)
	ev := d.ExpectedValue(min, mode, max)
	v := d.Variance(min, mode, max)

	fmt.Printf("Expected value for min %v, mode %v, max %v = %v", min, mode, max, ev)
	fmt.Println("")
	fmt.Printf("%v Distribution Variance %v", d.DistributionType, v)
	fmt.Println("")
	fmt.Printf("Actual mean of array values %v", m)
}
