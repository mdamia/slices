package main

import (
	"log"

	"github.com/mdamia/slices/slice"
	"github.com/mdamia/slices/utils"
)

func main() {
	people := []string{}
	people = slice.AddPerson(people, "jacob")

	for _, person := range people {
		log.Println(utils.FirstToUC(person))

	}

}
