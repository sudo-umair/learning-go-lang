package main

import (
	"fmt"

	"com/bank/file_ops"

	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	accountBalance, err := file_ops.ReadFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("-----------------------------------")
		// panic("Could not read balance file")
	}

	fmt.Println("Welcome to the bank!")
	fmt.Println(randomdata.PhoneNumber())

	for {
		fmt.Println("What would you like to do today?")
		presentOptions()

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
			file_ops.WriteFloatToFile(balanceFile, accountBalance)
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
			file_ops.WriteFloatToFile(balanceFile, accountBalance)
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
