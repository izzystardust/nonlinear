package nonlinear

import "math"

// Shamanskii solves f(x) = 0 using the Shamanskii method with derivative df, derivative update n, and
// initial iterate x to within tolerance. If maxits iterations pass, returns early.
func Shamanskii(f, df fn, n int, x, tolerance float64, maxits int) (float64, []float64, []float64) {
	xs := []float64{x}
	evals := []float64{f(x)}
	dfe := df(x)
	its := 0
	for math.Abs(f(x)) > tolerance && its < maxits {
		if n != 0 && its%n == 0 {
			dfe = df(x)
		}
		x = x - f(x)/dfe
		xs = append(xs, x)
		evals = append(evals, f(x))
		its++
	}
	return x, xs, evals
}
