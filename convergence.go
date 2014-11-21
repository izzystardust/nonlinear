package nonlinear

import "math"

// ConvergenceRate returns the average q given a slice of residuals
// based on solving |e_{n+1}| = |e_{n}|^q for q
func ConvergenceRate(residuals []float64) float64 {
	abs := math.Abs
	log := math.Log
	var qs []float64
	for i := 0; i < len(residuals)-1; i++ {
		q := log(abs(residuals[i+1])) / log(abs(residuals[i]))
		if !math.IsNaN(q) && !math.IsInf(q, 0) {
			qs = append(qs, q)
		}
	}
	total := 0.0
	for _, q := range qs {
		total += q
	}
	return total / float64(len(qs))
}
