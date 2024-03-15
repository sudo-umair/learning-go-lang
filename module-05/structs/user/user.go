package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if (firstName == "") || (lastName == "") || (birthdate == "") {
		return nil, errors.New("invalid user details")
	}

	user := User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}

	return &user, nil
}
