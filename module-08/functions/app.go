package main

import "fmt"

type TransformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// numbers2 := []int{6, 7, 8, 9, 10}

	// doubledNumbers := transformNumbers(&numbers, double)
	// fmt.Println("Numbers:", numbers, "Doubled Numbers:", doubledNumbers)

	// tripledNumbers := transformNumbers(&numbers, triple)
	// fmt.Println("Numbers:", numbers, "Tripled Numbers:", tripledNumbers)

	// transformerFn1 := getTransformerFn(&numbers)
	// transformedNumbers1 := transformNumbers(&numbers, transformerFn1)
	// fmt.Println("Numbers:", numbers, "Transformed Numbers:", transformedNumbers1)

	// transformerFn2 := getTransformerFn(&numbers2)
	// transformedNumbers2 := transformNumbers(&numbers2, transformerFn2)
	// fmt.Println("Numbers:", numbers2, "Transformed Numbers:", transformedNumbers2)

	numbersTransformed := transformNumbers(&numbers, func(i int) int {
		return i * 32
	})
	fmt.Println("Numbers:", numbers, "Transformed Numbers:", numbersTransformed)

}

func transformNumbers(numbers *[]int, transform TransformFunc) []int {
	transformedNumbers := make([]int, len(*numbers))
	for i, n := range *numbers {
		transformedNumbers[i] = transform(n)
	}
	return transformedNumbers
}

// func getTransformerFn(numbers *[]int) TransformFunc {
// 	if (*numbers)[0] == 1 {
// 		return double
// 	} else {
// 		return triple
// 	}
// }

// func double(n int) int {
// 	return n * 2
// }

// func triple(n int) int {
// 	return n * 3
// }
