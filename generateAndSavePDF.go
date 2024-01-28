package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot, errorRollPlot, errorPitchPlot, errorYawPlot, errorAltitudePlot *plot.Plot) error {
	// Creating new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// New page (first)
	pdf.AddPage()

	// Title for page
	pdf.SetFont("times", "B", 14)
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - quality indicators")
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

	// Integral criterion here (todo)

	// New page (second)
	pdf.AddPage()

	// Title for page
	pdf.SetFont("times", "B", 14)
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - control visualisation")
	pdf.SetFont("times", "B", 10)
	pdf.Ln(10)

	pngRollPlot := "rollPlot.png"
	if err := rollPlot.Save(10*vg.Inch, 3*vg.Inch, pngRollPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pngPitchPlot := "pitchPlot.png"
	if err := pitchPlot.Save(10*vg.Inch, 3*vg.Inch, pngPitchPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pngYawPlot := "yawPlot.png"
	if err := yawPlot.Save(10*vg.Inch, 3*vg.Inch, pngYawPlot); err != nil {
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
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.1: Roll data"))
	pdf.Ln(10)
	pdf.Image(pngPitchPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetPitch := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetPitch))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.2: Pitch data"))
	pdf.Ln(10)
	pdf.Image(pngYawPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetYaw := 85
	pdf.SetX(pdf.GetX() + float64(xOffsetYaw))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.3: Yaw data"))
	pdf.Ln(10)
	pdf.Image(pngAltitudePlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetAltitude := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetAltitude))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 1.4: Altitude data"))
	pdf.Ln(10)

	// New page (third)
	pdf.AddPage()

	// Title for page
	pdf.SetFont("times", "B", 14)
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - control errors")
	pdf.SetFont("times", "B", 10)
	pdf.Ln(10)

	pngRollErrorPlot := "rollErrorPlot.png"
	if err := errorRollPlot.Save(10*vg.Inch, 3*vg.Inch, pngRollErrorPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}
	pngPitchErrorPlot := "pitchErrorPlot.png"
	if err := errorPitchPlot.Save(10*vg.Inch, 3*vg.Inch, pngPitchErrorPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}
	pngYawErrorPlot := "yawErrorPlot.png"
	if err := errorYawPlot.Save(10*vg.Inch, 3*vg.Inch, pngYawErrorPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}
	pngAltitudeErrorPlot := "altitudeErrorPlot.png"
	if err := errorAltitudePlot.Save(10*vg.Inch, 3*vg.Inch, pngAltitudeErrorPlot); err != nil {
		return fmt.Errorf("Error when saving plot to PNG: %v", err)
	}

	pdf.Ln(10)
	pdf.Image(pngRollErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetRollError := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetRollError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.1: Roll error data"))
	pdf.Ln(10)
	pdf.Image(pngPitchErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetPitchError := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetPitchError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.2: Pitch error data"))
	pdf.Ln(10)
	pdf.Image(pngYawErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetYawError := 82
	pdf.SetX(pdf.GetX() + float64(xOffsetYawError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.3: Yaw error data"))
	pdf.Ln(10)
	pdf.Image(pngAltitudeErrorPlot, 10, pdf.GetY(), 190, 50, true, "", 0, "")
	xOffsetAltitudeError := 79
	pdf.SetX(pdf.GetX() + float64(xOffsetAltitudeError))
	pdf.Cell(210, 10, fmt.Sprintf("Fig. 2.4: Altitude error data"))

	// Zapisanie dokumentu do pliku
	err := pdf.OutputFileAndClose("Data-report.pdf")
	if err != nil {
		return fmt.Errorf("Error when saving file to PDF: %v", err)
	}

	fmt.Println("PDF created successfully.")
	return nil
}
