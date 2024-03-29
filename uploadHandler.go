package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"gonum.org/v1/plot"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}

	fmt.Println("Client IP address:", ip)

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create new csv file on server
	serverFile, err := os.Create("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
		return
	}
	defer serverFile.Close()

	// Copy the content of the form file to a file on the server
	_, err = io.Copy(serverFile, file)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "The contents of the file could not be copied", http.StatusInternalServerError)
		return
	}

	rolld, pitchd, yawd, altituded, roll, pitch, yaw, altitude, isClamp, err := readCSVFile("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_ = isClamp
	// Displaying unused data (sample)
	//fmt.Println("Column 9:", isClamp)

	//timeBeforePlots := time.Since(startTime).Microseconds()
	//fmt.Printf("Time before plots: %v microseconds\n", timeBeforePlots)

	//rollPlotChan := make(chan *plot.Plot)
	ch1 := make(chan *plot.Plot)
	ch2 := make(chan *plot.Plot)
	ch3 := make(chan *plot.Plot)
	ch4 := make(chan *plot.Plot)
	var wg sync.WaitGroup
	wg.Add(4)

	go generateAnglePlot(rolld, roll, "Roll_d", "Roll", "Sample", "Angle [rad]", ch1, &wg)
	go generateAnglePlot(pitchd, pitch, "Pitch_d", "Pitch", "Sample", "Angle [rad]", ch2, &wg)
	go generateAnglePlot(yawd, yaw, "Yaw_d", "Yaw", "Sample", "Angle [rad]", ch3, &wg)
	go generateAnglePlot(altituded, altitude, "Altitude_d", "Altitude", "Sample", "Altitude [m]", ch4, &wg)

	rollPlot, pitchPlot, yawPlot, altitudePlot := <-ch1, <-ch2, <-ch3, <-ch4
	close(ch1)
	close(ch2)
	close(ch3)
	close(ch4)
	wg.Wait()

	//timeAfterPlots := time.Since(startTime).Microseconds()
	//fmt.Printf("Time after plots: %v microseconds\n", timeAfterPlots)

	// Calculating errors
	var errorRoll []float64
	for i := 0; i < len(rolld); i++ {
		errorRoll = append(errorRoll, rolld[i]-roll[i])
	}
	var errorPitch []float64
	for i := 0; i < len(pitchd); i++ {
		errorPitch = append(errorPitch, pitchd[i]-pitch[i])
	}
	var errorYaw []float64
	for i := 0; i < len(yawd); i++ {
		errorYaw = append(errorYaw, yawd[i]-yaw[i])
	}
	var errorAltitude []float64
	for i := 0; i < len(altituded); i++ {
		errorAltitude = append(errorAltitude, altituded[i]-altitude[i])
	}

	//timeBeforeErrorPlots := time.Since(startTime).Microseconds()
	//fmt.Printf("Time before error plots: %v microseconds\n", timeBeforeErrorPlots)

	errorRollPlot, err := generateErrorPlot(errorRoll, "Roll_e", "Sample", "Angle error [rad]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	errorPitchPlot, err := generateErrorPlot(errorPitch, "Pitch_e", "Sample", "Angle error [rad]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	errorYawPlot, err := generateErrorPlot(errorYaw, "Yaw_e", "Sample", "Angle error [rad]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	errorAltitudePlot, err := generateErrorPlot(errorAltitude, "Altitude_e", "Sample", "Altitude error [m]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//timeAfterErrorPlots := time.Since(startTime).Microseconds()
	//fmt.Printf("Time after error plots: %v microseconds\n", timeAfterErrorPlots)

	rollRMS, rmsErr := computeRMS(errorRoll)
	pitchRMS, rmsErr := computeRMS(errorPitch)
	yawRMS, rmsErr := computeRMS(errorYaw)
	altitudeRMS, rmsErr := computeRMS(errorAltitude)
	rollSSE, sseErr := computeSSE(errorRoll)
	pitchSSE, sseErr := computeSSE(errorPitch)
	yawSSE, sseErr := computeSSE(errorYaw)
	altitudeSSE, sseErr := computeSSE(errorAltitude)
	if rmsErr != nil {
		fmt.Println("Error computing RMS:", rmsErr)
		return
	}
	if sseErr != nil {
		fmt.Println("Error computing SSE:", sseErr)
		return
	}

	//timeAfterIndicators := time.Since(startTime).Microseconds()
	//fmt.Printf("Time after indicators: %v microseconds\n", timeAfterIndicators)

	if err := generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot, errorRollPlot, errorPitchPlot, errorYawPlot, errorAltitudePlot,
		rollRMS, pitchRMS, yawRMS, altitudeRMS, rollSSE, pitchSSE, yawSSE, altitudeSSE); err != nil {
		fmt.Println("Error:", err)
	}

	pdfContent, err := ioutil.ReadFile("Data-report.pdf")
	if err != nil {
		http.Error(w, "Error reading PDF content", http.StatusInternalServerError)
		return
	}

	//timeAfterPDFGeneration := time.Since(startTime).Microseconds()
	//fmt.Printf("Time after PDF generation: %v microseconds\n", timeAfterPDFGeneration)

	// Set response headers
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=Data-report.pdf")

	// Send the modified PDF as response
	w.Write(pdfContent)

	filesToRemove := []string{"altitudePlot.png", "rollPlot.png", "yawPlot.png", "pitchPlot.png",
		"rollErrorPlot.png", "pitchErrorPlot.png", "yawErrorPlot.png", "altitudeErrorPlot.png", "Data-report.pdf"}

	for _, file := range filesToRemove {
		err := os.Remove(file)
		if err != nil {
			fmt.Println("Error deleting a file", file, ":", err)
		}
	}

	executionTime := time.Since(startTime).Microseconds()
	fmt.Printf("Execution time: %v microseconds\n", executionTime)

}
