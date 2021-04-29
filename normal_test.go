package dist

import (
	"fmt"
	"testing"
)

func Test_Normal(t *testing.T) {
	var numbers []float64
	d := NormalDistribution{
		DistributionType: "Normal",
	}

	for i := 1; i < 100; i++ {
		n, _ := d.RandVar(.5, 1)
		numbers = append(numbers, n[0])
		numbers = append(numbers, n[1])

		fmt.Println(n[0])
		fmt.Println(n[1])

	}
	// PlotHistogram(5, numbers)
}
