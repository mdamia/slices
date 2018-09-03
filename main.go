package main

import (
	"log"
)

func main() {
	people := []string{}
	people = addPerson(people, "Jacob", "Amira", "Mohammed")

	for _, person := range people {
		log.Println(person)
	}

}
