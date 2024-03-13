package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 5

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Println("Future value of investment:", futureValue)
}
