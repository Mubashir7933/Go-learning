package main

import "fmt"

// simple switch case
// func main() {
// 	i := 2
// 	switch i {
// 	case 1:
// 		fmt.Println("this is case 1")
// 	case 2:
// 		fmt.Println("this is case 2")
// 	default:
// 		fmt.Println("this is default case")
// 	}
// }

// multiple switch cases
// func main() {

// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's the weekend")
// 	case time.Monday:
// 		fmt.Println("It's Monday")

// 	case time.Tuesday:
// 		fmt.Println("It's Tuesday")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}
// }

func main() {
	typeCheck := func(i interface{}) {
		switch k := i.(type) {
		case int:
			fmt.Println("i is an integer", k)
		case string:
			fmt.Println("i is a string", k)
		case bool:
			fmt.Println("i is a boolean", k)
		default:
			fmt.Println("i is other type", k)
		}
	}

	typeCheck(55.9)
}
