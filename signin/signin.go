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
	fmt.Printf("Welcome!")
}
