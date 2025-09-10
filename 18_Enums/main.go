package main

import "fmt"

// enums are a collection of constants

type Roles string

const (
	Admin   Roles = "Admin"
	User          = "User"
	Guest         = "Guest"
	Student       = "Student"
)

func changeOfRoles(role Roles) {
	fmt.Println("role changed to", role)
}

func main() {
	changeOfRoles(Admin)
}
