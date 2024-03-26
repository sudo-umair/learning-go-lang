package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"Mobile", "Tablet"}
	prices := [4]float64{1.23, 2.34, 3.45, 0.0}

	fmt.Println(prices)
	fmt.Println(productNames[0])
	productNames[0] = "Laptop"
	fmt.Println(productNames[0])
}
