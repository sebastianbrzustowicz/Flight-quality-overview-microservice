package main

import (
	"fmt"

	"sync"

	"github.com/jung-kurt/gofpdf"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot, errorRollPlot, errorPitchPlot, errorYawPlot, errorAltitudePlot *plot.Plot,
	rollRMS, pitchRMS, yawRMS, altitudeRMS, rollSSE, pitchSSE, yawSSE, altitudeSSE string) error {

	// Creating new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// New page (first)
	pdf.AddPage()

	// Title for page
	pdf.SetFont("times", "B", 14)
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - control quality indicators")
	pdf.Ln(10)
	// Adding variables to PDF
	name := "Sebastian Brzustowicz"
	email := "Se.Brzustowicz@gmail.com"
	additionalInfo := "This microservice is distributed under MIT licencse."

	pdf.SetFont("times", "B", 10)
	pdf.Cell(0, 10, fmt.Sprintf("Author: %s", name))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("E-mail: %s", email))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Additional info: %s", additionalInfo))
	pdf.SetFont("times", "", 10)
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("In order to present the results of the control method RMS indicator was applied, which is the root mean square error."))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("This indicator shows what the spread of values around the mean error value is."))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("The RMS criterion for each rotary axes is defined as follows:"))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("J_RMS = sqrt(sum((e_n-m)^2)/N)"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("where"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("J_RMS - RMS quality indicator"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("e - actual error"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("m - mean error"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("N - number of samples"))
	pdf.Ln(10)

	// Headers setting for table
	pdf.SetFont("times", "B", 12)
	pdf.SetFillColor(200, 220, 255)

	// First row
	pdf.CellFormat(40, 10, "Measured variable", "1", 0, "", true, 0, "")
	pdf.CellFormat(100, 10, "RMS indicator", "1", 0, "", true, 0, "")

	// End of row
	pdf.Ln(-1)

	// Table style
	pdf.SetFont("times", "", 12)
	pdf.SetFillColor(255, 255, 255)

	// Table data
	pdf.CellFormat(40, 10, fmt.Sprintf("Roll"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(rollRMS), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(40, 10, fmt.Sprintf("Pitch"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(pitchRMS), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(40, 10, fmt.Sprintf("Yaw"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(yawRMS), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(40, 10, fmt.Sprintf("Altitude"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(altitudeRMS), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(5)

	pdf.SetFont("times", "", 10)
	pdf.Cell(0, 10, fmt.Sprintf("The second indicator is the sum of squared errors."))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Sum of squared errors for each rotary axes is defined as follows:"))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("J_SSE = sum((e_n)^2))"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("where"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("J_SSE - quality indicator"))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("e - actual error"))
	pdf.Ln(10)

	// Headers setting for table
	pdf.SetFont("times", "B", 12)
	pdf.SetFillColor(200, 220, 255)

	// First row
	pdf.CellFormat(40, 10, "Measured variable", "1", 0, "", true, 0, "")
	pdf.CellFormat(100, 10, "SSE indicator", "1", 0, "", true, 0, "")

	// End of row
	pdf.Ln(-1)

	// Table style
	pdf.SetFont("times", "", 12)
	pdf.SetFillColor(255, 255, 255)

	// Table data
	pdf.CellFormat(40, 10, fmt.Sprintf("Roll"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(rollSSE), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(40, 10, fmt.Sprintf("Pitch"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(pitchSSE), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(40, 10, fmt.Sprintf("Yaw"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(yawSSE), "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(40, 10, fmt.Sprintf("Altitude"), "1", 0, "", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf(altitudeSSE), "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	// New page (second)
	pdf.AddPage()

	// Title for page
	pdf.SetFont("times", "B", 14)
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - control visualisation")
	pdf.SetFont("times", "B", 10)
	pdf.Ln(10)

	// Goroutines
	var wg sync.WaitGroup

	wg.Add(4)
	pngRollPlot := "rollPlot.png"
	go savePlotAsync(rollPlot, pngRollPlot, &wg)
	pngPitchPlot := "pitchPlot.png"
	go savePlotAsync(pitchPlot, pngPitchPlot, &wg)
	pngYawPlot := "yawPlot.png"
	go savePlotAsync(yawPlot, pngYawPlot, &wg)
	pngAltitudePlot := "altitudePlot.png"
	go savePlotAsync(altitudePlot, pngAltitudePlot, &wg)
	wg.Wait()

	pdf.Ln(10)
	pdf.Image(pngRollPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetRoll := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetRoll))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.1: Roll data"))
	pdf.Ln(12)
	pdf.Image(pngPitchPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetPitch := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetPitch))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.2: Pitch data"))
	pdf.Ln(12)
	pdf.Image(pngYawPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetYaw := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetYaw))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.3: Yaw data"))
	pdf.Ln(12)
	pdf.Image(pngAltitudePlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetAltitude := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetAltitude))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.4: Altitude data"))

	// New page (third)
	pdf.AddPage()

	// Title for page
	pdf.SetFont("times", "B", 14)
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - control errors")
	pdf.SetFont("times", "B", 10)
	pdf.Ln(10)

	wg.Add(4)
	pngRollErrorPlot := "rollErrorPlot.png"
	go saveErrorPlotAsync(errorRollPlot, pngRollErrorPlot, &wg)
	pngPitchErrorPlot := "pitchErrorPlot.png"
	go saveErrorPlotAsync(errorPitchPlot, pngPitchErrorPlot, &wg)
	pngYawErrorPlot := "yawErrorPlot.png"
	go saveErrorPlotAsync(errorYawPlot, pngYawErrorPlot, &wg)
	pngAltitudeErrorPlot := "altitudeErrorPlot.png"
	go saveErrorPlotAsync(errorAltitudePlot, pngAltitudeErrorPlot, &wg)
	wg.Wait()

	pdf.Ln(10)
	pdf.Image(pngRollErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetRollError := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetRollError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.1: Roll error data"))
	pdf.Ln(12)
	pdf.Image(pngPitchErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetPitchError := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetPitchError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.2: Pitch error data"))
	pdf.Ln(12)
	pdf.Image(pngYawErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetYawError := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetYawError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.3: Yaw error data"))
	pdf.Ln(12)
	pdf.Image(pngAltitudeErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetAltitudeError := 79
	pdf.SetX(pdf.GetX() + float64(xOffsetAltitudeError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.4: Altitude error data"))

	// Saving pdf to file
	err := pdf.OutputFileAndClose("Data-report.pdf")
	if err != nil {
		return fmt.Errorf("Error when saving file to PDF: %v", err)
	}

	return nil
}

func savePlotAsync(plot *plot.Plot, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := plot.Save(10*vg.Inch, 3*vg.Inch, filename); err != nil {
		fmt.Printf("Error when saving plot to PNG (%s): %v\n", filename, err)
	}
}

func saveErrorPlotAsync(errorPlot *plot.Plot, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := errorPlot.Save(10*vg.Inch, 3*vg.Inch, filename); err != nil {
		fmt.Printf("Error when saving error plot to PNG (%s): %v\n", filename, err)
	}
}
