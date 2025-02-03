package main

import (
	"fmt"
	"projects/handler"
)

func main() {
	var RegisterNickname, RegisterPassword, PasswordCheck string

	fmt.Println("Create your nickname: ")
	fmt.Scan(&RegisterNickname)

	fmt.Println("Create your password: ")
	fmt.Scan(&RegisterPassword)
	for {
		fmt.Println("Confirm your password: ")
		fmt.Scan(&PasswordCheck)

		if PasswordCheck == RegisterPassword {
			handler.SuccessPage()
			break
		} else {
			fmt.Println("Passwords do not match. Please try again.")
		}
	}
}
