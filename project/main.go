package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/millere/nonlinear"
)

func main() {
	maxIterations := 10000
	useBisection(maxIterations)
	useChord(maxIterations)
	useNewton(maxIterations)
	useSecant(maxIterations)
	useShamanskii(maxIterations, 3)
}

// fns is a slice containing each function, its derivative, and the given x0.
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

// because some methods require multiple initial iterates, x2s holds those.
// x2s[i] is x1 for the ith equation in fns
var x2s = []float64{3.5, 0.0, -1.0, -1.0, -1.0, 3.0}

func useBisection(maxIterations int) {
	fmt.Println("Using Bisection Method")
	for i, c := range fns {
		x, xs, evals := nonlinear.Bisect(c.f, c.x0, x2s[i], math.Pow(10, -12), maxIterations)
		convergenceRate := nonlinear.ConvergenceRate(evals)
		fmt.Printf("%c: %E (converging at %f)\n", i+'A', x, convergenceRate)
		f, err := os.Create(fmt.Sprintf("bisection-%c.csv", i+'A'))
		if err != nil {
			log.Println("Unable to create", i+'A', err)
		}
		if err := writeResults(f, xs, evals); err != nil {
			log.Println("Unable to write", i+'A', err)
		}
	}
}

func useChord(maxIterations int) {
	fmt.Println("Using Chord Method")
	for i, c := range fns {
		x, xs, evals := nonlinear.Chord(c.f, c.df, c.x0, math.Pow(10, -12), maxIterations)
		convergenceRate := nonlinear.ConvergenceRate(evals)
		fmt.Printf("%c: %E (converging at %f)\n", i+'A', x, convergenceRate)
		f, err := os.Create(fmt.Sprintf("chord-%c.csv", i+'A'))
		if err != nil {
			log.Println("Unable to create", i+'A', err)
		}
		if err := writeResults(f, xs, evals); err != nil {
			log.Println("Unable to write", i+'A', err)
		}
	}
}

func useNewton(maxIterations int) {
	fmt.Println("Using Newton's Method")
	for i, c := range fns {
		x, xs, evals := nonlinear.Newton(c.f, c.df, c.x0, math.Pow(10, -12), maxIterations)
		convergenceRate := nonlinear.ConvergenceRate(evals)
		fmt.Printf("%c: %E (converging at %f)\n", i+'A', x, convergenceRate)
		f, err := os.Create(fmt.Sprintf("newton-%c.csv", i+'A'))
		if err != nil {
			log.Println("Unable to create", i+'A', err)
		}
		if err := writeResults(f, xs, evals); err != nil {
			log.Println("Unable to write", i+'A', err)
		}
	}
}

func useSecant(maxIterations int) {
	fmt.Println("Using Secant Method")
	for i, c := range fns {
		x, xs, evals := nonlinear.Secant(c.f, c.x0, x2s[i], math.Pow(10, -12), maxIterations)
		convergenceRate := nonlinear.ConvergenceRate(evals)
		fmt.Printf("%c: %E (converging at %f)\n", i+'A', x, convergenceRate)
		f, err := os.Create(fmt.Sprintf("secant-%c.csv", i+'A'))
		if err != nil {
			log.Println("Unable to create", i+'A', err)
		}
		if err := writeResults(f, xs, evals); err != nil {
			log.Println("Unable to write", i+'A', err)
		}
	}
}

func useShamanskii(maxIterations, n int) {
	fmt.Println("Using Shamanskii's Method (n = 3)")
	for i, c := range fns {
		x, xs, evals := nonlinear.Shamanskii(c.f, c.df, n, c.x0, math.Pow(10, -12), maxIterations)
		convergenceRate := nonlinear.ConvergenceRate(evals)
		fmt.Printf("%c: %E (converging at %f)\n", i+'A', x, convergenceRate)
		f, err := os.Create(fmt.Sprintf("shamanskii-%c.csv", i+'A'))
		if err != nil {
			log.Println("Unable to create", i+'A', err)
		}
		if err := writeResults(f, xs, evals); err != nil {
			log.Println("Unable to write", i+'A', err)
		}
	}
}
