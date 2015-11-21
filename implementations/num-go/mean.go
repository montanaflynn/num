package main

import (
	"errors"
	"math"
)

// Mean gets the average of a slice of numbers
func Mean(input []float64) float64 {
	return Sum(input) / float64(len(input))
}

// GeometricMean gets the geometric mean for a slice of numbers
func GeometricMean(input []float64) float64 {

	// Get the product of all the numbers
	var p float64
	for _, n := range input {
		if p == 0 {
			p = n
		} else {
			p *= n
		}
	}

	// Calculate the geometric mean
	return math.Pow(p, 1/float64(len(input)))
}

// HarmonicMean gets the harmonic mean for a slice of numbers
func HarmonicMean(input []float64) (float64, error) {

	// Get the sum of all the numbers reciprocals and return an
	// error for values that cannot be included in harmonic mean
	var p float64
	for _, n := range input {
		if n < 0 {
			return 0, errors.New("Input must not contain a negative number")
		} else if n == 0 {
			return 0, errors.New("Input must not contain a zero value")
		}
		p += (1 / n)
	}

	return float64(len(input)) / p, nil
}
