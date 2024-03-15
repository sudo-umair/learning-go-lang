package main

import (
	"com/structs/user"
	"fmt"
)

func main() {
	firstNameVal := getUserData("Please enter your first name: ")
	lastNameVal := getUserData("Please enter your last name: ")
	birthdateVal := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user1, err := user.New(firstNameVal, lastNameVal, birthdateVal)

	if err != nil {
		fmt.Printf("ERROR! %s\n", err)
		return
	}

	user1.OutputUserDetails()
	user1.ClearUserName()
	user1.OutputUserDetails()

	admin := user.NewAdmin("admin1@gmail.com", "password")
	admin.OutputUserDetails()
	admin.ClearUserName()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
