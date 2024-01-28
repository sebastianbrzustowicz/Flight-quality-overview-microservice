package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

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

	rollPlot, err := generateAndSavePlot(rolld, roll, "Roll_d", "Roll", "Sample", "Angle [rad]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	pitchPlot, err := generateAndSavePlot(pitchd, pitch, "Pitch_d", "Pitch", "Sample", "Angle [rad]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	yawPlot, err := generateAndSavePlot(yawd, yaw, "Yaw_d", "Yaw", "Sample", "Angle [rad]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	altitudePlot, err := generateAndSavePlot(altituded, altitude, "Altitude_d", "Altitude", "Sample", "Altitude [m]")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot); err != nil {
		fmt.Println("Error:", err)
	}

	pdfContent, err := ioutil.ReadFile("Data-report.pdf")
	if err != nil {
		http.Error(w, "Error reading PDF content", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=Data-report.pdf")

	// Send the modified PDF as response
	w.Write(pdfContent)

	filesToRemove := []string{"altitudePlot.png", "rollPlot.png", "yawPlot.png", "pitchPlot.png", "Data-report.pdf"}

	for _, file := range filesToRemove {
		err := os.Remove(file)
		if err != nil {
			fmt.Println("Error deleting a file", file, ":", err)
		}
	}
}
