package auth

import "fmt"

func LoginWithYourCredentials(username string, password string) {
	fmt.Println("You are logged in with username:", username)
	fmt.Println("\nYour password is:", password)
}
