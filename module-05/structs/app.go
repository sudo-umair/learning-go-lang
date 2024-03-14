package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*user, error) {

	if (firstName == "") || (lastName == "") || (birthdate == "") {
		return nil, errors.New("Invalid user details")
	}

	user := user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}

	return &user, nil
}

func main() {
	firstNameVal := getUserData("Please enter your first name: ")
	lastNameVal := getUserData("Please enter your last name: ")
	birthdateVal := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user1, err := newUser(firstNameVal, lastNameVal, birthdateVal)

	if err != nil {
		fmt.Printf("ERROR! %s\n", err)
		return
	}

	user1.outputUserDetails()
	user1.clearUserName()
	user1.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
