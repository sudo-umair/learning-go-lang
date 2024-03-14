package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}

func main() {
	firstNameVal := getUserData("Please enter your first name: ")
	lastNameVal := getUserData("Please enter your last name: ")
	birthdateVal := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user1 := user{
		firstName: firstNameVal,
		lastName:  lastNameVal,
		createdAt: time.Now(),
		birthdate: birthdateVal,
	}

	user1.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
