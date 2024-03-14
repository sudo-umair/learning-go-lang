package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to the bank!")
	fmt.Println("What would you like to do today?")
	fmt.Println("Please select an option:")

	fmt.Println("1. Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	fmt.Println("You selected:", choice)

	// if (choice < 1) || (choice > 4) {
	// 	fmt.Println("Invalid choice")
	// 	return
	// }

	if choice == 1 { // Balance
		fmt.Printf("Your account balance is: %.2f\n", accountBalance)
	} else if choice == 2 { // Deposit
		fmt.Print("Enter amount to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("Your new account balance is: %.2f\n", accountBalance)
	} else if choice == 3 { // Withdraw
		fmt.Print("Enter amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Printf("Your new account balance is: %.2f\n", accountBalance)
	} else if choice == 4 { // Exit
		fmt.Println("Thank you for using the bank!")
	} else {
		fmt.Println("Invalid choice")
	}

}