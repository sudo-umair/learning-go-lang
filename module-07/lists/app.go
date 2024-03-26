package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"Mobile", "Tablet", "Laptop", "Desktop"}
	prices := [4]float64{1.23, 2.34, 3.45, 0.0}

	// fmt.Println(prices)
	// fmt.Println(productNames[0])
	// productNames[0] = "Laptop"
	// fmt.Println(productNames[0])

	featuredPrices := prices[1:3]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)

	featuredProducts1 := productNames[:3]
	fmt.Println(featuredProducts1)

	featuredProducts2 := productNames[1:]
	fmt.Println(featuredProducts2)
}
