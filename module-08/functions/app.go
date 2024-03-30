package main

import "fmt"

type TransformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Println("Numbers:", numbers, "Doubled Numbers:", doubledNumbers)

	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Println("Numbers:", numbers, "Tripled Numbers:", tripledNumbers)
}

func transformNumbers(numbers *[]int, transform TransformFunc) []int {
	transformedNumbers := make([]int, len(*numbers))
	for i, n := range *numbers {
		transformedNumbers[i] = transform(n)
	}
	return transformedNumbers
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}
