package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func generateAnglePlot(arr1, arr2 []float64, legend1, legend2, xlabel, ylabel string) (*plot.Plot, error) {

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
	line1.LineStyle.Color = color.RGBA{R: 0, G: 114, B: 189, A: 255}
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
	line2.LineStyle.Color = color.RGBA{R: 217, G: 83, B: 25, A: 255}
	p.Add(line2)

	// Adding legend (position and value)
	p.Legend.Top = true
	p.Legend.Left = false
	p.Legend.XOffs = 0
	p.Legend.YOffs = -0.1
	p.Legend.Add(legend1, line1)
	p.Legend.Add(legend2, line2)

	// Adding labels
	p.X.Label.Text = xlabel
	p.Y.Label.Text = ylabel

	// Adding grid
	p.Add(plotter.NewGrid())

	return p, nil
}
