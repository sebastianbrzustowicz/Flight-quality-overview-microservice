package main

import (
	"math"
	"strconv"
)

func computeSSE(errorArr []float64) (string, error) {
	if len(errorArr) == 0 {
		return "Empty data", nil
	}

	sumOfSquares := 0.0

	for _, num := range errorArr {
		sumOfSquares += math.Pow(num, 2)
	}

	result := strconv.FormatFloat(sumOfSquares, 'f', -1, 64)
	return result, nil
}
