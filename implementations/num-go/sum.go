package main

// Sum adds all the numbers together
// https://en.wikipedia.org/wiki/Summation
func Sum(input []float64) (sum float64) {
	for _, n := range input {
		sum += n
	}
	return sum
}
