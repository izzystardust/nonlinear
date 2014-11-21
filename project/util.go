package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
)

var sin = math.Sin
var cos = math.Cos
var atan = math.Atan
var sq = func(x float64) float64 {
	return math.Pow(x, 2)
}

// writeResults writes the calculated points in csv format to dest
func writeResults(dest io.Writer, xs, evals []float64) error {
	if len(xs) != len(evals) {
		return fmt.Errorf("writeResults: mismatched lengths")
	}
	w := csv.NewWriter(dest)
	for i := range xs {
		x := fmt.Sprintf("%E", xs[i])
		e := fmt.Sprintf("%E", evals[i])
		if err := w.Write([]string{x, e}); err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}
