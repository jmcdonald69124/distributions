package dist

import (
	"fmt"
	"testing"
)

// Generate random numbers that fit the
// Bernoulli distribution

func Test_Bernoulli(t *testing.T) {
	var numbers []int
	var p = .25
	d := BernoulliDistribution{
		DistributionType: "Bernoulli",
	}

	for i := 1; i < 100; i++ {
		n, _ := d.RandVar(p)
		numbers = append(numbers, n)

		fmt.Println(n)
	}
	// Todo mean and std dev needs to be calculated

	m := d.ExpectedValue(p)

	fmt.Println(m)
}
