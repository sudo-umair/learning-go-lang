package main

import "fmt"

func main() {
	// task 1
	hobbies := [3]string{"reading", "coding", "gaming"}
	fmt.Println("Hobbies:", hobbies)

	// task 2
	fmt.Println("First hobby:", hobbies[0])
	fmt.Println("Second and third hobbies:", hobbies[1:])

	// task 3
	slice1 := hobbies[:2]
	slice2 := hobbies[0:2]
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	// task 4
	slice1 = hobbies[1:]
	fmt.Println("Resliced slice 1:", slice1)

	// task 5
	courseGoals := []string{"learn Go", "build projects"}
	fmt.Println("Course goals:", courseGoals)

	// task 6
	courseGoals[1] = "build more projects"
	courseGoals = append(courseGoals, "learn more about Go")
	fmt.Println("Updated course goals:", courseGoals)

	// task 7
	type Product struct {
		title string
		id    int
		price float64
	}

	products := []Product{
		{"Mobile", 1, 1000.0},
		{"Tablet", 2, 500.0},
	}
	products = append(products, Product{"Laptop", 3, 1500.0})
	fmt.Println("Products:", products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
