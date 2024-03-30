package main

import "fmt"

func main() {
	prices := []float64{10.99, 4.5}
	fmt.Println("prices:", prices[0:1])

	// prices[0] = 5.99
	// fmt.Println("prices:", prices)

	prices = append(prices, 3.5)
	fmt.Println("prices:", prices)

	discountedPrices := []float64{2.5, 1.5}
	prices = append(prices, discountedPrices...)
	fmt.Println("prices:", prices)
}

// func main() {
// 	// var productNames [4]string = [4]string{"Mobile", "Tablet", "Laptop", "Desktop"}
// 	prices := [4]float64{1.23, 2.34, 3.45, 0.0}

// 	// fmt.Println(prices)
// 	// fmt.Println(productNames[0])
// 	// productNames[0] = "Laptop"
// 	// fmt.Println(productNames[0])

// 	featuredPrices := prices[1:3]
// 	featuredPrices[0] = 9.99
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println("length and capacity of prices:", len(prices), cap(prices))
// 	fmt.Println("length and capacity of featuredPrices:", len(featuredPrices), cap(featuredPrices))
// 	fmt.Println("length and capacity of highlightedPrices:", len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println("length and capacity of highlightedPrices:", len(highlightedPrices), cap(highlightedPrices))

// }
