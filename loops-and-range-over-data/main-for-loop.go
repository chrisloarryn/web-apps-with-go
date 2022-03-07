package main

import "log"


func main() {
	/* animals := make(map[string]string)
	animals["dog"] = "Juanito Perro"
	animals["cat"] = "Juanito Gato" */
	/* firstLine := "Once upon the midnight dreary"

	for index, letter := range firstLine {
		log.Println(index, ":", letter)
	} */

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User 

	users = append(users, User{"Jhon", "Smith", "jhon@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
	/* for animalType, animal := range firstLine{
		log.Println(animalType, animal)
	} */
}

// func main() {
// 	animals := []string{"dog", "fish", "cow", "cat"}
// 	for _, animal := range animals{
// 		log.Println(animal)
// 	}
// }

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		log.Println(i)
// 	}
// }