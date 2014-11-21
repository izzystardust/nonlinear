package nonlinear

import "math"

// Secant solves f(x) = 0 using the secant method, with initial iterates x0 and x1 to
// within tolerance. If maxits iterations are exceeded, returns early.
func Secant(f fn, x0, x1, tolerance float64, maxits int) (float64, []float64, []float64) {
	xOlder := x0
	xOld := x1
	its := 0
	xs := []float64{x0, x1}
	evals := []float64{f(x0), f(x1)}
	for math.Abs(f(xOld)) > tolerance && its < maxits {
		x := xOld - f(xOld)*((xOld-xOlder)/(f(xOld)-f(xOlder)))
		xOlder = xOld
		xOld = x
		xs = append(xs, x)
		evals = append(evals, f(x))
		its++
	}
	return xOld, xs, evals
}
