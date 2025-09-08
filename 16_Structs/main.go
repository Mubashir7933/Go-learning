package main

import (
	"fmt"
	"time"
)

// struct embedding
type User struct {
	Name  string
	phone int
}

type Reservation struct {
	id           int
	Name         string
	Email        string
	createdAt    time.Time
	availability string
	price        float64
	User
}

// constructor with pointer receiver

// func newReservervation(id int, Name string, Email string, price float64) *Reservation {
// 	centerPort := Reservation{
// 		id:        id,
// 		Name:      Name,
// 		Email:     Email,
// 		createdAt: time.Now(),
// 		price:     price,
// 	}
// 	return &centerPort
// }

// func (res *Reservation) getStatus(availability string) {
// 	res.availability = availability
// 	fmt.Println(availability)
// }

// func (res Reservation) getPrice() float64 {
// 	return res.price
// }

func main() {

	// object of struct embedding
	hotel := Reservation{
		id:    1,
		Name:  "Mubashir",
		Email: "mubashir030601@gmail.com",
		User: User{
			Name:  "Mubashir",
			phone: 123456,
		},
	}
	hotel.phone = 987654
	fmt.Println(hotel.phone)

	// centerPort := Reservation{
	// 	id:           1,
	// 	Name:         "Mubashir",
	// 	Email:        "mubashir@gmail.com",
	// 	createdAt:    time.Now(),
	// 	availability: "open",
	// 	price:        500.00,
	// }

	// centerPort := newReservervation(1, "Mubashir", "mubashir030601@gmail.com", 500)

	// centerPort.getStatus("closed")
	// fmt.Println(centerPort.availability)
	// fmt.Println(centerPort.getPrice())

	//single use struct
	// centerPort := struct {
	// 	id    int
	// 	name  string
	// 	price float64
	// }{1, "Mubashir", 500.00}

	// struct with pointer

	// func ( res *Reservation ) getStatus(availability string)  {

	// }

	//Struct embedding

}
