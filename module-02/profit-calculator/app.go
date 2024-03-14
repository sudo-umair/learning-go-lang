package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile = "results.txt"

func main() {
	revenue, err1 := getInputFor("Enter revenue: ")
	expenses, err2 := getInputFor("Enter expenses: ")
	taxRate, err3 := getInputFor("Enter tax rate: ")

	// using if statement to check for all errors in one go
	if (err1 != nil) || (err2 != nil) || (err3 != nil) {
		fmt.Println("ERROR!")
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println("-----------------------------------")
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	storeResultsInFile(ebt, profit, ratio)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Earnings before tax to profit ratio: %.2f\n", ratio)
}

func getInputFor(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Invalid amount! Must be greater than 0")
	}

	return value, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResultsInFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("Earnings before tax: %.2f\nProfit: %.2f\nEarnings before tax to profit ratio: %.2f\n", ebt, profit, ratio)
	os.WriteFile(resultsFile, []byte(results), 0644)
}
