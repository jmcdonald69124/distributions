package dist

import (
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

func HistPlot(values plotter.Values) {
	p := plot.New()

	p.Title.Text = "histogram plot"

	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}
