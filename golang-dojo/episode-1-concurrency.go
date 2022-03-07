package main

import (
	"fmt"
	"time"
)

func main () {
	start := time.Now()
	defer func () {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Cristobal", "Juan", "Elias", "Luis", "Mauricio", "Aracely", "Valentin", "Maria"}

	for _, evilNinja := range evilNinjas {
		go Attack(evilNinja)
	}
}

//attack
func Attack(target string) {
	// fmt.Println('throwing ninja starts at ',  target)
	fmt.Println("throwing ninja starts at ",  target)
}