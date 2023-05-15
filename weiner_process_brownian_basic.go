package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func BrownianBasic() {
	// Seed the random number generator with the current time for diff results each time
	rand.Seed(time.Now().UnixNano())

	timeStep := 0.01
	timeFinal := 1.0

	numberOfSteps := int(timeFinal / timeStep)

	// Generate a slice of array with the number of all steps and capacity of all steps in platter XYs values
	var weinerProcessSeriesPlotValues plotter.XYs = make(plotter.XYs, numberOfSteps+1)
	weinerProcessSeriesPlotValues[0].X = 0.0
	weinerProcessSeriesPlotValues[0].Y = 0.0

	// Generate the weiner process for graph plotting
	for i := 1; i < numberOfSteps+1; i++ {
		randomVar := math.Sqrt(timeStep) * float64(rand.NormFloat64())
		weinerProcessSeriesPlotValues[i].X = float64(i)
		// Receive last value and add the random variable to it
		weinerProcessSeriesPlotValues[i].Y = weinerProcessSeriesPlotValues[i-1].Y + randomVar
		fmt.Println(weinerProcessSeriesPlotValues[i])
	}

	// Create a new plot
	p := plot.New()
	p.Title.Text = "Weiner Process Brownian Motion Simple Example"
	p.X.Label.Text = "Step"
	p.Y.Label.Text = "Value"

	// Create a plotter.Values value and fill it with the values from the respective series.
	line, err := plotter.NewLine(weinerProcessSeriesPlotValues)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{G: 128, A: 255}
	p.Add(line)

	if err := p.Save(7*vg.Inch, 5*vg.Inch, "weiner_process_brownian_basic.png"); err != nil {
		panic(err)
	}

	fmt.Println("Graph image generated successfully.")
}
