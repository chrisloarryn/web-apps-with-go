package main

import "log"

type myStruct struct {
	FirstName string `json:"first_name"`
}

func (m *myStruct) printFirstName() string {
	return m.FirstName 
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Valentin1"

	myVar2 := myStruct{
		FirstName: "Valentin2",
	}

	log.Println("myVar is set to: ", myVar.printFirstName())
	log.Println("myVar2 is set to: ", myVar2.printFirstName())
}