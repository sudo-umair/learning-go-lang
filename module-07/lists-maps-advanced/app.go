package main

import "fmt"

type FloatMap map[string]float64

func (f FloatMap) output() {
	fmt.Println("Output", f)
}

func (f FloatMap) keyValPairs() {
	for key, value := range f {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func main() {
	userNames := make([]string, 2) // make([]T, length, capacity)
	fmt.Println("userNames:", userNames)

	userNames[0] = "John"
	userNames[1] = "Doe"
	fmt.Println("userNames:", userNames)

	userNames = append(userNames, "Smith")
	fmt.Println("userNames:", userNames)

	courseRatings := make(FloatMap, 2)
	courseRatings["Go Fundamentals"] = 4.7
	courseRatings["Docker Deep Dive"] = 4.6
	courseRatings["Kubernetes for Developers"] = 4.8
	// fmt.Println("Course Ratings:", courseRatings)
	courseRatings.output()
	courseRatings.keyValPairs()
}
