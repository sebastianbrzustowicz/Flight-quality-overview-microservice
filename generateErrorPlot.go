package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func generateErrorPlot(arr1 []float64, legend1, xlabel, ylabel string) (*plot.Plot, error) {

	length := len(arr1)

	samplesArray := make([]float64, length)
	for i := 0; i < length; i++ {
		samplesArray[i] = float64(i + 1)
	}

	var resultArray1 []struct {
		X, Y float64
	}

	for i := 0; i < len(samplesArray); i++ {
		resultArray1 = append(resultArray1, struct{ X, Y float64 }{samplesArray[i], arr1[i]})
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
	line1.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	p.Add(line1)

	// Adding legend (position and value)
	p.Legend.Top = true
	p.Legend.Left = false
	p.Legend.XOffs = 0
	p.Legend.YOffs = -0.1
	p.Legend.Add(legend1, line1)

	// Adding labels
	p.X.Label.Text = xlabel
	p.Y.Label.Text = ylabel

	// Adding grid
	p.Add(plotter.NewGrid())

	return p, nil
}
