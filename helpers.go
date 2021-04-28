package dist

import (
	"math"
	"math/rand"
)

// The joy of polymorphism in Go!
type Distribution interface {
	ExpectedValue([]interface{}) interface{}
	Variance([]interface{}) float64
	RandVar(interface{}) ([]interface{}, error)
}

// To get a number between 0 and 1
func RandomFloat() float64 {
	for {
		x := rand.Float64()
		if x >= 0 && x <= 1 {
			return x
		}
	}

}

func ArrayMean(values []float64) float64 {
	total := float64(0)
	l := 1
	for _, number := range values {

		total = total + math.Ceil(number*100)/100
		l = l + 1
	}
	x := float64(total) / float64(l)
	return math.Ceil(x*100) / 100
}
