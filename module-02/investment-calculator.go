package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years := 1000.0, 5.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("Future value of investment:", futureValue)
}
