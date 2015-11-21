package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
	"strconv"
)

func printResultOrError(result interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func printAndExit(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {

	if len(os.Args) < 3 {
		printAndExit("Missing function and data")
	}

	var cmd = os.Args[1]

	if len(os.Args) == 3 {
		datum, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			printAndExit(err)
		}
		if cmd == "round" {
			printResultOrError(stats.Round(datum, 0))
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
			printResultOrError(stats.Sum(data))
		}

		if cmd == "mean" || cmd == "average" {
			printResultOrError(stats.Mean(data))
		}

		if cmd == "gmean" || cmd == "geometric-mean" {
			printResultOrError(stats.GeometricMean(data))
		}

		if cmd == "hmean" || cmd == "harmonic-mean" {
			printResultOrError(stats.HarmonicMean(data))
		}
	}
}
