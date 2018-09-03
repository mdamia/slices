package slice

// AddPerson func
func AddPerson(people []string, person string) []string {
	people = append(people, person)
	return people
}
