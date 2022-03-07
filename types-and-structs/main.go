package main

import (
	"log"
	"time"
)
 
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Valentin",
		LastName: "Contreras",
		PhoneNumber: "123",
		Age: 0,
		BirthDate: time.Date(2020, 9, 23, 21, 50, 0,0, time.Now().UTC().Location()),
	}

	log.Fatalln(user.BirthDate)
}
