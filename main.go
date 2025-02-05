package main

import (
	"fmt"
	"projects/signin"
)

func main() {
	var User signin.UserPass
	var passwordCheck string
	fmt.Println("Enter your email")
	fmt.Scan(&User.Email)

	fmt.Println("Create your password: ")
	fmt.Scan(&User.Password)

	for {

		fmt.Println("Confirm your password: ")
		fmt.Scan(&passwordCheck)

		if signin.VerifyPassword(User.Password, passwordCheck) {
			signin.SigninPage(User)
		} else {
			fmt.Println("Passwords are not simple.")
		}
	}
}
