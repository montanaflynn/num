package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	command := os.Args[1]
	data := os.Args[2:]
	if command == "sum" {
		var floatData []float64
		for i := 0; i < len(data); i++ {
			datum, err := strconv.ParseFloat(data[i], 64)
			if err == nil {
				floatData = append(floatData, datum)
			}

		}
		fmt.Println(Sum(floatData))
	}
}
