package main

import (
	cmdmanager "com/calculator/cmd-manager"
	"com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("output-%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
