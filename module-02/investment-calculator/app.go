package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter number of years: ")
	fmt.Scan(&years)

	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFutureValue := fmt.Sprintf("%.2f", futureValue)
	formattedFutureRealValue := fmt.Sprintf("%.2f", futureRealValue)

	fmt.Println("Future value of investment:", formattedFutureValue)
	fmt.Println("Future real value of investment:", formattedFutureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}
