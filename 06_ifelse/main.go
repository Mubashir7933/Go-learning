package main

import "fmt"

func main() {
	// if else statements

	// x := 30
	// if x <= 20 {
	// 	fmt.Println("x is less than or equal to 20")
	// } else {
	// 	fmt.Println("x is greater than 20")
	// }

	// defining a variable in if statement
	// if x := 30; x <= 30 {
	// 	fmt.Println("x is less than or equal to 30")
	// } else {
	// 	fmt.Println("x is greater than 30")
	// }

	// role based if else access statements
	role := "admin"
	IsPermission := true

	if role == "admin" && IsPermission == false {
		println("you can  access the admin dashboard")
	} else if role == "admin" && IsPermission == true {
		fmt.Println("you can access the admin dashboard with full permissions")
	} else {
		fmt.Println("you can't access the admin dashboard with full permissions")
	}

}

// terminary operators are not available in golang yet
