package main

import "fmt"

func main() {
	age := 30
	agePointer := &age

	fmt.Println("Age", age)
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years", adultYears)

}

func getAdultYears(age *int) int {
	return *age - 18
}
