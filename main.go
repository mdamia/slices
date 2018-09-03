package main

import (
	"log"

	"github.com/mdamia/slices/slice"
	_ "github.com/mdamia/utils"
)

func main() {
	people := []string{}
	people = slice.AddPerson(people, "Jacob")

	for _, person := range people {
		log.Println(person)
	}

}
