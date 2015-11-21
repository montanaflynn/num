package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3}, 6},
		{[]float64{1, 1.1, 1.2, 2.2}, 5.5},
		{[]float64{1, -1, 2, -3}, -1},
	} {
		got := Sum(c.in)
		if c.out != got {
			t.Errorf("Sum(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
}
