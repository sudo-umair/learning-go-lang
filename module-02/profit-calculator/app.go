package main

import "fmt"

func main() {
	revenue := getInputFor("Enter revenue: ")
	expenses := getInputFor("Enter expenses: ")
	taxRate := getInputFor("Enter tax rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Earnings before tax to profit ratio: %.2f\n", ratio)
}

func getInputFor(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
