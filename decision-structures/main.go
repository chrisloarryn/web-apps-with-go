package main

import "log"

func main() {
	cat := "cats"

	switch cat {
	case "cat":
		log.Println("cat is set to:", cat)
	case "dog":
		log.Println("cat is set to:", cat)
	case "fish":
		log.Println("cat is set to:", cat)
	default:
		log.Println("cat is something else")
	}
}

// func main() {
// 	cat := "cat"

// 	if cat == "cat" {
// 		log.Println("cat is cat")
// 	} else {
// 		log.Println("cat is not cat")
// 	}
// }