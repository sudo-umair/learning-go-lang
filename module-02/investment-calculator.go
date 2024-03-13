package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFutureValue := fmt.Sprintf("%.2f", futureValue)
	formattedFutureRealValue := fmt.Sprintf("%.2f", futureRealValue)

	fmt.Println("Future value of investment:", formattedFutureValue)
	fmt.Println("Future real value of investment:", formattedFutureRealValue)
}
