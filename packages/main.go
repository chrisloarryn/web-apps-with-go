package main

import (
	"github.com/ccontreraso/web-apps-with-go/helpers"
	"log"
)

func main() {
	log.Println("Hello.World")

	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}