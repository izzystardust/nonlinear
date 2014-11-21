package nonlinear

import (
	"math"
	"testing"
)

var sin = math.Sin
var cos = math.Cos
var atan = math.Atan
var sq = func(x float64) float64 {
	return math.Pow(x, 2)
}

func TestShamanskii(t *testing.T) {
	var fns = []struct {
		f, df func(float64) float64
		x0    float64
	}{
		{func(x float64) float64 { return sin(x) }, func(x float64) float64 { return cos(x) }, 2.9},
		{func(x float64) float64 { return x - cos(x) }, func(x float64) float64 { return 1 + sin(x) }, 1},
		{func(x float64) float64 { return atan(x) }, func(x float64) float64 { return 1 / (sq(x) + 1) }, 1},
		{func(x float64) float64 { return atan(x) }, func(x float64) float64 { return 1 / (sq(x) + 1) }, 12},
		{func(x float64) float64 { return sq(x) + 3 }, func(x float64) float64 { return 2 * x }, 10},
		{func(x float64) float64 { return math.Pow(x-2, 3) }, func(x float64) float64 { return 3 * sq(x-2) }, 1.5},
	}
	for _, c := range fns {
		chord, _, _ := Shamanskii(c.f, c.df, 0, c.x0, 10e-12, 1000)
		expect, _, _ := Chord(c.f, c.df, c.x0, 10e-12, 1000)
		if chord != expect && !(math.IsNaN(chord) && math.IsNaN(expect)) {
			t.Errorf("n=0 (Chord): Got %f, expected %f", chord, expect)
		}
	}

	for _, c := range fns {
		newton, _, _ := Shamanskii(c.f, c.df, 1, c.x0, 10e-12, 1000)
		expect, _, _ := Newton(c.f, c.df, c.x0, 10e-12, 1000)
		if newton != expect && !(math.IsNaN(newton) && math.IsNaN(expect)) {
			t.Errorf("n=1 (Newton): Got %f, expected %f", newton, expect)
		}
	}

}
