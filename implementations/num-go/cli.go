package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

// Num can take commands through cli and data
// through both cli or piped in from stdin
//
// Here's the fucking logic:
// 1) Check for commands
// 2) Check for data through cli
// 3) Check for data through stdin
// 4) Calculate commands and display results

func main() {

	var data []float64
	var sin = stdinData()
	var arg = argData()

	if len(os.Args) == 1 {
		printAndExit("Missing commands, see --help")
	}

	if len(os.Args) == 2 && len(sin) == 0 {
		printAndExit("Missing data, see --help")
	}

	data = append(data, sin...)
	data = append(data, arg...)

	var cmd = os.Args[1]

	if len(data) == 1 {
		if cmd == "round" {
			printResultOrError(stats.Round(data[0], 0))
		}
	}

	if len(data) > 1 {
		if cmd == "sum" || cmd == "total" {
			printResultOrError(stats.Sum(data))
		}

		if cmd == "mean" || cmd == "average" || cmd == "avg" {
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

func argData() (data []float64) {
	for i := 0; i < len(os.Args[2:]); i++ {
		datum, err := strconv.ParseFloat(os.Args[2:][i], 64)
		if err == nil {
			data = append(data, datum)
		}
	}
	return
}

func stdinData() (data []float64) {

	// let's just grab it all at once shall we
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return []float64{}
	}

	// Clean up any space and convert to a string
	clean := strings.TrimSpace(string(stdin))

	// break up the numbers into individual nums
	pieces := strings.Split(clean, " ")

	for i := 0; i < len(pieces); i++ {

		// for each number we convert to float64
		// ignoring any errors (non-data) we may find
		datum, err := strconv.ParseFloat(pieces[i], 64)
		if err == nil {
			data = append(data, datum)
		}
	}

	return
}

func printResultOrError(result interface{}, err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", result)
}

func printAndExit(msg interface{}) {
	fmt.Printf("%+v\n", msg)
	os.Exit(1)
}
