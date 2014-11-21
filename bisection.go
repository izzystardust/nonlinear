package nonlinear

// Bisect solves f(x) = 0 on the interval (a,b) using the bisection method to within tolerance.
// If more than maxits iterations pass, return early
func Bisect(f fn, a, b, tolerance float64, maxits int) (float64, []float64, []float64) {
	ms := []float64{}
	fms := []float64{}
	i := 0
	if a > b {
		a, b = b, a
	}
	if a == b {
		panic("a can't equal b!")
	}
	for b-a > tolerance && i < maxits {
		m := (b + a) / 2
		ms = append(ms, m)
		fms = append(fms, f(m))
		if f(a)*f(m) <= 0 {
			b = m
		} else {
			a = m
		}
		i++
	}

	return (b + a) / 2, ms, fms
}
