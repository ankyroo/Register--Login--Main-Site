package signin

import "fmt"

type UserPass struct {
	Email    string
	Password string
}

func VerifyPassword(password, confirmPassword string) bool {
	return password == confirmPassword
}

func SigninPage(User UserPass) {
	var EmailCheck, PasswordCheck string
	for {
		fmt.Println("Enter your email: ")
		fmt.Scan(&User.Email)

		fmt.Println("Enter your password: ")
		fmt.Scan(&User.Password)

		if EmailCheck == User.Password {
			fmt.Println("Works")
		}

		if PasswordCheck == User.Password {
			fmt.Println("It workrs")
			return
		} else {
			fmt.Println("Email or Password is wrong")
		}
	}
}
