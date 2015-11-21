package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Missing function and data")
		os.Exit(1)
	}

	var cmd = os.Args[1]

	if len(os.Args) == 3 {
		datum, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if cmd == "round" {
			result, err := Round(datum, 1)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(result)
		}
	}

	if len(os.Args) > 3 {
		var data []float64
		for i := 0; i < len(os.Args[2:]); i++ {
			datum, err := strconv.ParseFloat(os.Args[2:][i], 64)
			if err == nil {
				data = append(data, datum)
			}
		}
		if cmd == "sum" || cmd == "total" {
			fmt.Println(Sum(data))
		}

		if cmd == "mean" || cmd == "average" {
			fmt.Println(Mean(data))
		}

		if cmd == "gmean" || cmd == "geometric-mean" {
			fmt.Println(GeometricMean(data))
		}

		if cmd == "hmean" || cmd == "harmonic-mean" {
			result, err := HarmonicMean(data)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(result)
		}
	}
}
