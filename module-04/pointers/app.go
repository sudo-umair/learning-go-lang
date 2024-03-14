package main

import "fmt"

func main() {
	age := 30
	agePointer := &age

	fmt.Println("Age", *agePointer)
	editAgeUsingPointer(agePointer)
	fmt.Println("Adult years", *agePointer)

}

func editAgeUsingPointer(age *int) {
	*age -= 18
}
