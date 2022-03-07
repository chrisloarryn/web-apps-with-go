package main

import (
	"github.com/ccontreraso/web-apps-with-go/helpers"
	"log"
)

const numPool = 10

func CalculateValue (intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan
	log.Println(num)
}
