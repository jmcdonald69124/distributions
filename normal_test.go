package dist

import (
	"fmt"
	"testing"

	stats "github.com/montanaflynn/stats"
)

func Test_Normal(t *testing.T) {
	var numbers []float64

	for i := 1; i < 10; i++ {
		n := Normal()
		numbers = append(numbers, n)
		fmt.Println(n)
	}

	v := stats.NormFit(numbers)
	fmt.Printf("NormFit : %v", v)
}
