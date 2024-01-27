package main

import (
	"fmt"
)

func main() {

	rolld, pitchd, yawd, altituded, roll, pitch, yaw, altitude, isClamp, err := readCSVFile("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Displaying unused data (sample)
	fmt.Println("Column 9:", isClamp)

	rollPlot, err := generateAndSavePlot(rolld, roll)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	pitchPlot, err := generateAndSavePlot(pitchd, pitch)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	yawPlot, err := generateAndSavePlot(yawd, yaw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	altitudePlot, err := generateAndSavePlot(altituded, altitude)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot); err != nil {
		fmt.Println("Error:", err)
	}
}
