package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"
const defaultBalance = 1000.0

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("-----------------------------------")
		// panic("Could not read balance file")
	}

	fmt.Println("Welcome to the bank!")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)

	// https://www.redhat.com/sysadmin/linux-file-permissions-explained
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return defaultBalance, errors.New("could not read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return defaultBalance, errors.New("could not parse balance")
	}

	// https://golang.org/pkg/strconv/#ParseFloat
	return balance, nil
}
