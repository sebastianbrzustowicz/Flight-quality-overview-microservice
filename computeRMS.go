package main

import (
	"math"
	"strconv"
)

func computeRMS(errorArr []float64) (string, error) {
	if len(errorArr) == 0 {
		return "Empty data", nil
	}

	sum := 0.0
	for _, value := range errorArr {
		sum += value
	}

	average := sum / float64(len(errorArr))

	newArr := make([]float64, len(errorArr))
	for i, value := range errorArr {
		newArr[i] = value - average
	}

	sumOfSquares := 0.0

	for _, num := range newArr {
		sumOfSquares += math.Pow(num, 2)
	}

	averageSumOfSquares := sumOfSquares / float64(len(errorArr))

	JRMS := math.Sqrt(averageSumOfSquares)

	result := strconv.FormatFloat(JRMS, 'f', -1, 64)
	return result, nil
}
