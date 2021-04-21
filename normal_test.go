package dist

import (
	"fmt"
	"testing"
)

func Test_Normal(t *testing.T) {
	var numbers []float64

	for i := 1; i < 100; i++ {
		n, o := Normal()
		numbers = append(numbers, n)
		numbers = append(numbers, o)

		fmt.Println(n)
		fmt.Println(o)
	}
	// Todo mean and std dev needs to be calculated

}
