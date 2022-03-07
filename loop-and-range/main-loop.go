package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	myMap := make(map[string]User) // map[string]int => myMap["son"] = 1

	me := User {
		FirstName: "Valentin",
		LastName: "Contreras",
	}

	myMap["me"] = me

	log.Println("my map", myMap["me"].FirstName)
}