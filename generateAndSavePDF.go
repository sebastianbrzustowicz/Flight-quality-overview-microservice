package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func generateAndSavePDF(rollPlot, pitchPlot, yawPlot, altitudePlot *plot.Plot) error {
	// Creating new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Font size
	pdf.SetFont("times", "B", 14)

	// Title for PDF
	pdf.Cell(40, 10, "Filght data to PDF raport microservice - plot visualisation")
	pdf.Ln(10)
	// Adding variables to PDF
	//name := "Sebastian Brzustowicz"
	//title := "Programmer"
	//email := "me.example@gmail.com"
	//
	//pdf.Ln(10) // Dodanie pustego wiersza
	pdf.SetFont("times", "B", 10)
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

	// Zapisanie dokumentu do pliku
	err := pdf.OutputFileAndClose("Data-report.pdf")
	if err != nil {
		return fmt.Errorf("Error when saving file to PDF: %v", err)
	}

	fmt.Println("PDF created successfully.")
	return nil
}
