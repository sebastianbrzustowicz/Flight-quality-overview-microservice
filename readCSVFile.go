package main

import (
	"fmt"

	"encoding/csv"
	"os"
	"strconv"
)

func readCSVFile(filename string) (col1, col2, col3, col4, col5, col6, col7, col8 []float64, col9 []bool, err error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Read data from the file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error reading data from file: %v", err)
	}

	// Process each record
	for _, record := range records {
		// Convert values to the appropriate types
		val1, _ := strconv.ParseFloat(record[0], 64)
		val2, _ := strconv.ParseFloat(record[1], 64)
		val3, _ := strconv.ParseFloat(record[2], 64)
		val4, _ := strconv.ParseFloat(record[3], 64)
		val5, _ := strconv.ParseFloat(record[4], 64)
		val6, _ := strconv.ParseFloat(record[5], 64)
		val7, _ := strconv.ParseFloat(record[6], 64)
		val8, _ := strconv.ParseFloat(record[7], 64)
		val9, _ := strconv.ParseBool(record[8])

		// Add values to the respective arrays
		col1 = append(col1, val1)
		col2 = append(col2, val2)
		col3 = append(col3, val3)
		col4 = append(col4, val4)
		col5 = append(col5, val5)
		col6 = append(col6, val6)
		col7 = append(col7, val7)
		col8 = append(col8, val8)
		col9 = append(col9, val9)
	}

	return col1, col2, col3, col4, col5, col6, col7, col8, col9, nil
}
