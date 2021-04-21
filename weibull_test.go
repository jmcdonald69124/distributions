package dist

import (
	"fmt"
	"testing"
)

func Test_Weibull(t *testing.T) {
	var numbers []float64

	for i := 1; i < 100; i++ {
		n := Weibull()
		numbers = append(numbers, n)

		fmt.Println(n)
	}
	// Todo mean and std dev needs to be calculated

}
