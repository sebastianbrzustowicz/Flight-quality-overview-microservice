package main

import (
	"fmt"

	"encoding/csv"
	"os"
	"strconv"

	"image/color"

	"github.com/jung-kurt/gofpdf"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	rolld, pitchd, yawd, altituded, roll, pitch, yaw, altitude, isClamp, err := readCSVFile("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Displaying data (sample)
	//fmt.Println("Column 1:", rolld)
	//fmt.Println("Column 2:", pitchd)
	//fmt.Println("Column 3:", yawd)
	//fmt.Println("Column 4:", altituded)
	//fmt.Println("Column 5:", roll)
	//fmt.Println("Column 6:", pitch)
	//fmt.Println("Column 7:", yaw)
	//fmt.Println("Column 8:", altitude)
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

func generateAndSavePlot(arr1, arr2 []float64) (*plot.Plot, error) {

	if len(arr2) != len(arr1) {
		return nil, fmt.Errorf("Data is incorrect: different sizes")
	}

	length := len(arr1)

	samplesArray := make([]float64, length)
	for i := 0; i < length; i++ {
		samplesArray[i] = float64(i + 1)
	}

	var resultArray1 []struct {
		X, Y float64
	}
	var resultArray2 []struct {
		X, Y float64
	}

	for i := 0; i < len(samplesArray); i++ {
		resultArray1 = append(resultArray1, struct{ X, Y float64 }{samplesArray[i], arr1[i]})
	}
	for i := 0; i < len(samplesArray); i++ {
		resultArray2 = append(resultArray2, struct{ X, Y float64 }{samplesArray[i], arr2[i]})
	}

	// Creating new plot
	p := plot.New()

	// Adding data to plot
	points := make(plotter.XYs, len(resultArray1))
	for i, d := range resultArray1 {
		points[i].X = d.X
		points[i].Y = d.Y
	}
	line1, err := plotter.NewLine(points)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating plot: %v", err)
	}
	line1.LineStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	p.Add(line1)

	points2 := make(plotter.XYs, len(resultArray2))
	for i, d := range resultArray2 {
		points2[i].X = d.X
		points2[i].Y = d.Y
	}
	line2, err := plotter.NewLine(points2)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating plot: %v", err)
	}
	line2.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	p.Add(line2)

	// Saving plot to file
	//if err := p.Save(4*vg.Inch, 4*vg.Inch, "graph.png"); err != nil {
	//	return fmt.Errorf("error during saving file: %v", err)
	//}
	//fmt.Println("Plot created successfully and saved to file 'graph.png'.")

	return p, nil
}

func generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot *plot.Plot) error {
	// Creating new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Font size
	pdf.SetFont("Arial", "B", 16)

	// Title for PDF
	pdf.Cell(40, 10, "Data to PDF raport microservice")

	// Adding variables to PDF
	//name := "Sebastian Brzustowicz"
	//title := "Programmer"
	//email := "me.example@gmail.com"
	//
	//pdf.Ln(10) // Dodanie pustego wiersza
	pdf.SetFont("Arial", "", 12)
	//pdf.Cell(0, 10, fmt.Sprintf("Name and surname: %s", name))
	//pdf.Ln(10)
	//pdf.Cell(0, 10, fmt.Sprintf("Title: %d", title))
	//pdf.Ln(10)
	//pdf.Cell(0, 10, fmt.Sprintf("E-mail: %s", email))

	pngRollPlot := "rollPlot.png"
	if err := rollPlot.Save(10*vg.Inch, 3*vg.Inch, pngRollPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pngPitchPlot := "pitchPlot.png"
	if err := pitchPlot.Save(10*vg.Inch, 3*vg.Inch, pngPitchPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pngYawPlot := "yawPlot.png"
	if err := rollPlot.Save(10*vg.Inch, 3*vg.Inch, pngYawPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pngAltitudePlot := "altitudePlot.png"
	if err := altitudePlot.Save(10*vg.Inch, 3*vg.Inch, pngAltitudePlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pdf.Ln(10)
	pdf.Image(pngRollPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetRoll := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetRoll))
	pdf.Cell(210, 10, fmt.Sprintf("Roll data"))
	pdf.Ln(10)
	pdf.Image(pngPitchPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetPitch := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetPitch))
	pdf.Cell(210, 10, fmt.Sprintf("Pitch data"))
	pdf.Ln(10)
	pdf.Image(pngYawPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetYaw := 86
	pdf.SetX(pdf.GetX() + float64(xOffsetYaw))
	pdf.Cell(210, 10, fmt.Sprintf("Yaw data"))
	pdf.Ln(10)
	pdf.Image(pngAltitudePlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetAltitude := 83
	pdf.SetX(pdf.GetX() + float64(xOffsetAltitude))
	pdf.Cell(210, 10, fmt.Sprintf("Altitude data"))
	pdf.Ln(10)

	// Zapisanie dokumentu do pliku
	err := pdf.OutputFileAndClose("Data-report.pdf")
	if err != nil {
		return fmt.Errorf("Error when saving file to PDF: %v", err)
	}

	fmt.Println("PDF created successfully.")
	return nil
}
