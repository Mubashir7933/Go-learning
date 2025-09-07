package main

import (
	"fmt"
	"time"
)

type Reservation struct {
	id           int
	Name         string
	Email        string
	createdAt    time.Time
	availability string
}

func (res *Reservation) getStatus(availability string) {
	res.availability = availability
	fmt.Println(availability)
}

func main() {

	centerPort := Reservation{
		id:           1,
		Name:         "Mubashir",
		Email:        "mubashir@gmail.com",
		createdAt:    time.Now(),
		availability: "open",
	}

	centerPort.getStatus("closed")
	fmt.Println(centerPort.availability)

	// struct with pointer

	// func ( res *Reservation ) getStatus(availability string)  {

	// }

}
