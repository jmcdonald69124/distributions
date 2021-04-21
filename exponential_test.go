package dist

import (
	"fmt"
	"testing"
)

func Test_Exponential(t *testing.T) {
	var numbers []float64

	for i := 1; i < 100; i++ {
		n := Exp()
		numbers = append(numbers, n)

		fmt.Println(n)
	}
	// Todo mean and std dev needs to be calculated

}
