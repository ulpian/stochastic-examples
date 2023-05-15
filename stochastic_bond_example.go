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

func BondBasic() {
	rand.Seed(time.Now().UnixNano())

	// Parameters for a simple bond with fixed rate term
	termInYears := 15.00    // Bond term in years
	initialPrice := 1000.00 // Par value
	interestRate := 0.05    // Interest rate
	volatility := 0.08      // Volatility

	numberOfSimulations := 3
	numberOfSteps := 12 * int(termInYears)
	step := termInYears / float64(numberOfSteps)
	discountFactor := 1.00 / (1.00 + interestRate*step)

	simulationPlot := make([]plotter.XYs, numberOfSimulations)

	// Performing simulations
	for x := 0; x < numberOfSimulations; x++ {
		bondPrice := initialPrice

		// Graph plotting
		simulationPoint := make(plotter.XYs, int(numberOfSteps)+1)
		simulationPoint[0].X = 0.0
		simulationPoint[0].Y = bondPrice

		for s := 0; s <= numberOfSteps; s++ {
			driftPrice := interestRate * bondPrice * float64(step)
			randomVolatility := volatility * bondPrice * math.Sqrt(step) * float64(rand.NormFloat64())

			bondPrice = (bondPrice + driftPrice + randomVolatility) * discountFactor
			// Add to graph
			simulationPoint[s].X = float64(s)
			simulationPoint[s].Y = bondPrice
		}

		simulationPlot[x] = simulationPoint
		fmt.Printf("Simulation Cycle %d: Bond price = %.2f\n", x+1, bondPrice)
	}

	p := plot.New()
	p.Title.Text = "Stochastic Bond Price Example Simulation"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Bond Price"

	for x := 0; x < numberOfSimulations; x++ {
		line, err := plotter.NewLine(simulationPlot[x])
		if err != nil {
			panic(err)
		}
		line.Color = color.RGBA{G: 128, A: 255}
		p.Add(line)
	}

	// Save the plot to an image file
	if err := p.Save(10*vg.Inch, 5*vg.Inch, "bond_price_simulations.png"); err != nil {
		panic(err)
	}

	fmt.Println("Graph image generated successfully.")
}
