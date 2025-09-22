package main

import (
	"fmt"

	"github.com/Mubashir7933/Go-learning/auth"
	"github.com/Mubashir7933/Go-learning/user"
)

func main() {
	// Learning of how to work with packages in Go
	auth.LoginWithYourCredentials("Mubashir", "mypassword123")
	session := auth.Session()

	println("Your session id is:", session)

	UserObj := user.User{
		Username: "Mubashir",
		Email:    "mubashir@gmail.com",
	}
	fmt.Println("User details are:", UserObj)
	fmt.Println("Username is:", UserObj.Username)
	fmt.Println("Email is:", UserObj.Email)
}
