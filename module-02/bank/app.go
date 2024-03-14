package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to the bank!")

	for {
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

		switch choice {
		case 1: // Balance
			fmt.Printf("Your account balance is: %.2f\n", accountBalance)

		case 2: // Deposit
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount!, must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Your new account balance is: %.2f\n", accountBalance)

		case 3: // Withdraw
			fmt.Print("Enter amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount!, must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Printf("Insufficient funds! Your account balance is: %.2f\n", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Your new account balance is: %.2f\n", accountBalance)

		case 4: // Exit
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using the bank!")
			return

		default: // Invalid choice
			fmt.Println("Invalid choice")

		}

	}

}
