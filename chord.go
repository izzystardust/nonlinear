package nonlinear

import "math"

// Chord solves f(x) = 0 using the Chord Method variation of Newton's Method using an
// initial iterate x to within tolerance. If maxits iterations pass, returns early.
func Chord(f, df fn, x, tolerance float64, maxits int) (float64, []float64, []float64) {
	xs := []float64{x}
	evals := []float64{f(x)}
	df0 := df(x)
	its := 0
	for math.Abs(f(x)) > tolerance && its < maxits {
		x = x - f(x)/df0
		xs = append(xs, x)
		evals = append(evals, f(x))
		its++
	}
	return x, xs, evals
}
